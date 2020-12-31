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
 * where the recipe came from (i.e. book/website)
 * @export
 * @interface RecipeSource
 */
export interface RecipeSource {
    /**
     * url
     * @type {string}
     * @memberof RecipeSource
     */
    url?: string;
    /**
     * title (if book)
     * @type {string}
     * @memberof RecipeSource
     */
    title?: string;
    /**
     * page number/section (if book)
     * @type {string}
     * @memberof RecipeSource
     */
    page?: string;
}

export function RecipeSourceFromJSON(json: any): RecipeSource {
    return RecipeSourceFromJSONTyped(json, false);
}

export function RecipeSourceFromJSONTyped(json: any, ignoreDiscriminator: boolean): RecipeSource {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'url': !exists(json, 'url') ? undefined : json['url'],
        'title': !exists(json, 'title') ? undefined : json['title'],
        'page': !exists(json, 'page') ? undefined : json['page'],
    };
}

export function RecipeSourceToJSON(value?: RecipeSource | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'url': value.url,
        'title': value.title,
        'page': value.page,
    };
}

