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
        
                #[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
                pub struct ClientSecretJwtParamsIn {
                        /// The base64-encoded secret used for signing the JWT.
                        #[serde(rename = "secretBase64")]
                        pub secret_base64: String,
                        /// Optional secret identifier. If supplied, this will be populated in the JWT header in the `kid` field.
                        #[serde(rename = "secretId", skip_serializing_if = "Option::is_none")]
                        pub secret_id: Option<String>,
                        #[serde(rename = "signingAlgorithm")]
                        pub signing_algorithm: models::OauthJwsSigningAlgorithm,
                        /// Optional number of seconds after which the JWT should expire. Defaults to 300 seconds.
                        #[serde(rename = "tokenExpirySecs", skip_serializing_if = "Option::is_none")]
                        pub token_expiry_secs: Option<i32>,
                    }

                    impl ClientSecretJwtParamsIn {
                    pub fn new(secret_base64: String, signing_algorithm: models::OauthJwsSigningAlgorithm) -> ClientSecretJwtParamsIn {
                ClientSecretJwtParamsIn {
                    secret_base64,
                    secret_id: None,
                    signing_algorithm,
                    token_expiry_secs: None,
                    }
                    }
                    }

