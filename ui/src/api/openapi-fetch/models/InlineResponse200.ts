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
    GooglePhotosAlbum,
    GooglePhotosAlbumFromJSON,
    GooglePhotosAlbumFromJSONTyped,
    GooglePhotosAlbumToJSON,
} from './';

/**
 * 
 * @export
 * @interface InlineResponse200
 */
export interface InlineResponse200 {
    /**
     * The list of albums
     * @type {Array<GooglePhotosAlbum>}
     * @memberof InlineResponse200
     */
    albums?: Array<GooglePhotosAlbum>;
}

export function InlineResponse200FromJSON(json: any): InlineResponse200 {
    return InlineResponse200FromJSONTyped(json, false);
}

export function InlineResponse200FromJSONTyped(json: any, ignoreDiscriminator: boolean): InlineResponse200 {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'albums': !exists(json, 'albums') ? undefined : ((json['albums'] as Array<any>).map(GooglePhotosAlbumFromJSON)),
    };
}

export function InlineResponse200ToJSON(value?: InlineResponse200 | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'albums': value.albums === undefined ? undefined : ((value.albums as Array<any>).map(GooglePhotosAlbumToJSON)),
    };
}


