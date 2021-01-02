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
    Ingredient,
    IngredientFromJSON,
    IngredientFromJSONTyped,
    IngredientToJSON,
    RecipeDetail,
    RecipeDetailFromJSON,
    RecipeDetailFromJSONTyped,
    RecipeDetailToJSON,
} from './';

/**
 * Ingredients in a single section
 * @export
 * @interface SectionIngredient
 */
export interface SectionIngredient {
    /**
     * id
     * @type {string}
     * @memberof SectionIngredient
     */
    id: string;
    /**
     * what kind of ingredient
     * @type {string}
     * @memberof SectionIngredient
     */
    kind: SectionIngredientKindEnum;
    /**
     * 
     * @type {RecipeDetail}
     * @memberof SectionIngredient
     */
    recipe?: RecipeDetail;
    /**
     * 
     * @type {Ingredient}
     * @memberof SectionIngredient
     */
    ingredient?: Ingredient;
    /**
     * weight in grams
     * @type {number}
     * @memberof SectionIngredient
     */
    grams: number;
    /**
     * amount
     * @type {number}
     * @memberof SectionIngredient
     */
    amount?: number;
    /**
     * unit
     * @type {string}
     * @memberof SectionIngredient
     */
    unit?: string;
    /**
     * adjective
     * @type {string}
     * @memberof SectionIngredient
     */
    adjective?: string;
    /**
     * optional
     * @type {boolean}
     * @memberof SectionIngredient
     */
    optional?: boolean;
    /**
     * raw line item (pre-import/scrape)
     * @type {string}
     * @memberof SectionIngredient
     */
    original?: string;
}

export function SectionIngredientFromJSON(json: any): SectionIngredient {
    return SectionIngredientFromJSONTyped(json, false);
}

export function SectionIngredientFromJSONTyped(json: any, ignoreDiscriminator: boolean): SectionIngredient {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'kind': json['kind'],
        'recipe': !exists(json, 'recipe') ? undefined : json['recipe'],
        'ingredient': !exists(json, 'ingredient') ? undefined : json['ingredient'],
        'grams': json['grams'],
        'amount': !exists(json, 'amount') ? undefined : json['amount'],
        'unit': !exists(json, 'unit') ? undefined : json['unit'],
        'adjective': !exists(json, 'adjective') ? undefined : json['adjective'],
        'optional': !exists(json, 'optional') ? undefined : json['optional'],
        'original': !exists(json, 'original') ? undefined : json['original'],
    };
}

export function SectionIngredientToJSON(value?: SectionIngredient | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'kind': value.kind,
        'recipe': value.recipe,
        'ingredient': value.ingredient,
        'grams': value.grams,
        'amount': value.amount,
        'unit': value.unit,
        'adjective': value.adjective,
        'optional': value.optional,
        'original': value.original,
    };
}

/**
* @export
* @enum {string}
*/
export enum SectionIngredientKindEnum {
    RECIPE = 'recipe',
    INGREDIENT = 'ingredient'
}


