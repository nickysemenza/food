/*
 * Gourd Recipe Database
 *
 * API for https://github.com/nickysemenza/gourd
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: n@nickysemenza.com
 * Generated by: https://openapi-generator.tech
 */

/// PaginatedPhotos : pages of GooglePhoto



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct PaginatedPhotos {
    #[serde(rename = "photos", skip_serializing_if = "Option::is_none")]
    pub photos: Option<Vec<crate::models::GooglePhoto>>,
    #[serde(rename = "meta", skip_serializing_if = "Option::is_none")]
    pub meta: Option<Box<crate::models::Items>>,
}

impl PaginatedPhotos {
    /// pages of GooglePhoto
    pub fn new() -> PaginatedPhotos {
        PaginatedPhotos {
            photos: None,
            meta: None,
        }
    }
}


