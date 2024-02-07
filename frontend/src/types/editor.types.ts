import { InjectionKey } from "vue";
import {
  type SchemaProperty,
  type SchemaPropertyType,
  type PropertiesMap,
} from "./properties.types";

export type JsonSchema = {
  title: string;
  description: string;
  properties: PropertiesMap;
};

export type SchemaNode = [string, SchemaProperty];

export const JsonSchemaKey: InjectionKey<JsonSchema> = Symbol("JsonSchema");

export type CsvFile = {
  fileName: string;
  fileLocation: string;
};

export const CsvFileKey: InjectionKey<CsvFile> = Symbol("CsvFile");

type PropertyDescriptor = {
  key: string;
  type: SchemaPropertyType;
  path: string;
};

type HeaderDescriptor = {
  isSelected: boolean;
  headerText: string;
  headerIndex: number | undefined;
  propertyDescriptor: PropertyDescriptor | undefined;
};

export type CsvSchemaProperty = {
  schemaPath: string;
  header: string | null;
  headerIdx: number | null;
  dataType: SchemaPropertyType;
};

export type CsvSchemaMapValue =
  | CsvSchemaProperty
  | Map<string, CsvSchemaMapValue>;
export type CsvSchemaMap = Map<string, CsvSchemaMapValue>;

export type CsvSchema = {
  headerDescriptors: HeaderDescriptor[];
  usedHeaderIndexes: number[];
  // should become reference to file path?
  schema: JsonSchema;
  map: CsvSchemaMap;
};

export const CsvSchemaKey: InjectionKey<CsvSchema> = Symbol("CsvSchema");
