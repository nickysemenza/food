/* tslint:disable */
/* eslint-disable */
/**
 * Gourd Recipe Database
 * API for https://github.com/nickysemenza/gourd
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: n@nickysemenza.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    List,
    ListFromJSON,
    ListFromJSONTyped,
    ListToJSON,
    Recipe,
    RecipeFromJSON,
    RecipeFromJSONTyped,
    RecipeToJSON,
} from './';

/**
 * pages of Recipe
 * @export
 * @interface PaginatedRecipes
 */
export interface PaginatedRecipes {
    /**
     * 
     * @type {Array<Recipe>}
     * @memberof PaginatedRecipes
     */
    recipes?: Array<Recipe>;
    /**
     * 
     * @type {List}
     * @memberof PaginatedRecipes
     */
    meta?: List;
}

export function PaginatedRecipesFromJSON(json: any): PaginatedRecipes {
    return PaginatedRecipesFromJSONTyped(json, false);
}

export function PaginatedRecipesFromJSONTyped(json: any, ignoreDiscriminator: boolean): PaginatedRecipes {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'recipes': !exists(json, 'recipes') ? undefined : ((json['recipes'] as Array<any>).map(RecipeFromJSON)),
        'meta': !exists(json, 'meta') ? undefined : ListFromJSON(json['meta']),
    };
}

export function PaginatedRecipesToJSON(value?: PaginatedRecipes | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'recipes': value.recipes === undefined ? undefined : ((value.recipes as Array<any>).map(RecipeToJSON)),
        'meta': ListToJSON(value.meta),
    };
}


