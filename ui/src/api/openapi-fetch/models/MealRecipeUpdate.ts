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
/**
 * an update to the recipes on a mea
 * @export
 * @interface MealRecipeUpdate
 */
export interface MealRecipeUpdate {
    /**
     * Recipe Id
     * @type {string}
     * @memberof MealRecipeUpdate
     */
    recipe_id: string;
    /**
     * multiplier
     * @type {number}
     * @memberof MealRecipeUpdate
     */
    multiplier: number;
    /**
     * todo
     * @type {string}
     * @memberof MealRecipeUpdate
     */
    action: MealRecipeUpdateActionEnum;
}

/**
* @export
* @enum {string}
*/
export enum MealRecipeUpdateActionEnum {
    ADD = 'add',
    REMOVE = 'remove'
}

export function MealRecipeUpdateFromJSON(json: any): MealRecipeUpdate {
    return MealRecipeUpdateFromJSONTyped(json, false);
}

export function MealRecipeUpdateFromJSONTyped(json: any, ignoreDiscriminator: boolean): MealRecipeUpdate {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'recipe_id': json['recipe_id'],
        'multiplier': json['multiplier'],
        'action': json['action'],
    };
}

export function MealRecipeUpdateToJSON(value?: MealRecipeUpdate | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'recipe_id': value.recipe_id,
        'multiplier': value.multiplier,
        'action': value.action,
    };
}


