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
    Amount,
    AmountFromJSON,
    AmountFromJSONTyped,
    AmountToJSON,
    UnitMapping,
    UnitMappingFromJSON,
    UnitMappingFromJSONTyped,
    UnitMappingToJSON,
} from './';

/**
 * 
 * @export
 * @interface UnitConversionRequest
 */
export interface UnitConversionRequest {
    /**
     * 
     * @type {string}
     * @memberof UnitConversionRequest
     */
    target?: UnitConversionRequestTargetEnum;
    /**
     * multiple amounts to try
     * @type {Array<Amount>}
     * @memberof UnitConversionRequest
     */
    input: Array<Amount>;
    /**
     * mappings of equivalent units
     * @type {Array<UnitMapping>}
     * @memberof UnitConversionRequest
     */
    unit_mappings: Array<UnitMapping>;
}

/**
* @export
* @enum {string}
*/
export enum UnitConversionRequestTargetEnum {
    WEIGHT = 'weight',
    VOLUME = 'volume',
    MONEY = 'money',
    OTHER = 'other'
}

export function UnitConversionRequestFromJSON(json: any): UnitConversionRequest {
    return UnitConversionRequestFromJSONTyped(json, false);
}

export function UnitConversionRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UnitConversionRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'target': !exists(json, 'target') ? undefined : json['target'],
        'input': ((json['input'] as Array<any>).map(AmountFromJSON)),
        'unit_mappings': ((json['unit_mappings'] as Array<any>).map(UnitMappingFromJSON)),
    };
}

export function UnitConversionRequestToJSON(value?: UnitConversionRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'target': value.target,
        'input': ((value.input as Array<any>).map(AmountToJSON)),
        'unit_mappings': ((value.unit_mappings as Array<any>).map(UnitMappingToJSON)),
    };
}

