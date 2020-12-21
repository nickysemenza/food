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
    PaginatedRecipes,
    PaginatedRecipesFromJSON,
    PaginatedRecipesToJSON,
    RecipeWrapper,
    RecipeWrapperFromJSON,
    RecipeWrapperToJSON,
    SearchResult,
    SearchResultFromJSON,
    SearchResultToJSON,
} from '../models';

export interface RecipesApiCreateRecipesRequest {
    recipeWrapper: RecipeWrapper;
}

export interface RecipesApiGetRecipeByIdRequest {
    recipeId: string;
}

export interface RecipesApiListRecipesRequest {
    offset?: number;
    limit?: number;
}

export interface RecipesApiSearchRequest {
    name: string;
    offset?: number;
    limit?: number;
}

/**
 * 
 */
export class RecipesApi extends runtime.BaseAPI {

    /**
     * Create a recipe
     */
    async createRecipesRaw(requestParameters: RecipesApiCreateRecipesRequest): Promise<runtime.ApiResponse<RecipeWrapper>> {
        if (requestParameters.recipeWrapper === null || requestParameters.recipeWrapper === undefined) {
            throw new runtime.RequiredError('recipeWrapper','Required parameter requestParameters.recipeWrapper was null or undefined when calling createRecipes.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = typeof token === 'function' ? token("bearerAuth", []) : token;

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/recipes`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: RecipeWrapperToJSON(requestParameters.recipeWrapper),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => RecipeWrapperFromJSON(jsonValue));
    }

    /**
     * Create a recipe
     */
    async createRecipes(requestParameters: RecipesApiCreateRecipesRequest): Promise<RecipeWrapper> {
        const response = await this.createRecipesRaw(requestParameters);
        return await response.value();
    }

    /**
     * Info for a specific recipe
     */
    async getRecipeByIdRaw(requestParameters: RecipesApiGetRecipeByIdRequest): Promise<runtime.ApiResponse<RecipeWrapper>> {
        if (requestParameters.recipeId === null || requestParameters.recipeId === undefined) {
            throw new runtime.RequiredError('recipeId','Required parameter requestParameters.recipeId was null or undefined when calling getRecipeById.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = typeof token === 'function' ? token("bearerAuth", []) : token;

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/recipes/{recipe_id}`.replace(`{${"recipe_id"}}`, encodeURIComponent(String(requestParameters.recipeId))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => RecipeWrapperFromJSON(jsonValue));
    }

    /**
     * Info for a specific recipe
     */
    async getRecipeById(requestParameters: RecipesApiGetRecipeByIdRequest): Promise<RecipeWrapper> {
        const response = await this.getRecipeByIdRaw(requestParameters);
        return await response.value();
    }

    /**
     * List all recipes
     */
    async listRecipesRaw(requestParameters: RecipesApiListRecipesRequest): Promise<runtime.ApiResponse<PaginatedRecipes>> {
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
            path: `/recipes`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => PaginatedRecipesFromJSON(jsonValue));
    }

    /**
     * List all recipes
     */
    async listRecipes(requestParameters: RecipesApiListRecipesRequest): Promise<PaginatedRecipes> {
        const response = await this.listRecipesRaw(requestParameters);
        return await response.value();
    }

    /**
     * Search recipes and ingredients
     */
    async searchRaw(requestParameters: RecipesApiSearchRequest): Promise<runtime.ApiResponse<SearchResult>> {
        if (requestParameters.name === null || requestParameters.name === undefined) {
            throw new runtime.RequiredError('name','Required parameter requestParameters.name was null or undefined when calling search.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.name !== undefined) {
            queryParameters['name'] = requestParameters.name;
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
            path: `/search`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => SearchResultFromJSON(jsonValue));
    }

    /**
     * Search recipes and ingredients
     */
    async search(requestParameters: RecipesApiSearchRequest): Promise<SearchResult> {
        const response = await this.searchRaw(requestParameters);
        return await response.value();
    }

}
