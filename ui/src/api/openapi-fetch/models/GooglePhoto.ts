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
 * A google photo
 * @export
 * @interface GooglePhoto
 */
export interface GooglePhoto {
    /**
     * id
     * @type {string}
     * @memberof GooglePhoto
     */
    id: string;
    /**
     * public image
     * @type {string}
     * @memberof GooglePhoto
     */
    baseUrl: string;
    /**
     * blur hash
     * @type {string}
     * @memberof GooglePhoto
     */
    blurHash?: string;
    /**
     * when it was taken
     * @type {Date}
     * @memberof GooglePhoto
     */
    created: Date;
    /**
     * width px
     * @type {number}
     * @memberof GooglePhoto
     */
    width: number;
    /**
     * height px
     * @type {number}
     * @memberof GooglePhoto
     */
    height: number;
}

export function GooglePhotoFromJSON(json: any): GooglePhoto {
    return GooglePhotoFromJSONTyped(json, false);
}

export function GooglePhotoFromJSONTyped(json: any, ignoreDiscriminator: boolean): GooglePhoto {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'baseUrl': json['base_url'],
        'blurHash': !exists(json, 'blur_hash') ? undefined : json['blur_hash'],
        'created': (new Date(json['created'])),
        'width': json['width'],
        'height': json['height'],
    };
}

export function GooglePhotoToJSON(value?: GooglePhoto | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'base_url': value.baseUrl,
        'blur_hash': value.blurHash,
        'created': (value.created.toISOString()),
        'width': value.width,
        'height': value.height,
    };
}


