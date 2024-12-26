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
                pub struct ListResponseStreamOut {
                        #[serde(rename = "data")]
                        pub data: Vec<models::StreamOut>,
                        #[serde(rename = "done")]
                        pub done: bool,
                        #[serde(rename = "iterator", deserialize_with = "Option::deserialize")]
                        pub iterator: Option<String>,
                        #[serde(rename = "prevIterator", skip_serializing_if = "Option::is_none")]
                        pub prev_iterator: Option<String>,
                    }

                    impl ListResponseStreamOut {
                    pub fn new(data: Vec<models::StreamOut>, done: bool, iterator: Option<String>) -> ListResponseStreamOut {
                ListResponseStreamOut {
                    data,
                    done,
                    iterator,
                    prev_iterator: None,
                    }
                    }
                    }

