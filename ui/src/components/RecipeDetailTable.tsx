import React from "react";
import IngredientSearch from "./IngredientSearch";
import { Link } from "react-router-dom";
import {
  RecipeWrapper,
  Ingredient,
  RecipeSection,
} from "../api/openapi-hooks/api";
import { getIngredient } from "../util";
export interface UpdateIngredientProps {
  sectionID: number;
  ingredientID: number;
  value: string;
  attr: "grams" | "name" | "amount" | "unit" | "adjective" | "optional";
}

export interface TableProps {
  recipe: RecipeWrapper;
  updateIngredient: (i: UpdateIngredientProps) => void;
  updateIngredientInfo: (
    sectionID: number,
    ingredientID: number,
    ingredient: Pick<Ingredient, "id" | "name">,
    kind: "recipe" | "ingredient"
  ) => void;
  updateInstruction: (
    sectionID: number,
    instructionID: number,
    value: string
  ) => void;

  getIngredientValue: (
    sectionID: number,
    ingredientID: number,
    value: number,
    attr: "grams" | "amount"
  ) => number;
  edit: boolean;
  addInstruction: (sectionID: number) => void;
  addIngredient: (sectionID: number) => void;
  addSection: () => void;
}
const RecipeDetailTable: React.FC<TableProps> = ({
  recipe,
  updateIngredient,
  updateIngredientInfo,
  updateInstruction,
  getIngredientValue,
  edit,
  addInstruction,
  addIngredient,
  addSection,
}) => {
  // for baker's percentage cauclation we need the total mass of all flours (which together are '100%')
  const flourMass = (recipe.detail.sections || []).reduce(
    (acc, section) =>
      acc +
      section.ingredients
        .filter((item) => item.ingredient?.name.includes("flour"))
        .reduce((acc, ingredient) => acc + ingredient?.grams, 0),
    0
  );
  const showBP = flourMass > 0;

  const renderRow = (section: RecipeSection, x: number) => (
    <TableRow key={x}>
      <TableCell>
        <div className="inline-block bg-blue-200 text-blue-800 text-xs px-2 rounded-full uppercase font-semibold tracking-wide">
          {String.fromCharCode(65 + x)}
        </div>
      </TableCell>
      <TableCell>{section.minutes}</TableCell>
      <TableCell>
        {section.ingredients.map((ingredient, y) => {
          const bp = Math.round((ingredient.grams / flourMass) * 100);
          return (
            <div className="ing-table-row" key={y}>
              <TableInput
                data-cy="grams-input"
                edit={edit}
                softEdit
                value={getIngredientValue(x, y, ingredient.grams || 0, "grams")}
                onChange={(e) =>
                  updateIngredient({
                    sectionID: x,
                    ingredientID: y,
                    value: e.target.value,
                    attr: "grams",
                  })
                }
              />
              <div className="flex space-x-0.5">
                <div className="text-gray-600">g</div>
                {showBP && (
                  <div
                    className={`${
                      bp > 0 ? "text-gray-600" : "text-red-300"
                    } italic`}
                  >
                    ({bp}%)
                  </div>
                )}
              </div>
              {edit ? (
                <IngredientSearch
                  initial={getIngredient(ingredient).name}
                  callback={(item, kind) =>
                    updateIngredientInfo(x, y, item, kind)
                  }
                />
              ) : (
                <div className="text-gray-600">
                  {ingredient.kind === "recipe" ? (
                    <Link
                      to={`/recipe/${ingredient.recipe?.id}`}
                      className="link"
                    >
                      {ingredient.recipe?.name}
                    </Link>
                  ) : (
                    ingredient.ingredient?.name
                  )}
                </div>
              )}
              <TableInput
                data-cy="amount-input"
                // width={16}
                edit={edit}
                softEdit
                value={getIngredientValue(
                  x,
                  y,
                  ingredient.amount || 0,
                  "amount"
                )}
                onChange={(e) =>
                  updateIngredient({
                    sectionID: x,
                    ingredientID: y,
                    value: e.target.value,
                    attr: "amount",
                  })
                }
              />
              <TableInput
                data-cy="unit-input"
                width={16}
                edit={edit}
                value={ingredient.unit}
                onChange={(e) =>
                  updateIngredient({
                    sectionID: x,
                    ingredientID: y,
                    value: e.target.value,
                    attr: "unit",
                  })
                }
              />
              <TableInput
                data-cy="adjective-input"
                width={16}
                edit={edit}
                value={ingredient.adjective}
                onChange={(e) =>
                  updateIngredient({
                    sectionID: x,
                    ingredientID: y,
                    value: e.target.value,
                    attr: "adjective",
                  })
                }
              />
              {/* TODO: optional toggle */}
            </div>
          );
        })}
        {edit && (
          <div className="add-item" onClick={() => addIngredient(x)}>
            add ingredient
          </div>
        )}
      </TableCell>
      <TableCell>
        {/* <ol className="list-decimal list-inside"> */}
        {section.instructions.map((instruction, y) => (
          <div key={y} className="flex font-serif">
            <div className="mr-2 w-4">{y + 1}. </div>
            <TableInput
              data-cy="instruction-input"
              width={72}
              tall
              edit={edit}
              value={instruction.instruction}
              onChange={(e) => updateInstruction(x, y, e.target.value)}
            />
          </div>
        ))}
        {/* </ol> */}
        {edit && (
          <div className="add-item" onClick={() => addInstruction(x)}>
            add instruction
          </div>
        )}
      </TableCell>
    </TableRow>
  );

  return (
    <div className="border-gray-900 shadow-xl bg-gray-100">
      <TableRow header>
        <TableCell>Section</TableCell>
        <TableCell>Minutes</TableCell>
        <TableCell>
          <div className="ing-table-row font-mono">
            <div>x</div>
            <div>grams (BP)</div>
            <div>of y</div>
            <div>z</div>
            <div>units</div>
            <div>modifier</div>
          </div>
        </TableCell>
        <TableCell>Instructions</TableCell>
      </TableRow>
      {recipe.detail.sections?.map((section, x) => renderRow(section, x))}
      {edit && (
        <div className="add-item" onClick={() => addSection()}>
          add section
        </div>
      )}
    </div>
  );
};
export default RecipeDetailTable;

const TableCell: React.FC = ({ children }) => (
  <div className="border-solid border border-gray-600 p-1">{children}</div>
);
const TableRow: React.FC<{ header?: boolean }> = ({
  children,
  header = false,
}) => (
  <div className={`rec-table-row ${header && "font-semibold"}`}>{children}</div>
);

const TableInput: React.FC<{
  edit: boolean;
  softEdit?: boolean;
  value: string | number;
  width?: number;
  tall?: boolean;
  onChange: (
    event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => void;
}> = ({ edit, softEdit = false, width = 10, tall, ...props }) => {
  const className = `border-2 border-dashed p-0 h-${
    tall ? 18 : 6
  } w-${width} border-gray-200 disabled:border-red-100 hover:border-black ${
    softEdit && !edit && "bg-transparent"
  } focus:bg-gray-200`;
  return edit || softEdit ? (
    tall ? (
      <textarea {...props} className={className} rows={3} />
    ) : (
      <input
        {...props}
        className={className}
        disabled={!edit && props.value === 0}
      />
    )
  ) : (
    <p className="flex flex-wrap">{formatText(props.value)}</p>
  );
};
const re = /[\d]* ?F/g;
const formatText = (text: React.ReactText) => {
  if (typeof text === "number") {
    return text;
  }

  let pairs = [];
  const matches = [...text.matchAll(re)];
  if (matches.length === 0) {
    return text;
  }

  console.log(matches);
  // matches.next
  let lastProcessed = 0;
  for (const match of matches) {
    const matchStart = match.index || 0;
    const matchEnd = matchStart + match[0].length;
    pairs.push(text.substring(lastProcessed, matchStart));
    pairs.push(
      <code className="text-red-800 mx-1">
        {text.substring(matchStart, matchEnd)}
      </code>
    );
    // pairs.push()
    lastProcessed = matchEnd;
    // pairs.push([, ]);
  }
  pairs.push(text.substring(lastProcessed));
  // let res = [];
  // for
  return pairs;

  // console.log(pairs);
};
