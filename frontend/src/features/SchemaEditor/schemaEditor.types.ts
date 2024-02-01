export type PropertiesMap = Map<string, string | Map<string, any>>;

export type SchemaValues = {
  title: string;
  description: string;
  properties: PropertiesMap;
};

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
