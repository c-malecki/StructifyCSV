import { InjectionKey, type ComputedRef } from "vue";
import { entity } from "../../../wailsjs/go/models";

export const CsvFileKey: InjectionKey<entity.CsvFileData> = Symbol("CsvFile");

export const HeaderOptsKey: InjectionKey<
  ComputedRef<{ header: string; index: number }[]>
> = Symbol("HeaderOptsKey");
