/* tslint:disable */
/* eslint-disable */
/**
 * Swagger Recipestore
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
    Recipe,
    RecipeFromJSON,
    RecipeFromJSONTyped,
    RecipeToJSON,
} from './';

/**
 * An Ingredient
 * @export
 * @interface Ingredient
 */
export interface Ingredient {
    /**
     * UUID
     * @type {string}
     * @memberof Ingredient
     */
    id: string;
    /**
     * Ingredient name
     * @type {string}
     * @memberof Ingredient
     */
    name: string;
    /**
     * Recipes referencing this ingredient
     * @type {Array<Recipe>}
     * @memberof Ingredient
     */
    recipes?: Array<Recipe>;
    /**
     * Ingredients that are equivalent
     * @type {Array<Ingredient>}
     * @memberof Ingredient
     */
    children?: Array<Ingredient>;
}

export function IngredientFromJSON(json: any): Ingredient {
    return IngredientFromJSONTyped(json, false);
}

export function IngredientFromJSONTyped(json: any, ignoreDiscriminator: boolean): Ingredient {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'name': json['name'],
        'recipes': !exists(json, 'recipes') ? undefined : ((json['recipes'] as Array<any>).map(RecipeFromJSON)),
        'children': !exists(json, 'children') ? undefined : ((json['children'] as Array<any>).map(IngredientFromJSON)),
    };
}

export function IngredientToJSON(value?: Ingredient | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'name': value.name,
        'recipes': value.recipes === undefined ? undefined : ((value.recipes as Array<any>).map(RecipeToJSON)),
        'children': value.children === undefined ? undefined : ((value.children as Array<any>).map(IngredientToJSON)),
    };
}

