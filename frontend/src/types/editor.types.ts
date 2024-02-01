import { InjectionKey } from "vue";

export type PropertiesMap = Map<string, string | Map<string, any>>;

export type SchemaValues = {
  title: string;
  description: string;
  properties: PropertiesMap;
};

export const SchemaValuesKey: InjectionKey<SchemaValues> =
  Symbol("SchemaValues");

type DataTypeName =
  | "string (text)"
  | "float (decimal)"
  | "integer"
  | "object"
  | "array"
  | "boolean"
  | "null";

type DataTypeValue =
  | "string"
  | "number"
  | "integer"
  | "object"
  | "array"
  | "boolean"
  | "null";

export type DataTypeOpts = { name: DataTypeName; value: DataTypeValue };

export const dataTypeOpts: DataTypeOpts[] = [
  { name: "string (text)", value: "string" },
  { name: "float (decimal)", value: "number" },
  { name: "integer", value: "integer" },
  { name: "object", value: "object" },
  { name: "array", value: "array" },
  { name: "boolean", value: "boolean" },
  { name: "null", value: "null" },
];

export type ColumnHeader = {
  name: string;
  selected: boolean;
};

export type CsvFileData = {
  fileName: string;
  fileLocation: string;
  headers: ColumnHeader[];
};
