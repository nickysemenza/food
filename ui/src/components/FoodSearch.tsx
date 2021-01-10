import React, { useEffect, useState } from "react";
import { PlusCircle } from "react-feather";
import { PaginatedFoods, FoodApi } from "../api/openapi-fetch";
import { getOpenapiFetchConfig } from "../config";
import { Code } from "../util";
import { ButtonGroup } from "./Button";

const FoodSearch: React.FC<{
  name: string;
  highlightId?: number;
  onLink?: (fdcId: number) => void;
}> = ({ name, highlightId, onLink }) => {
  const [foods, setFoods] = useState<PaginatedFoods>();

  useEffect(() => {
    const fetchData = async () => {
      const bar = new FoodApi(getOpenapiFetchConfig());
      const result = await bar.searchFoods({ name, limit: 5 });
      setFoods(result);
    };

    fetchData();
  }, [name]);

  if (!foods || !foods.foods) return null;
  return (
    <div className="">
      <ul className="list-disc list-outside pl-4">
        {(foods.foods || []).map((r) => {
          const isHighlighted = highlightId === r.fdcId;
          return (
            <li
              className={`border ${
                isHighlighted ? "border-red-600 " : "border-indigo-600"
              } ${isHighlighted && "bg-indigo-200"} flex`}
              key={`${name}@${r.fdcId}`}
            >
              <Code>{r.fdcId}</Code>
              <div className="flex">
                <div className="">{r.description}</div>{" "}
                <Code>{r.dataType}</Code>
              </div>
              {!!r.brandedInfo && (
                <div className="italic">{r.brandedInfo.brandOwner}</div>
              )}
              {/* <div className="flex"> */}
              <div className="font-bold flex">portions:</div>
              <Code>{r.portions?.length}</Code>
              <div className="font-bold flex ml-1">nutrition:</div>
              <Code>{r.nutrients.length}</Code>
              {/* </div> */}
              <div></div>
              {onLink !== undefined && (
                <ButtonGroup
                  compact
                  buttons={[
                    {
                      onClick: () => {
                        onLink(r.fdcId);
                      },
                      text: "link",
                      disabled: isHighlighted,
                      IconLeft: PlusCircle,
                    },
                  ]}
                />
              )}
            </li>
          );
        })}
      </ul>
      {/* <Debug data={foods.foods} /> */}
    </div>
  );
};
export default FoodSearch;
