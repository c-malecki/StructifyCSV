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

export type PropertiesMapValue =
  | JsonSchemaDataType
  | Map<string, PropertiesMapValue>;
export type SchemaPropertiesMap = Map<string, PropertiesMapValue>;

export type JsonSchema = {
  title: string;
  description: string;
  properties: SchemaPropertiesMap;
};

export const JsonSchemaKey: InjectionKey<JsonSchema> = Symbol("JsonSchema");

export type CsvFile = {
  fileName: string;
  fileLocation: string;
};

export const CsvFileKey: InjectionKey<CsvFile> = Symbol("CsvFile");

type SchemaProperty = {
  key: string;
  value: JsonSchemaDataType;
  path: string;
};

type HeaderDescriptor = {
  isSelected: boolean;
  headerText: string;
  headerIndex: number | undefined;
  schemaProperty: SchemaProperty | undefined;
};

export type CsvModelProperty = {
  schemaPath: string;
  header: string | null;
  headerIdx: number | null;
  dataType: JsonSchemaDataType;
};

export type CsvModelMapValue = CsvModelProperty | Map<string, CsvModelMapValue>;
export type CsvModelMap = Map<string, CsvModelMapValue>;

export type CsvModel = {
  headerDescriptors: HeaderDescriptor[];
  usedHeaderIndexes: number[];
  // should become reference to file path?
  schema: JsonSchema;
  map: CsvModelMap;
};

export const CsvModelKey: InjectionKey<CsvModel> = Symbol("CsvModel");
