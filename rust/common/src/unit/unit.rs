#[derive(Clone, PartialEq, PartialOrd, Debug, Default)]

pub struct Measurement(String, f32);
#[derive(Clone, PartialEq, PartialOrd, Debug)]
pub enum Measure {
    Other(Measurement),
    Grams(f32),
    Ml(f32),
    Teaspoon(f32),
}

#[derive(Clone, PartialEq, PartialOrd, Debug)]
pub enum MeasureKind {
    Weight,
    Volume,
    Other,
}
#[derive(Clone, PartialEq, PartialOrd, Debug)]
pub enum Unit {
    Gram,
    Kilogram,
    Liter,
    Milliliter,
    Teaspoon,
    Tablespoon,
    Cup,
    Quart,
    FluidOunce,
    Ounce,
    Other(String),
}

impl Unit {
    pub fn from_str(s: &str) -> Self {
        match s {
            "gram" | "g" => Self::Gram,
            "kilogram" | "kg" => Self::Kilogram,

            "oz" | "ounce" => Self::Ounce,

            "ml" => Self::Milliliter,
            "l" => Self::Liter,

            "tsp" | "teaspoon" => Self::Teaspoon,
            "tbsp" | "tablespoon" => Self::Tablespoon,
            "c" | "cup" => Self::Cup,
            "q" | "quart" => Self::Quart,
            "fl oz" | "fluid oz" => Self::FluidOunce,

            _ => Self::Other(s.to_string()),
        }
    }
    pub fn to_str(self) -> String {
        match self {
            Unit::Gram => "g",
            Unit::Kilogram => "kg",
            Unit::Liter => "l",
            Unit::Milliliter => "ml",
            Unit::Teaspoon => "tsp",
            Unit::Tablespoon => "tbsp",
            Unit::Cup => "cup",
            Unit::Quart => "quart",
            Unit::FluidOunce => "fl oz",
            Unit::Ounce => "oz",
            Unit::Other(s) => return s,
        }
        .to_string()
    }
}

// multiplication factors
const TSP_TO_TBSP: f32 = 3.0;
const TSP_TO_FL_OZ: f32 = 2.0;
const G_TO_K: f32 = 1000.0;
const CUP_TO_QUART: f32 = 4.0;
const TSP_TO_CUP: f32 = 48.0;
const GRAM_TO_OZ: f32 = 28.3495;

impl Measure {
    pub fn from_string(s: String) -> Measure {
        let a = ingredient::parse_amount(s.as_str()).unwrap()[0].clone();
        return Measure::parse(Measurement(a.unit, a.value));
    }
    pub fn parse(m: Measurement) -> Measure {
        return match Unit::from_str(singular(m.0.as_ref()).as_ref()) {
            Unit::Gram => Self::Grams(m.1),
            Unit::Kilogram => Self::Grams(m.1 * G_TO_K),

            Unit::Ounce => Self::Grams(m.1 * GRAM_TO_OZ),

            Unit::Milliliter => Self::Ml(m.1),
            Unit::Liter => Self::Ml(m.1 * G_TO_K),

            Unit::Teaspoon => Self::Teaspoon(m.1),
            Unit::Tablespoon => Self::Teaspoon(m.1 * TSP_TO_TBSP),
            Unit::Cup => Self::Teaspoon(m.1 * TSP_TO_CUP),
            Unit::Quart => Self::Teaspoon(m.1 * CUP_TO_QUART * TSP_TO_CUP),
            Unit::FluidOunce => Self::Teaspoon(m.1 * TSP_TO_FL_OZ),

            Unit::Other(_) => Self::Other(m),
        };
    }
    pub fn kind(self) -> MeasureKind {
        return match self {
            Measure::Other(_) => MeasureKind::Other,
            Measure::Grams(_) => MeasureKind::Weight,
            Measure::Teaspoon(_) | Measure::Ml(_) => MeasureKind::Volume,
        };
    }

    pub fn normalize(self) -> Measurement {
        let (m, u, f) = match self {
            Measure::Other(m) => (m.1, Unit::Other(m.0), 1.0),
            Measure::Grams(m) => {
                if m < 1000.0 {
                    (m, Unit::Gram, 1.0)
                } else {
                    (m, Unit::Kilogram, G_TO_K)
                }
            }
            Measure::Ml(m) => {
                if m < 1000.0 {
                    (m, Unit::Milliliter, 1.0)
                } else {
                    (m, Unit::Liter, G_TO_K)
                }
            }
            Measure::Teaspoon(m) => match m {
                m if { m < 3.0 } => (m, Unit::Teaspoon, 1.0),
                m if { m < 12.0 } => (m, Unit::Tablespoon, TSP_TO_TBSP),
                m if { m < CUP_TO_QUART * TSP_TO_CUP } => (m, Unit::Cup, TSP_TO_CUP),
                _ => (m, Unit::Teaspoon, 1.0),
            },
        };
        return Measurement(u.to_str(), m / f);
    }

    // Err("todo".to_string())
}
pub fn singular(s: &str) -> String {
    s.strip_suffix("s").unwrap_or(s).to_lowercase()
}

#[cfg(test)]
mod tests {

    use super::*;
    #[test]
    fn test_measure() {
        // let m1 = Measure::parse(Measurement("Tbsp".to_string(), 16.0));
        let m1 = Measure::from_string("16 tbsp".to_string());
        assert_eq!(m1, Measure::Teaspoon(48.0));
        assert_eq!(m1.normalize(), Measurement("cup".to_string(), 1.0));
        assert_eq!(
            Measure::from_string("25.2 grams".to_string()).normalize(),
            Measurement("g".to_string(), 25.2)
        );
        assert_eq!(
            Measure::from_string("2500.2 grams".to_string()).normalize(),
            Measurement("kg".to_string(), 2.5002)
        );
        assert_eq!(
            Measure::from_string("12 foo".to_string()).normalize(),
            Measurement("foo".to_string(), 12.0)
        );
    }
}
