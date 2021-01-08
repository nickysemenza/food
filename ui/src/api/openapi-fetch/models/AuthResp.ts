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
 * todo
 * @export
 * @interface AuthResp
 */
export interface AuthResp {
    /**
     * 
     * @type {object}
     * @memberof AuthResp
     */
    user: object;
    /**
     * 
     * @type {string}
     * @memberof AuthResp
     */
    jwt: string;
}

export function AuthRespFromJSON(json: any): AuthResp {
    return AuthRespFromJSONTyped(json, false);
}

export function AuthRespFromJSONTyped(json: any, ignoreDiscriminator: boolean): AuthResp {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'user': json['user'],
        'jwt': json['jwt'],
    };
}

export function AuthRespToJSON(value?: AuthResp | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'user': value.user,
        'jwt': value.jwt,
    };
}


