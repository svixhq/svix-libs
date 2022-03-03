use bb8::Pool;
use bb8_redis::RedisConnectionManager;
use redis::{AsyncCommands, RedisError};
use serde::{de::DeserializeOwned, Deserialize, Serialize};

use crate::core::types::{ApplicationId, OrganizationId};

/// Errors internal to the cache
#[derive(thiserror::Error, Debug)]
pub enum Error {
    #[error("error deserializing Redis value")]
    Deserialization(#[from] serde_json::error::Error),

    #[error("Redis pool error")]
    Pool(#[from] bb8::RunError<RedisError>),

    #[error("Redis databse error")]
    Database(#[from] RedisError),
}
type Result<T> = std::result::Result<T, Error>;

pub trait CacheKey: AsRef<str> {
    const PREFIX_CACHE: &'static str = "USER_CACHE";
}
pub trait CacheValue: DeserializeOwned + Serialize {
    type Key: CacheKey;
}

/// A macro that creates a [`CacheKey`] and ties it to any value that implements
/// DeserializeOwned`] and [`Serialize`]
macro_rules! kv_def {
    ($key_id:ident, $val_struct:ident) => {
        #[derive(Clone, Debug)]
        pub struct $key_id(String);

        impl AsRef<str> for $key_id {
            fn as_ref(&self) -> &str {
                &self.0
            }
        }

        impl CacheKey for $key_id {}

        impl CacheValue for $val_struct {
            type Key = $key_id;
        }
    };
}

// FIXME: Find out the actual data that needs to be cached
#[derive(Deserialize, Serialize, Debug, Clone)]
pub struct AppEndpointValue {
    test_value: String,
}
kv_def!(AppEndpointKey, AppEndpointValue);
impl AppEndpointKey {
    // FIXME: Rewrite doc comment when AppEndpointValue members are known
    /// Returns a key for fetching all cached endpoints for a given organization and application.
    pub fn new(org: OrganizationId, app: ApplicationId) -> AppEndpointKey {
        AppEndpointKey(format!("{}_APP_v3_{}_{}", Self::PREFIX_CACHE, org, app))
    }
}

/// A Redis-based cache of data to avoid expensive fetches from PostgreSQL. Simply a wrapper over
/// Redis.
pub struct RedisCache {
    redis: Pool<RedisConnectionManager>,
}

impl RedisCache {
    pub fn new(redis: Pool<RedisConnectionManager>) -> RedisCache {
        RedisCache { redis }
    }

    pub async fn get<T: CacheValue>(&self, key: &T::Key) -> Result<Option<T>> {
        let mut pool = self.redis.get().await?;
        let fetched = pool.get::<&str, Option<String>>(key.as_ref()).await?;

        if let Some(fetched) = fetched {
            Ok(Some(serde_json::from_str(&fetched)?))
        } else {
            Ok(None)
        }
    }

    pub async fn set<T: CacheValue>(
        &self,
        key: &T::Key,
        value: &T,
        ttl: Option<usize>,
    ) -> Result<()> {
        let mut pool = self.redis.get().await?;

        if let Some(ttl) = ttl {
            pool.set_ex(key.as_ref(), serde_json::to_string(value)?, ttl)
                .await?;
        } else {
            pool.set(key.as_ref(), serde_json::to_string(value)?)
                .await?;
        }

        Ok(())
    }

    pub async fn delete<T: CacheKey>(&self, key: &T) -> Result<()> {
        let mut pool = self.redis.get().await?;
        pool.del(key.as_ref()).await?;

        Ok(())
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    // Test structures

    #[derive(Deserialize, Serialize, Debug, PartialEq)]
    struct TestValA(usize);
    kv_def!(TestKeyA, TestValA);
    impl TestKeyA {
        fn new(id: String) -> TestKeyA {
            TestKeyA(format!("SVIX_TEST_KEY_A_{}", id))
        }
    }

    #[derive(Deserialize, Serialize, Debug, PartialEq)]
    struct TestValB(String);
    kv_def!(TestKeyB, TestValB);
    impl TestKeyB {
        fn new(id: String) -> TestKeyB {
            TestKeyB(format!("SVIX_TEST_KEY_B_{}", id))
        }
    }

    #[tokio::test]
    async fn test_cache_crud_no_ttl() {
        // FIXME: Configurable addr
        let redis_pool = bb8::Pool::builder()
            .build(RedisConnectionManager::new("redis://localhost:6379").unwrap())
            .await
            .unwrap();

        let cache = RedisCache::new(redis_pool.clone());

        let (first_key, first_val_a, first_val_b) =
            (TestKeyA::new("1".to_owned()), TestValA(1), TestValA(2));
        let (second_key, second_val_a, second_val_b) = (
            TestKeyB::new("1".to_owned()),
            TestValB("1".to_owned()),
            TestValB("2".to_owned()),
        );

        // Create
        assert!(cache.set(&first_key, &first_val_a, None).await.is_ok());
        assert!(cache.set(&second_key, &second_val_a, None).await.is_ok());

        // Read
        assert_eq!(cache.get(&first_key).await.unwrap(), Some(first_val_a));
        assert_eq!(cache.get(&second_key).await.unwrap(), Some(second_val_a));

        // Update (overwrite)
        assert!(cache.set(&first_key, &first_val_b, None).await.is_ok());
        assert!(cache.set(&second_key, &second_val_b, None).await.is_ok());

        // Confirm update
        assert_eq!(cache.get(&first_key).await.unwrap(), Some(first_val_b));
        assert_eq!(cache.get(&second_key).await.unwrap(), Some(second_val_b));

        // Delete
        assert!(cache.delete(&first_key).await.is_ok());
        assert!(cache.delete(&second_key).await.is_ok());

        // Confirm deletion
        assert_eq!(cache.get::<TestValA>(&first_key).await.unwrap(), None);
        assert_eq!(cache.get::<TestValB>(&second_key).await.unwrap(), None);
    }

    #[tokio::test]
    async fn test_cache_ttl() {
        let redis_pool = bb8::Pool::builder()
            .build(RedisConnectionManager::new("redis://localhost:6379").unwrap())
            .await
            .unwrap();
        let cache = RedisCache::new(redis_pool.clone());
        let key = TestKeyA::new("key".to_owned());

        assert!(cache.set(&key, &TestValA(1), Some(1)).await.is_ok());
        tokio::time::sleep(std::time::Duration::from_millis(1200)).await;
        assert_eq!(cache.get::<TestValA>(&key).await.unwrap(), None);
    }
}
