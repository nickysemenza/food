/*
 * Gourd Recipe Database
 *
 * API for https://github.com/nickysemenza/gourd
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: n@nickysemenza.com
 * Generated by: https://openapi-generator.tech
 */

/// RecipeWrapper : A recipe with subcomponents



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct RecipeWrapper {
    /// id
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "detail")]
    pub detail: crate::models::RecipeDetail,
}

impl RecipeWrapper {
    /// A recipe with subcomponents
    pub fn new(id: String, detail: crate::models::RecipeDetail) -> RecipeWrapper {
        RecipeWrapper {
            id,
            detail,
        }
    }
}


