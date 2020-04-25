import React, { useState } from "react";
import { useGetRecipeByUuidQuery } from "../generated/graphql";

import { Box, Button } from "rebass";
import { useParams } from "react-router-dom";
import RecipeTable from "../components/RecipeTable";
import Debug from "../components/Debug";
import RecipeCard from "../components/RecipeCard";

type override = {
  sectionID: number;
  ingredientID: number;
  value: number;
};
const RecipeDetail: React.FC = () => {
  let { uuid } = useParams();
  const { loading, error, data } = useGetRecipeByUuidQuery({
    variables: { uuid: uuid || "" },
  });
  const [multiplier, setMultiplier] = useState(1.0);
  const [override, setOverride] = useState<override>();
  const [edit, setEdit] = useState(false);
  const [recipe, setRecipe] = useState(data?.recipe);
  if (error) {
    console.error({ error });
    return (
      <Box color="primary" fontSize={4}>
        {error.message}
      </Box>
    );
  }
  if (!recipe && !!data?.recipe) {
    setRecipe(data.recipe);
  }
  if (!recipe) return null;

  const updateIngredient = (
    sectionID: number,
    ingredientID: number,
    value: string,
    attr: "grams" | "name"
  ) => {
    const newValue = parseFloat(value.endsWith(".") ? value + "0" : value);
    // console.log(newValue);
    attr === "grams" &&
      !edit &&
      setOverride({
        sectionID,
        ingredientID,
        value: newValue,
      });
    const { grams } = recipe.sections[sectionID]!.ingredients[ingredientID];
    edit
      ? setRecipe({
          ...recipe,
          sections: recipe.sections.map((item, index) =>
            index === sectionID
              ? {
                  ...item,
                  ingredients: recipe.sections[sectionID].ingredients.map(
                    (item2, index2) =>
                      index2 === ingredientID
                        ? {
                            ...item2,
                            grams:
                              attr === "grams"
                                ? parseFloat(value)
                                : item2.grams,
                            info: {
                              ...item2.info,
                              name: attr === "name" ? value : item2.info.name,
                            },
                          }
                        : item2
                  ),
                }
              : item
          ),
        })
      : setMultiplier(
          Math.round((newValue / grams + Number.EPSILON) * 100) / 100
        );
  };

  const getIngredientValue = (
    sectionID: number,
    ingredientID: number,
    value: number
  ) => {
    if (
      override?.ingredientID === ingredientID &&
      override.sectionID === sectionID
    )
      return override.value;
    return value * multiplier;
  };

  const updateInstruction = (
    sectionID: number,
    instructionID: number,
    value: string
  ) => {
    edit &&
      setRecipe({
        ...recipe,
        sections: recipe.sections.map((item, index) =>
          index === sectionID
            ? {
                ...item,
                instructions: recipe.sections[
                  sectionID
                ].instructions.map((item2, index2) =>
                  index2 === instructionID
                    ? { ...item2, instruction: value }
                    : item2
                ),
              }
            : item
        ),
      });
  };

  const addInstruction = (sectionID: number) => {
    setRecipe({
      ...recipe,
      sections: recipe.sections.map((item, index) =>
        index === sectionID
          ? {
              ...item,
              instructions: [
                ...(recipe.sections[sectionID]?.instructions || []),
                { uuid: "x", instruction: "" },
              ],
            }
          : item
      ),
    });
  };

  const addIngredient = (sectionID: number) => {
    setRecipe({
      ...recipe,
      sections: recipe.sections.map((item, index) =>
        index === sectionID
          ? {
              ...item,
              ingredients: [
                ...(recipe.sections[sectionID].ingredients || []),
                { uuid: "x", grams: 1, info: { name: "" } },
              ],
            }
          : item
      ),
    });
  };

  return (
    <div>
      <Button onClick={() => setMultiplier(1)}>Reset</Button>
      <Button
        onClick={() => {
          setMultiplier(1);
          setEdit(!edit);
        }}
      >
        {edit ? "edit" : "view"}
      </Button>
      <RecipeCard recipe={recipe} />
      <RecipeTable
        updateIngredient={updateIngredient}
        recipe={recipe}
        getIngredientValue={getIngredientValue}
        edit={edit}
        addInstruction={addInstruction}
        addIngredient={addIngredient}
        updateInstruction={updateInstruction}
      />
      <Debug data={{ recipe, loading, error, data, multiplier, override }} />
    </div>
  );
};

export default RecipeDetail;
