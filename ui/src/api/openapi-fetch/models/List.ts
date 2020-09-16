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
/**
 * 
 * @export
 * @interface List
 */
export interface List {
    /**
     * What number apge this is
     * @type {number}
     * @memberof List
     */
    pageNumber: number;
    /**
     * How many items were requested for this page
     * @type {number}
     * @memberof List
     */
    limit: number;
    /**
     * Total number of items across all pages
     * @type {number}
     * @memberof List
     */
    totalCount: number;
}

export function ListFromJSON(json: any): List {
    return ListFromJSONTyped(json, false);
}

export function ListFromJSONTyped(json: any, ignoreDiscriminator: boolean): List {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'pageNumber': json['page_number'],
        'limit': json['limit'],
        'totalCount': json['total_count'],
    };
}

export function ListToJSON(value?: List | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'page_number': value.pageNumber,
        'limit': value.limit,
        'total_count': value.totalCount,
    };
}


