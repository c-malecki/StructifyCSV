import { InjectionKey } from "vue";

export type JsonSchemaDataType =
  | "string"
  | "number"
  | "integer"
  | "object"
  | "array"
  | "boolean"
  | "null";

export const dataTypeOpts: JsonSchemaDataType[] = [
  "string",
  "number",
  "integer",
  "object",
  "array",
  "boolean",
  "null",
];

export type PropertyMapValue =
  | JsonSchemaDataType
  | Map<string, PropertyMapValue>;
export type PropertiesMap = Map<string, PropertyMapValue>;

export type SchemaValues = {
  title: string;
  description: string;
  properties: PropertiesMap;
};

export const SchemaValuesKey: InjectionKey<SchemaValues> =
  Symbol("SchemaValues");

type SchemaProperty = {
  key: string;
  schemaPath: string;
  value: JsonSchemaDataType;
} | null;

export type CsvHeader = {
  isSelected: boolean;
  header: string;
  schemaProperty: SchemaProperty | null;
};

export type CsvFile = {
  fileName: string;
  fileLocation: string;
};

export const CsvFileKey: InjectionKey<CsvFile> = Symbol("CsvFile");

export type CsvModelProperty = {
  schemaPath: string;
  header: string | null;
  headerIdx: number | null;
  dataType: JsonSchemaDataType;
};

export type CsvModelMapValue = CsvModelProperty | Map<string, CsvModelMapValue>;
export type CsvModelMap = Map<string, CsvModelMapValue>;

export type CsvModel = {
  headers: CsvHeader[];
  map: CsvModelMap;
};

export const CsvModelKey: InjectionKey<CsvModel> = Symbol("CsvModel");
