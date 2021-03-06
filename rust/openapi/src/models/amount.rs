/*
 * Gourd Recipe Database
 *
 * API for https://github.com/nickysemenza/gourd
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: n@nickysemenza.com
 * Generated by: https://openapi-generator.tech
 */

/// Amount : amount and unit



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct Amount {
    /// unit
    #[serde(rename = "unit")]
    pub unit: String,
    /// value
    #[serde(rename = "value")]
    pub value: f64,
}

impl Amount {
    /// amount and unit
    pub fn new(unit: String, value: f64) -> Amount {
        Amount {
            unit,
            value,
        }
    }
}


