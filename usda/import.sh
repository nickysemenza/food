#!/bin/sh
set -euf -o pipefail
tables=(food_category food food_attribute_type acquisition_sample agricultural_acquisition branded_food food_attribute food_nutrient_conversion_factor food_calorie_conversion_factor food_component nutrient food_nutrient_source food_nutrient_derivation food_nutrient measure_unit food_portion food_protein_conversion_factor foundation_food input_food lab_method lab_method_code lab_method_nutrient market_acquisition nutrient_incoming_name retention_factor sample_food sr_legacy_food sub_sample_food sub_sample_result wweia_food_category survey_fndds_food)

alias p='psql "postgresql://food:food@localhost:5555/food"'

p -c "select count(*) from food";
p -c "truncate table food_category cascade;"
for f in ${tables[@]}; do
    echo $f
    headers=$(head -n1 $f.csv | tr -d '"')
    sed 's/""/NULL/g' $f.csv > $f_n.csv
    p -c "\copy $f($headers) from '$f_n.csv' (format csv, null \"NULL\", DELIMITER ',', HEADER);"
done