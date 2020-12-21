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


import * as runtime from '../runtime';
import {
    PaginatedPhotos,
    PaginatedPhotosFromJSON,
    PaginatedPhotosToJSON,
} from '../models';

export interface PhotosApiListPhotosRequest {
    offset?: number;
    limit?: number;
}

/**
 * 
 */
export class PhotosApi extends runtime.BaseAPI {

    /**
     * List all photos
     */
    async listPhotosRaw(requestParameters: PhotosApiListPhotosRequest): Promise<runtime.ApiResponse<PaginatedPhotos>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = typeof token === 'function' ? token("bearerAuth", []) : token;

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/photos`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => PaginatedPhotosFromJSON(jsonValue));
    }

    /**
     * List all photos
     */
    async listPhotos(requestParameters: PhotosApiListPhotosRequest): Promise<PaginatedPhotos> {
        const response = await this.listPhotosRaw(requestParameters);
        return await response.value();
    }

}
