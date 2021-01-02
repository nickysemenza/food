import React, { useState } from "react";
import AsyncCreatableSelect, {
  Props as AsyncProps,
} from "react-select/async-creatable";
import { Ingredient, IngredientsApi } from "../api/openapi-fetch";

import { useSearch } from "../api/openapi-hooks/api";
import { getOpenapiFetchConfig } from "../config";
import { ActionMeta, OptionsType, OptionTypeBase } from "react-select";

import { toast } from "react-toastify";
import { IngredientKind } from "./RecipeEditorUtils";

export interface Results {
  ingredients: ResultType;
  recipes: ResultType;
}

export interface ResultItem {
  // title: string;
  id: string;
  kind: IngredientKind;
}
export interface ResultType {
  name: "ingredients" | "recipes";
  results: ResultItem[];
}

export interface PropsTest extends AsyncProps<OptionTypeBase, false> {
  name: string;
}

const IngredientSearch: React.FC<{
  callback: (
    ingredient: Pick<Ingredient, "id" | "name">,
    kind: IngredientKind
  ) => void;
  initial?: string;
}> = ({ callback, initial }) => {
  const i = initial || "";
  const [value, setValue] = useState(i);
  const [v, setV] = useState<any>({ label: i });

  const iApi = new IngredientsApi(getOpenapiFetchConfig());
  const { data } = useSearch({ queryParams: { name: value } });

  const handleCreate = async (inputValue: any) => {
    console.log("foo", inputValue);
    let res = await iApi.createIngredients({
      ingredient: { name: inputValue.value, id: "" },
    });
    toast(`ingredient ${res.id} created`);

    // let res = (await createIngredientMutation()).data;
    if (res) {
      callback({ id: res.id, name: res.name }, "ingredient");
    }
  };

  const handleChange = async (newValue: any, actionMeta: ActionMeta<any>) => {
    console.group("Value Changed");
    console.log(newValue);
    console.log(`action: ${actionMeta.action}`);
    console.groupEnd();
    if (newValue.__isNew__) {
      handleCreate(newValue);
    } else {
      callback({ name: newValue.label, id: newValue.id }, newValue.kind);
    }
    setV(newValue);
  };

  const loadOptions = (
    inputValue: string,
    callback: (options: OptionsType<any>) => void
  ) => {
    setValue(inputValue || "");
    const x = [
      ...(data?.ingredients || []).map((i) => ({
        label: i.name,
        kind: "ingredient",
        id: i.id,
      })),
      ...(data?.recipes || []).map((i) => ({
        label: i.detail.name + " (Recipe) (v " + i.detail.version + ")",
        kind: "recipe",
        id: i.id,
      })),
    ];
    callback(x);
  };

  return (
    <div data-cy="name-input">
      <AsyncCreatableSelect
        onChange={handleChange}
        loadOptions={loadOptions}
        onCreateOption={handleCreate}
        value={v}
      />
    </div>
  );
};
export default IngredientSearch;
