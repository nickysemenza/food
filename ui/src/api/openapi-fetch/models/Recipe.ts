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
/**
 * A recipe
 * @export
 * @interface Recipe
 */
export interface Recipe {
    /**
     * UUID
     * @type {string}
     * @memberof Recipe
     */
    id: string;
    /**
     * recipe name
     * @type {string}
     * @memberof Recipe
     */
    name: string;
    /**
     * todo
     * @type {number}
     * @memberof Recipe
     */
    totalMinutes?: number;
    /**
     * book or website? deprecated?
     * @type {string}
     * @memberof Recipe
     */
    source?: string;
    /**
     * num servings
     * @type {number}
     * @memberof Recipe
     */
    servings?: number;
    /**
     * serving quantity
     * @type {number}
     * @memberof Recipe
     */
    quantity: number;
    /**
     * serving unit
     * @type {string}
     * @memberof Recipe
     */
    unit: string;
}

export function RecipeFromJSON(json: any): Recipe {
    return RecipeFromJSONTyped(json, false);
}

export function RecipeFromJSONTyped(json: any, ignoreDiscriminator: boolean): Recipe {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'name': json['name'],
        'totalMinutes': !exists(json, 'total_minutes') ? undefined : json['total_minutes'],
        'source': !exists(json, 'source') ? undefined : json['source'],
        'servings': !exists(json, 'servings') ? undefined : json['servings'],
        'quantity': json['quantity'],
        'unit': json['unit'],
    };
}

export function RecipeToJSON(value?: Recipe | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'name': value.name,
        'total_minutes': value.totalMinutes,
        'source': value.source,
        'servings': value.servings,
        'quantity': value.quantity,
        'unit': value.unit,
    };
}


