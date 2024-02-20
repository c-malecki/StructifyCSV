import { entity } from "../../wailsjs/go/models";
import {
  type PropertyConstructorFormValues,
  type PropertyConstructor,
  type ArrayItemType,
} from "../features/SchemaEditor/SchemaEditor.types";
import { type PropertyType } from "../store/schema";

const nullToUndefined = (properties: Record<string, any>) => {
  const result = {} as Record<string, any>;
  const entries = Object.entries(properties);
  for (let [k, v] of entries) {
    if (v === null) {
      result[k] = undefined;
    } else if (k === "properties") {
      result[k] = nullToUndefined(v);
    } else if (k === "items") {
      result[k] = v;
    } else if (k === "required") {
      result[k] = v;
    } else if (v instanceof Object) {
      result[k] = nullToUndefined(v);
    } else {
      result[k] = v;
    }
  }
  return result;
};

/**
 * When a "_.schema.json" file is unmarshalled, Wails returns `PropertySchema` struct.
 * Values unprovided to the struct from the unmarshalled JSON are returned as `null` instead of `undefined`.
 * To align with the Wails generated `models.ts` TypeScript class `PropertySchema`, this converts the `null` form to `undefined`.
 */
export const fixWailsJsonSchemaImport = (
  properties: entity.PropertySchema["properties"]
) => {
  const result = {} as Record<string, any>;
  const entries = Object.entries(properties);
  for (let [k, v] of entries) {
    result[k] = nullToUndefined(v);
  }
  return result;
};

/**
 * Vuetify inputs use `null` as a default value. To align with the Wails generated `models.ts`
 * TypeScript class `PropertySchema`, this converts the `null` values to `undefined`.
 */
export const propertyFormNullToUndefined = (
  formValues: PropertyConstructorFormValues
) => {
  const result = {} as PropertyConstructor;
  for (let key of Object.keys(formValues)) {
    const value = formValues[key as keyof PropertyConstructorFormValues];
    if (value === null || value === "") {
      result[key as keyof Omit<PropertyConstructor, "type">] = undefined;
    } else {
      if (key === "type") {
        result[key] = value as PropertyType;
      } else if (key === "items") {
        result["items"] = { type: value as ArrayItemType };
      } else {
        result[key as keyof Omit<PropertyConstructor, "items" | "type">] =
          parseInt(value as string);
      }
    }
  }
  return result;
};

export const getPropertyAttributesDisplay = (
  property: entity.PropertySchema
) => {
  const entries = Object.entries(property);
  const filterNoDisplay = entries.filter(
    ([k, v]) =>
      v !== undefined &&
      k !== "type" &&
      k !== "properties" &&
      k !== "required" &&
      k !== "items" &&
      k !== "onOf" &&
      k !== "csvHeaderIndex"
  );
  return filterNoDisplay.map(([k, v]) => {
    if (k.includes("numM")) {
      return [k.replace("num", ""), v];
    }
    if (k.includes("intM")) {
      return [k.replace("int", ""), v];
    }
    if (k.includes("min")) {
      const reformat = k.replace("min", "Min");
      const str1 = reformat.substring(0, 3);
      const str2 = reformat.substring(3);
      return [`${str1} ${str2}`, v];
    }
    if (k.includes("max")) {
      const reformat = k.replace("max", "Max");
      const str1 = reformat.substring(0, 3);
      const str2 = reformat.substring(3);
      return [`${str1} ${str2}`, v];
    }

    return [k, v];
  });
};
