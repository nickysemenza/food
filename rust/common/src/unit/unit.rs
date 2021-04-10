#[derive(Clone, PartialEq, PartialOrd, Debug, Default)]
pub struct BareMeasurement(String, f32);

#[derive(Clone, PartialEq, PartialOrd, Debug)]
pub struct Measure(Unit, f32);

// #[derive(Clone, PartialEq, PartialOrd, Debug)]
// pub enum Measure {
//     Other(BareMeasurement),
//     Grams(f32),
//     Ml(f32),
//     Teaspoon(f32),
//     Cent(f32),
// }

#[derive(Clone, PartialEq, PartialOrd, Debug)]
pub enum MeasureKind {
    Weight,
    Volume,
    Money,
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
    Cent,
    Dollar,
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

            "dollar" | "$" => Self::Dollar,
            "cent" => Self::Cent,
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
            Unit::Cent => "cent",
            Unit::Dollar => "$",
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
        return Measure::parse(BareMeasurement(a.unit, a.value));
    }
    pub fn normalize(self) -> Measure {
        let m = self.clone();
        let foo = match self.0 {
            Unit::Gram => (Unit::Gram, m.1),
            Unit::Kilogram => (Unit::Gram, m.1 * G_TO_K),

            Unit::Ounce => (Unit::Gram, m.1 * GRAM_TO_OZ),

            Unit::Milliliter => (Unit::Milliliter, m.1),
            Unit::Liter => (Unit::Milliliter, m.1 * G_TO_K),

            Unit::Teaspoon => (Unit::Teaspoon, m.1),
            Unit::Tablespoon => (Unit::Teaspoon, m.1 * TSP_TO_TBSP),
            Unit::Cup => (Unit::Teaspoon, m.1 * TSP_TO_CUP),
            Unit::Quart => (Unit::Teaspoon, m.1 * CUP_TO_QUART * TSP_TO_CUP),
            Unit::FluidOunce => (Unit::Teaspoon, m.1 * TSP_TO_FL_OZ),

            Unit::Cent => (Unit::Cent, m.1),
            Unit::Dollar => (Unit::Cent, m.1 * 100.0),

            Unit::Other(x) => (Unit::Other(x), m.1),
        };
        return Measure(foo.0, foo.1);
    }
    pub fn parse(m: BareMeasurement) -> Measure {
        let foo = Measure(Unit::from_str(singular(m.0.as_ref()).as_ref()), m.1).normalize();
        return Measure(foo.0, foo.1);
    }
    pub fn kind(self) -> MeasureKind {
        return match self.0 {
            Unit::Gram => MeasureKind::Weight,
            Unit::Cent => MeasureKind::Money,
            Unit::Teaspoon | Unit::Milliliter => MeasureKind::Volume,

            Unit::Other(_) | _ => MeasureKind::Other,
        };
    }

    pub fn convert(self, target: MeasureKind, mappings: Vec<(Measure, Measure)>) -> Measure {
        let curr_kind = self.clone().kind();
        // let inp = self.as_bare();
        for m in mappings.iter() {
            let (a, b) = (m.0.clone().kind(), m.1.clone().kind());
            if a == target && b == curr_kind {
                return Measure(
                    m.0.clone().normalize().0.clone(),
                    m.0.clone().normalize().1 / m.1.clone().normalize().1 * self.clone().1,
                );
            }
            if a == curr_kind && b == target {
                dbg!(m);
                return Measure(
                    dbg!(m.1.clone().normalize()).0.clone(),
                    m.1.clone().normalize().1 / dbg!(m.0.clone().normalize()).1 * self.clone().1,
                );
            }
        }

        Measure(Unit::Other("foo".to_string()), 1.0)
    }
    pub fn as_bare(self) -> BareMeasurement {
        let m = self.1;
        let (val, u, f) = match self.0 {
            Unit::Gram => {
                if m < 1000.0 {
                    (m, Unit::Gram, 1.0)
                } else {
                    (m, Unit::Kilogram, G_TO_K)
                }
            }
            Unit::Milliliter => {
                if m < 1000.0 {
                    (m, Unit::Milliliter, 1.0)
                } else {
                    (m, Unit::Liter, G_TO_K)
                }
            }
            Unit::Teaspoon => match m {
                m if { m < 3.0 } => (m, Unit::Teaspoon, 1.0),
                m if { m < 12.0 } => (m, Unit::Tablespoon, TSP_TO_TBSP),
                m if { m < CUP_TO_QUART * TSP_TO_CUP } => (m, Unit::Cup, TSP_TO_CUP),
                _ => (m, Unit::Teaspoon, 1.0),
            },

            Unit::Cent => (m, Unit::Cent, 1.0),
            Unit::Other(o) => (m, Unit::Other(o), 1.0),
            _ => (m, Unit::Other("".to_string()), 1.0),
        };
        return BareMeasurement(u.to_str(), val / f);
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
        assert_eq!(m1, Measure(Unit::Teaspoon, 48.0));
        assert_eq!(m1.as_bare(), BareMeasurement("cup".to_string(), 1.0));
        assert_eq!(
            Measure::from_string("25.2 grams".to_string()).as_bare(),
            BareMeasurement("g".to_string(), 25.2)
        );
        assert_eq!(
            Measure::from_string("2500.2 grams".to_string()).as_bare(),
            BareMeasurement("kg".to_string(), 2.5002)
        );
        assert_eq!(
            Measure::from_string("12 foo".to_string()).as_bare(),
            BareMeasurement("foo".to_string(), 12.0)
        );
    }

    #[test]
    fn test_convert() {
        let m = Measure::from_string("1 tbsp".to_string());
        let tbsp_dollars = (
            Measure::from_string("2 tbsp".to_string()),
            Measure::from_string("4 dollars".to_string()),
        );
        assert_eq!(
            Measure::from_string("2 dollars".to_string()),
            m.convert(MeasureKind::Money, vec![tbsp_dollars])
        );
    }
}
