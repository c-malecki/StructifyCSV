import { InjectionKey } from "vue";
import { entity } from "../../wailsjs/go/models";

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

export type CsvModelProperty = {
  schemaPath: string;
  header: string | null;
  headerIdx: number | null;
  dataType: JsonSchemaDataType;
};

export type CsvModelMapValue = CsvModelProperty | Map<string, CsvModelMapValue>;
export type CsvModelMap = Map<string, CsvModelMapValue>;

export type CsvModel = entity.CsvModel;

export const CsvModelKey: InjectionKey<CsvModel> = Symbol("CsvModel");
