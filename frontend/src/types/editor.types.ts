import { InjectionKey, type ComputedRef } from "vue";
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
  headers: string[];
};

export const CsvFileKey: InjectionKey<CsvFile> = Symbol("CsvFile");

// type PropertyDescriptor = {
//   key: string;
//   type: SchemaPropertyType;
//   path: string;
// };

// type HeaderDescriptor = {
//   isSelected: boolean;
//   headerText: string;
//   headerIndex: number | undefined;
//   propertyDescriptor: PropertyDescriptor | undefined;
// };

export type CsvSchemaProperty = {
  headerIndexes: number | number[] | null;
  // headerText: string | string[] | null;
  // schemaPath: string;
  schemaPropertyType: SchemaPropertyType;
};

export type CsvSchemaMapValue =
  | CsvSchemaProperty
  | Map<string, CsvSchemaMapValue>;
export type CsvSchemaMap = Map<string, CsvSchemaMapValue>;

export const CsvSchemaMapKey: InjectionKey<CsvSchemaMap> =
  Symbol("CsvSchemaMap");

export const HeaderOptsKey: InjectionKey<
  ComputedRef<{ header: string; index: number }[]>
> = Symbol("HeaderOptsKey");
