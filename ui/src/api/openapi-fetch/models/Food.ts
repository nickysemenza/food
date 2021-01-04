/* tslint:disable */
/* eslint-disable */
/**
 * Gourd Recipe Database
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    BrandedFood,
    BrandedFoodFromJSON,
    BrandedFoodFromJSONTyped,
    BrandedFoodToJSON,
    FoodCategory,
    FoodCategoryFromJSON,
    FoodCategoryFromJSONTyped,
    FoodCategoryToJSON,
    FoodNutrient,
    FoodNutrientFromJSON,
    FoodNutrientFromJSONTyped,
    FoodNutrientToJSON,
    FoodPortion,
    FoodPortionFromJSON,
    FoodPortionFromJSONTyped,
    FoodPortionToJSON,
} from './';

/**
 * A top level food
 * @export
 * @interface Food
 */
export interface Food {
    /**
     * FDC Id
     * @type {number}
     * @memberof Food
     */
    fdcId: number;
    /**
     * Food description
     * @type {string}
     * @memberof Food
     */
    description: string;
    /**
     * todo
     * @type {string}
     * @memberof Food
     */
    dataType: FoodDataTypeEnum;
    /**
     * 
     * @type {FoodCategory}
     * @memberof Food
     */
    category?: FoodCategory;
    /**
     * todo
     * @type {Array<FoodNutrient>}
     * @memberof Food
     */
    nutrients: Array<FoodNutrient>;
    /**
     * portion datapoints
     * @type {Array<FoodPortion>}
     * @memberof Food
     */
    portions?: Array<FoodPortion>;
    /**
     * 
     * @type {BrandedFood}
     * @memberof Food
     */
    brandedInfo?: BrandedFood;
}

/**
* @export
* @enum {string}
*/
export enum FoodDataTypeEnum {
    FOUNDATION_FOOD = 'foundation_food',
    SAMPLE_FOOD = 'sample_food',
    MARKET_ACQUISITION = 'market_acquisition',
    SURVEY_FNDDS_FOOD = 'survey_fndds_food',
    SUB_SAMPLE_FOOD = 'sub_sample_food',
    AGRICULTURAL_ACQUISITION = 'agricultural_acquisition',
    SR_LEGACY_FOOD = 'sr_legacy_food',
    BRANDED_FOOD = 'branded_food'
}

export function FoodFromJSON(json: any): Food {
    return FoodFromJSONTyped(json, false);
}

export function FoodFromJSONTyped(json: any, ignoreDiscriminator: boolean): Food {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'fdcId': json['fdc_id'],
        'description': json['description'],
        'dataType': json['data_type'],
        'category': !exists(json, 'category') ? undefined : json['category'],
        'nutrients': ((json['nutrients'] as Array<any>).map(FoodNutrientFromJSON)),
        'portions': !exists(json, 'portions') ? undefined : ((json['portions'] as Array<any>).map(FoodPortionFromJSON)),
        'brandedInfo': !exists(json, 'branded_info') ? undefined : json['branded_info'],
    };
}

export function FoodToJSON(value?: Food | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'fdc_id': value.fdcId,
        'description': value.description,
        'data_type': value.dataType,
        'category': value.category,
        'nutrients': ((value.nutrients as Array<any>).map(FoodNutrientToJSON)),
        'portions': value.portions === undefined ? undefined : ((value.portions as Array<any>).map(FoodPortionToJSON)),
        'branded_info': value.brandedInfo,
    };
}


