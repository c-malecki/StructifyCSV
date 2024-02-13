import { entity } from "../../wailsjs/go/models";
import {
  type PropertyConstructorFormValues,
  type SchemaPropertyConstructor,
  type ArrayItemType,
  SchemaPropertyType,
} from "../features/SchemaEditor/SchemaEditor.types";

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
 * When a schema.json file is unmarshalled, Wails returns `Schema` struct.
 * Values unprovided to the struct from the unmarshalled JSON are returned as `null` instead of `undefined`.
 * To align with the Wails generated `models.ts` TypeScript class `Schema`, this converts the `null` form to `undefined`.
 */
export const fixWailSchemaImport = (
  properties: entity.Schema["properties"]
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
 * TypeScript class `Schema`, this converts the `null` values to `undefined`.
 */
export const propertyFormNullToUndefined = (
  formValues: PropertyConstructorFormValues
) => {
  const result = {} as SchemaPropertyConstructor;
  for (let key of Object.keys(formValues)) {
    const value = formValues[key as keyof PropertyConstructorFormValues];
    if (value === null || value === "") {
      result[key as keyof Omit<SchemaPropertyConstructor, "type">] = undefined;
    } else {
      if (key === "type") {
        result[key] = value as SchemaPropertyType;
      } else if (key === "items") {
        result["items"] = { type: value as ArrayItemType };
      } else {
        result[key as keyof Omit<SchemaPropertyConstructor, "items" | "type">] =
          parseInt(value as string);
      }
    }
  }
  return result;
};

export const getSchemaAttributesDisplay = (schema: entity.Schema) => {
  const entries = Object.entries(schema);
  const filterNoDisplay = entries.filter(
    ([k, v]) =>
      v !== undefined &&
      k !== "type" &&
      k !== "properties" &&
      k !== "required" &&
      k !== "items" &&
      k !== "onOf"
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

// export const transformForCsvModelMap = (data: PropertiesMap, path = "") => {
//   const map: CsvSchemaMap = new Map();
//   for (let [key, val] of data) {
//     let schemaPath = path.length ? `${path}.${key}` : key;
//     if (val.type === "object") {
//       map.set(key, transformForCsvModelMap(val.properties, schemaPath));
//     } else {
//       if (val.type === "array") {
//         map.set(key, {
//           headerIndexes: [],
//           schemaPropertyType: val.type,
//         });
//       } else {
//         map.set(key, {
//           headerIndexes: null,
//           schemaPropertyType: val.type,
//         });
//       }
//     }
//   }
//   return map;
// };

// export const transformCsvModelMaptoObject = (
//   data: CsvSchemaMap
// ): Record<string, any> => {
//   let obj = {} as Record<string, any>;
//   for (let [key, val] of data) {
//     if (val instanceof Map) {
//       obj[key] = transformCsvModelMaptoObject(val);
//     } else {
//       let indexes = [] as number[];
//       if (val.schemaPropertyType !== "array" && val.headerIndexes !== null) {
//         indexes = [val.headerIndexes as number];
//       } else if (val.schemaPropertyType === "array") {
//         indexes = [...(val.headerIndexes as number[])];
//       }
//       obj[key] = { ...val, headerIndexes: indexes };
//     }
//   }
//   return obj;
// };
