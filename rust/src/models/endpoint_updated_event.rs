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
        
            /// EndpointUpdatedEvent : Sent when an endpoint is updated.
                #[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
                pub struct EndpointUpdatedEvent {
                        #[serde(rename = "data")]
                        pub data: Box<models::EndpointUpdatedEventData>,
                        #[serde(rename = "type")]
                        pub r#type: Type,
                    }

                    impl EndpointUpdatedEvent {
                        /// Sent when an endpoint is updated.
                    pub fn new(data: models::EndpointUpdatedEventData, r#type: Type) -> EndpointUpdatedEvent {
                EndpointUpdatedEvent {
                    data: Box::new(data),
                    r#type,
                    }
                    }
                    }
                    /// 
                    #[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
                    pub enum Type {
                            #[serde(rename = "endpoint.updated")]
                        EndpointPeriodUpdated,
                    }

                    impl Default for Type {
                    fn default() -> Type {
                        Self::EndpointPeriodUpdated
                    }
                    }

