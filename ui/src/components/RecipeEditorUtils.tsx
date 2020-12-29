import { Ingredient, RecipeWrapper, TimeRange } from "../api/openapi-hooks/api";
import update from "immutability-helper";
export type Override = {
  sectionID: number;
  ingredientID: number;
  value: number;
  attr: IngredientAttr;
};

export type RecipeTweaks = {
  override?: Override;
  multiplier: number;
  edit: boolean;
};

export type IngredientAttr = "grams" | "amount";
export type IngredientKind = "recipe" | "ingredient";

export const updateIngredientInfo = (
  recipe: RecipeWrapper,
  sectionID: number,
  ingredientID: number,
  ingredient: Ingredient,
  kind: IngredientKind
) => {
  const { id, name } = ingredient;
  return update(recipe, {
    detail: {
      sections: {
        [sectionID]: {
          ingredients: {
            [ingredientID]: {
              recipe: {
                $set:
                  kind === "recipe"
                    ? { id, name, quantity: 0, unit: "", sections: [] }
                    : undefined,
              },
              ingredient: {
                $set: kind === "ingredient" ? { id, name } : undefined,
              },
              kind: { $set: kind },
            },
          },
        },
      },
    },
  });
};

export const isOverride = (
  tweaks: RecipeTweaks,
  sectionID: number,
  ingredientID: number,
  attr: IngredientAttr
) =>
  tweaks.override?.ingredientID === ingredientID &&
  tweaks.override.sectionID === sectionID &&
  tweaks.override.attr === attr;
export const getIngredientValue = (
  tweaks: RecipeTweaks,
  sectionID: number,
  ingredientID: number,
  value: number,
  attr: IngredientAttr
) =>
  (isOverride(tweaks, sectionID, ingredientID, attr) &&
    tweaks.override?.value) ||
  value * tweaks.multiplier;

export const updateInstruction = (
  recipe: RecipeWrapper,
  tweaks: RecipeTweaks,
  sectionID: number,
  instructionID: number,
  value: string
) =>
  tweaks.edit
    ? update(recipe, {
        detail: {
          sections: {
            [sectionID]: {
              instructions: {
                [instructionID]: { instruction: { $set: value } },
              },
            },
          },
        },
      })
    : recipe;

export const updateRecipeName = (recipe: RecipeWrapper, value: string) =>
  update(recipe, {
    detail: { name: { $set: value } },
  });

export const addInstruction = (recipe: RecipeWrapper, sectionID: number) =>
  update(recipe, {
    detail: {
      sections: {
        [sectionID]: {
          instructions: {
            $push: [{ id: "", instruction: "" }],
          },
        },
      },
    },
  });

export const addIngredient = (recipe: RecipeWrapper, sectionID: number) =>
  update(recipe, {
    detail: {
      sections: {
        [sectionID]: {
          ingredients: {
            $push: [
              {
                id: "",
                grams: 1,
                kind: "ingredient",
                // info: { name: "", id: "", __typename: "Ingredient" },
                amount: 0,
                unit: "",
                adjective: "",
                optional: false,
              },
            ],
          },
        },
      },
    },
  });
export const addSection = (recipe: RecipeWrapper) =>
  update(recipe, {
    detail: {
      sections: {
        $push: [
          {
            id: "",
            duration: { min: 0, max: 0 },
            ingredients: [],
            instructions: [],
          },
        ],
      },
    },
  });

export const updateTimeRange = (
  recipe: RecipeWrapper,
  sectionID: number,
  value?: TimeRange
) =>
  !!value
    ? update(recipe, {
        detail: {
          sections: {
            [sectionID]: {
              duration: {
                $set: value,
              },
            },
          },
        },
      })
    : recipe;

export type I = "ingredients" | "instructions";
const calculateMoveI = (
  recipe: RecipeWrapper,
  sectionIndex: number,
  index: number,
  movingUp: boolean,
  i: I
) => {
  const { sections } = recipe.detail;

  const numI = sections[sectionIndex][i].length;
  const numSections = sections.length;
  const firstInSection = index === 0;
  const lastInSection = index === numI - 1;

  let newSectionIndex = sectionIndex;
  let newInIndex: number;
  if (firstInSection && movingUp) {
    // needs to go to prior section
    newSectionIndex--;
    if (newSectionIndex < 0) {
      // out of bounds
      return null;
    }
    newInIndex = sections[newSectionIndex][i].length;
  } else if (!firstInSection && movingUp) {
    // prior row in same section
    newInIndex = index - 1;
  } else if (lastInSection && !movingUp) {
    // needs to go to next section
    newSectionIndex++;
    if (newSectionIndex > numSections - 1) {
      // out of bounds
      return null;
    }
    newInIndex = 0;
  } else {
    // next row in same section
    newInIndex = index + 1;
  }

  return { newSectionIndex, newInIndex };
};
export const canMoveI = (
  recipe: RecipeWrapper,
  sectionIndex: number,
  index: number,
  movingUp: boolean,
  i: I
) => !!calculateMoveI(recipe, sectionIndex, index, movingUp, i);
export const moveI = (
  recipe: RecipeWrapper,
  sectionIndex: number,
  index: number,
  movingUp: boolean,
  i: I
) => {
  const coords = calculateMoveI(recipe, sectionIndex, index, movingUp, i);
  if (!coords) return recipe;
  const { newSectionIndex, newInIndex } = coords;
  console.log("moving!", {
    sectionIndex,
    newSectionIndex,
    index,
    newInIndex,
  });
  const target = recipe.detail.sections[sectionIndex][i][index];
  return update(recipe, {
    detail: {
      sections:
        newSectionIndex === sectionIndex
          ? {
              [sectionIndex]: {
                [i]: {
                  $splice: [
                    [index, 1],
                    [newInIndex, 0, target],
                  ],
                },
              },
            }
          : {
              [sectionIndex]: {
                [i]: {
                  $splice: [[index, 1]],
                },
              },
              [newSectionIndex]: {
                [i]: {
                  $splice: [[newInIndex, 0, target]],
                },
              },
            },
    },
  });
};
export const delI = (
  recipe: RecipeWrapper,
  sectionIndex: number,
  index: number,
  i: I
) =>
  update(recipe, {
    detail: {
      sections: {
        [sectionIndex]: {
          [i]: {
            $splice: [[index, 1]],
          },
        },
      },
    },
  });

// returns the 1-indexed count of the instruction, across all sections.
export const getGlobalInstructionNumber = (
  recipe: RecipeWrapper,
  sectionIndex: number,
  instructionIndex: number
) =>
  recipe.detail.sections
    .slice(0, sectionIndex)
    .map((x) => x.instructions.length)
    .reduce((a, b) => a + b, 0) +
  instructionIndex +
  1;
