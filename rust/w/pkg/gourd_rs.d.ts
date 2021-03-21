/* tslint:disable */
/* eslint-disable */
/**
* @param {string} input
* @returns {string}
*/
export function parse(input: string): string;
/**
* @param {string} input
* @returns {Ingredient}
*/
export function parse2(input: string): Ingredient;
/**
* @param {string} input
* @returns {any}
*/
export function parse3(input: string): any;
/**
* @param {string} input
* @returns {any}
*/
export function parse4(input: string): any;
/**
* @param {Ingredient} val
* @returns {string}
*/
export function format_ingredient(val: Ingredient): string;
/**
* @param {any} recipe_detail
* @returns {any}
*/
export function sum_ingr(recipe_detail: any): any;

interface Ingredient {
    amounts: Amount[];
    modifier?: string;
    name: string;
  }
  
  interface Amount {
    unit: string;
    value: number;
  }


