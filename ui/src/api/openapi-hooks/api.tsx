/* Generated by restful-react */

import React from "react";
import {
  Get,
  GetProps,
  useGet,
  UseGetProps,
  Mutate,
  MutateProps,
  useMutate,
  UseMutateProps,
} from "restful-react";

export type Omit<T, K extends keyof T> = Pick<T, Exclude<keyof T, K>>;

const encodingFn = encodeURIComponent;

const encodingTagFactory = (encodingFn: typeof encodeURIComponent) => (
  strings: TemplateStringsArray,
  ...params: (string | number | boolean)[]
) =>
  strings.reduce(
    (accumulatedPath, pathPart, idx) =>
      `${accumulatedPath}${pathPart}${
        idx < params.length ? encodingFn(params[idx]) : ""
      }`,
    ""
  );

const encode = encodingTagFactory(encodingFn);

/**
 * A recipe
 */
export interface Recipe {
  /**
   * UUID
   */
  id: string;
  /**
   * recipe name
   */
  name: string;
  /**
   * todo
   */
  total_minutes?: number;
  /**
   * book or website? deprecated?
   */
  source?: string;
  /**
   * num servings
   */
  servings?: number;
  /**
   * serving quantity
   */
  quantity: number;
  /**
   * serving unit
   */
  unit: string;
}

/**
 * An Ingredient
 */
export interface Ingredient {
  /**
   * UUID
   */
  id: string;
  /**
   * Ingredient name
   */
  name: string;
  /**
   * Recipes referencing this ingredient
   */
  recipes?: Recipe[];
  /**
   * Ingredients that are equivalent
   */
  children?: Ingredient[];
}

export interface Error {
  code: number;
  message: string;
}

export interface List {
  /**
   * What number page this is
   */
  page_number: number;
  /**
   * How many items were requested for this page
   */
  limit: number;
  /**
   * todo
   */
  offset: number;
  /**
   * Total number of items across all pages
   */
  total_count: number;
}

export interface PaginatedRecipes {
  recipes?: Recipe[];
  meta?: List;
}

export interface PaginatedIngredients {
  ingredients?: Ingredient[];
  meta?: List;
}

export interface ListIngredientsQueryParams {
  /**
   * The number of items to skip before starting to collect the result set.
   */
  offset?: number;
  /**
   * The numbers of items to return.
   */
  limit?: number;
}

export type ListIngredientsProps = Omit<
  GetProps<PaginatedIngredients, Error, ListIngredientsQueryParams, void>,
  "path"
>;

/**
 * List all ingredients
 */
export const ListIngredients = (props: ListIngredientsProps) => (
  <Get<PaginatedIngredients, Error, ListIngredientsQueryParams, void>
    path={encode`/ingredients`}
    {...props}
  />
);

export type UseListIngredientsProps = Omit<
  UseGetProps<PaginatedIngredients, Error, ListIngredientsQueryParams, void>,
  "path"
>;

/**
 * List all ingredients
 */
export const useListIngredients = (props: UseListIngredientsProps) =>
  useGet<PaginatedIngredients, Error, ListIngredientsQueryParams, void>(
    encode`/ingredients`,
    props
  );

export interface ListRecipesQueryParams {
  /**
   * The number of items to skip before starting to collect the result set.
   */
  offset?: number;
  /**
   * The numbers of items to return.
   */
  limit?: number;
}

export type ListRecipesProps = Omit<
  GetProps<PaginatedRecipes, Error, ListRecipesQueryParams, void>,
  "path"
>;

/**
 * List all recipes
 */
export const ListRecipes = (props: ListRecipesProps) => (
  <Get<PaginatedRecipes, Error, ListRecipesQueryParams, void>
    path={encode`/recipes`}
    {...props}
  />
);

export type UseListRecipesProps = Omit<
  UseGetProps<PaginatedRecipes, Error, ListRecipesQueryParams, void>,
  "path"
>;

/**
 * List all recipes
 */
export const useListRecipes = (props: UseListRecipesProps) =>
  useGet<PaginatedRecipes, Error, ListRecipesQueryParams, void>(
    encode`/recipes`,
    props
  );

export type CreateRecipesProps = Omit<
  MutateProps<Recipe, Error, void, void, void>,
  "path" | "verb"
>;

/**
 * Create a recipe
 */
export const CreateRecipes = (props: CreateRecipesProps) => (
  <Mutate<Recipe, Error, void, void, void>
    verb="POST"
    path={encode`/recipes`}
    {...props}
  />
);

export type UseCreateRecipesProps = Omit<
  UseMutateProps<Recipe, Error, void, void, void>,
  "path" | "verb"
>;

/**
 * Create a recipe
 */
export const useCreateRecipes = (props: UseCreateRecipesProps) =>
  useMutate<Recipe, Error, void, void, void>("POST", encode`/recipes`, props);

export interface GetRecipeByIdPathParams {
  /**
   * The id of the recipe to retrieve
   */
  recipe_id: string;
}

export type GetRecipeByIdProps = Omit<
  GetProps<Recipe, Error, void, GetRecipeByIdPathParams>,
  "path"
> &
  GetRecipeByIdPathParams;

/**
 * Info for a specific recipe
 */
export const GetRecipeById = ({ recipe_id, ...props }: GetRecipeByIdProps) => (
  <Get<Recipe, Error, void, GetRecipeByIdPathParams>
    path={encode`/recipes/${recipe_id}`}
    {...props}
  />
);

export type UseGetRecipeByIdProps = Omit<
  UseGetProps<Recipe, Error, void, GetRecipeByIdPathParams>,
  "path"
> &
  GetRecipeByIdPathParams;

/**
 * Info for a specific recipe
 */
export const useGetRecipeById = ({
  recipe_id,
  ...props
}: UseGetRecipeByIdProps) =>
  useGet<Recipe, Error, void, GetRecipeByIdPathParams>(
    (paramsInPath: GetRecipeByIdPathParams) =>
      encode`/recipes/${paramsInPath.recipe_id}`,
    { pathParams: { recipe_id }, ...props }
  );