import { InjectionKey, type ComputedRef } from "vue";
import { entity } from "../../../wailsjs/go/models";
import { type SchemaPropertyType } from "../SchemaEditor/SchemaEditor.types";

export const CsvFileKey: InjectionKey<entity.CsvFileData> = Symbol("CsvFile");

export type CsvSchemaProperty = {
  headerIndexes: number | number[] | null;
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
