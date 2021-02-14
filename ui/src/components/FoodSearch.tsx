import React, { useEffect, useState } from "react";
import { PlusCircle } from "react-feather";
import { PaginatedFoods, FoodApi } from "../api/openapi-fetch";
import { getOpenapiFetchConfig } from "../config";
import { Code } from "../util";
import { ButtonGroup } from "./Button";
import { getCalories } from "./RecipeEditorUtils";

const FoodSearch: React.FC<{
  name: string;
  highlightId?: number;
  onLink?: (fdc_id: number) => void;
}> = ({ name, highlightId, onLink }) => {
  const [foods, setFoods] = useState<PaginatedFoods>();

  useEffect(() => {
    const fetchData = async () => {
      const bar = new FoodApi(getOpenapiFetchConfig());
      const result = await bar.searchFoods({
        name,
        limit: 5,
        dataTypes: [
          // FoodDataType.BRANDED_FOOD,
          // FoodDataType.FOUNDATION_FOOD,
        ],
      });
      setFoods(result);
    };

    fetchData();
  }, [name]);

  if (!foods || !foods.foods) return null;
  return (
    <div className="">
      <ul className="list-disc list-outside pl-4">
        {(foods.foods || []).map((r) => {
          const isHighlighted = highlightId === r.fdc_id;
          return (
            <div
              style={{ gridTemplateColumns: "5rem 15rem 5rem 5rem 5rem 5rem" }}
              className={`border ${
                isHighlighted ? "border-red-600 " : "border-indigo-600"
              } ${isHighlighted && "bg-indigo-200"} grid`}
              key={`${name}@${r.fdc_id}`}
            >
              <div className="flex flex-col">
                <Code>{r.fdc_id}</Code>
                <a
                  href={`https://fdc.nal.usda.gov/fdc-app.html#/food-details/${r.fdc_id}/nutrients`}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-sm pr-1"
                >
                  (view)
                </a>
              </div>
              <div className="flex flex-col">
                <div className="">{r.description}</div>{" "}
                <Code>{r.data_type}</Code>
              </div>
              {!!r.branded_info && (
                <div className="italic">{r.branded_info.brand_owner}</div>
              )}
              {/* <div className="flex"> */}
              <div className="flex flex-col">
                <div className="font-bold flex">nutrients:</div>
                <Code>{r.nutrients?.length}</Code>
              </div>
              <div className="flex flex-col">
                <div className="font-bold flex ml-1">nutrition:</div>
                <div>{`${getCalories(r)} kcal/100g`}</div>
              </div>
              {/* </div> */}
              {onLink !== undefined && (
                <ButtonGroup
                  compact
                  buttons={[
                    {
                      onClick: () => {
                        onLink(r.fdc_id);
                      },
                      text: "link",
                      disabled: isHighlighted,
                      IconLeft: PlusCircle,
                    },
                  ]}
                />
              )}
            </div>
          );
        })}
      </ul>
      {/* <Debug data={foods.foods} /> */}
    </div>
  );
};
export default FoodSearch;
