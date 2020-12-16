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
    Recipe,
    RecipeFromJSON,
    RecipeFromJSONTyped,
    RecipeToJSON,
} from './';

/**
 * A search result wrapper, which contains ingredients and recipes
 * @export
 * @interface SearchResult
 */
export interface SearchResult {
    /**
     * The ingredients
     * @type {Array<Ingredient>}
     * @memberof SearchResult
     */
    ingredients?: Array<Ingredient>;
    /**
     * The recipes
     * @type {Array<Recipe>}
     * @memberof SearchResult
     */
    recipes?: Array<Recipe>;
}

export function SearchResultFromJSON(json: any): SearchResult {
    return SearchResultFromJSONTyped(json, false);
}

export function SearchResultFromJSONTyped(json: any, ignoreDiscriminator: boolean): SearchResult {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'ingredients': !exists(json, 'ingredients') ? undefined : ((json['ingredients'] as Array<any>).map(IngredientFromJSON)),
        'recipes': !exists(json, 'recipes') ? undefined : ((json['recipes'] as Array<any>).map(RecipeFromJSON)),
    };
}

export function SearchResultToJSON(value?: SearchResult | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'ingredients': value.ingredients === undefined ? undefined : ((value.ingredients as Array<any>).map(IngredientToJSON)),
        'recipes': value.recipes === undefined ? undefined : ((value.recipes as Array<any>).map(RecipeToJSON)),
    };
}

