import type { DataTypes, MapValue } from "./schemaEditor.types";

export const dataTypes: DataTypes[] = [
  { name: "string (text)", value: "string" },
  { name: "float (decimal)", value: "number" },
  { name: "integer", value: "integer" },
  { name: "object", value: "object" },
  { name: "array", value: "array" },
  { name: "boolean", value: "boolean" },
  { name: "null", value: "null" },
];

const hoverColorScheme = {
  1: { hover: "blue-lighten-5", font: "black" },
  2: { hover: "blue-lighten-4", font: "black" },
  3: { hover: "blue-lighten-3", font: "black" },
  4: { hover: "blue-lighten-2", font: "black" },
  5: { hover: "blue-lighten-1", font: "white" },
  6: { hover: "blue-darken-1", font: "white" },
  7: { hover: "blue-darken-2", font: "white" },
  8: { hover: "blue-darken-3", font: "white" },
  9: { hover: "blue-darken-4", font: "white" },
};

export const getHoverColorScheme = (schemaNodeLevel: number) => {
  return hoverColorScheme[schemaNodeLevel as keyof typeof hoverColorScheme];
};

export const getLeftIndent = (schemaNodeLevel: number) => {
  return `margin-left: ${schemaNodeLevel * 16}px`;
};

export const convertMaptoObject = (data: MapValue): Record<string, any> => {
  let obj = {} as Record<string, any>;
  for (let [k, v] of data) {
    if (v instanceof Map) {
      obj[k] = convertMaptoObject(v);
    } else {
      obj[k] = v;
    }
  }
  return obj;
};

export const convertObjectToMap = (obj: Record<string, any>) => {
  let map = new Map();
  for (let key of Object.keys(obj)) {
    if (obj[key] instanceof Object) {
      map.set(key, convertObjectToMap(obj[key]));
    } else {
      map.set(key, obj[key]);
    }
  }
  return map;
};

const address = new Map<string, string | Map<string, any>>([
  ["line_one", "string"],
  ["line_two", "string"],
  ["city", "string"],
  ["state", "string"],
  ["country", "string"],
  ["postal_code", "string"],
]);

const bids = new Map<string, string | Map<string, any>>([
  ["initial_price", "number"],
  ["last_bid", "null"],
  ["total_bids", "integer"],
]);

const seller = new Map<string, string | Map<string, any>>([
  ["name", "string"],
  ["address", address],
]);

export const exampleSchema = new Map<string, string | Map<string, any>>([
  ["id", "integer"],
  ["name", "string"],
  ["description", "string"],
  ["featured", "boolean"],
  ["images", "array"],
  ["bids", bids],
  ["seller", seller],
]);
