/*
 * Svix API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.1.1
 * 
 * Generated by: https://openapi-generator.tech
 */

#[allow(unused_imports)]
use crate::models;
#[allow(unused_imports)]
use serde::{Deserialize, Serialize};
        
                                        #[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
                        #[serde(untagged)]
                        pub enum StreamSinkOut {
                        StreamSinkInOneOf(Box<models::StreamSinkInOneOf>),
                        StreamSinkInOneOf1(Box<models::StreamSinkInOneOf1>),
                        StreamSinkInOneOf2(Box<models::StreamSinkInOneOf2>),
                        StreamSinkInOneOf3(Box<models::StreamSinkInOneOf3>),
                        StreamSinkInOneOf4(Box<models::StreamSinkInOneOf4>),
                        StreamSinkInOneOf5(Box<models::StreamSinkInOneOf5>),
                        StreamSinkInOneOf6(Box<models::StreamSinkInOneOf6>),
                        StreamSinkInOneOf7(Box<models::StreamSinkInOneOf7>),
                        }

                        impl Default for StreamSinkOut {
                        fn default() -> Self {
                    Self::StreamSinkInOneOf(Default::default())
                        }
                        }
                    /// 
                    #[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
                    pub enum Type {
                            #[serde(rename = "azureBlobStorage")]
                        AzureBlobStorage,
                            #[serde(rename = "otelV1HttpTrace")]
                        OtelV1HttpTrace,
                            #[serde(rename = "http")]
                        Http,
                            #[serde(rename = "amazonS3")]
                        AmazonS3,
                            #[serde(rename = "snowflake")]
                        Snowflake,
                            #[serde(rename = "googleCloudStorage")]
                        GoogleCloudStorage,
                            #[serde(rename = "redshift")]
                        Redshift,
                            #[serde(rename = "bigQuery")]
                        BigQuery,
                    }

                    impl Default for Type {
                    fn default() -> Type {
                        Self::AzureBlobStorage
                    }
                    }

