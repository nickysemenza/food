import React from "react";
import { Link } from "react-router-dom";
import { RecipeDetail, UnitMapping } from "../api/openapi-hooks/api";
import { scaledRound } from "../util";

export interface Props {
  recipe: RecipeDetail;
  multiplier?: number;
}
export const RecipeLink: React.FC<Props> = ({
  recipe: { name, version, is_latest_version, id },
  multiplier,
}) => (
  <div className="flex space-x-0.5">
    <Link to={`/recipe/${id}?multiplier=${multiplier || 1}`} className="link">
      <div
        className={`${is_latest_version ? "text-blue-800" : "text-blue-200"}`}
      >
        {name}
      </div>
    </Link>
    <div className="flex font-mono">v{version}</div>
    {multiplier && <div className="font-mono">@{multiplier}x</div>}
  </div>
);

export const UnitMappingList: React.FC<{ unit_mappings: UnitMapping[] }> = ({
  unit_mappings,
}) => (
  <div className="w-60">
    {unit_mappings &&
      unit_mappings.map((m, x) => (
        <div
          key={x}
          // style={{ gridTemplateColumns: "7fr 1fr 8fr 4fr" }}
          className="flex text-sm text-gray-700"
        >
          <p>
            {scaledRound(m.a.value)} {m.a.unit}
          </p>
          <p className="text-center px-1">=</p>
          <p>
            {scaledRound(m.b.value)} {m.b.unit}
          </p>
          <p className="text-xs pl-1">{m.source}</p>
        </div>
      ))}
  </div>
);
