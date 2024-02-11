import type { CsvSchemaMap } from "../types/editor.types";
import {
  ObjectProperty,
  type PropertiesMap,
  type SchemaProperty,
} from "../types/properties.types";
import { createSchemaProperty } from "../util/create";

export const transformMapToObjectForWails = (
  data: PropertiesMap
): Record<string, any> => {
  let result = {} as Record<string, any>;
  for (let [key, val] of data) {
    if (val.type === "object") {
      result[key] = {
        ...val,
        properties: transformMapToObjectForWails(val.properties),
      };
    } else {
      result[key] = val;
    }
  }
  return result;
};

export const transformWailsObjectToMap = (
  obj: Record<string, any>,
  map: PropertiesMap = new Map<string, SchemaProperty>()
): PropertiesMap => {
  const result = map;
  for (let key of Object.keys(obj)) {
    const value = obj[key];
    if (value.type === "object") {
      const propMap = new Map<string, SchemaProperty>();
      const properties = transformWailsObjectToMap(value.properties, propMap);
      const object = new ObjectProperty({ ...value, properties });
      result.set(key, object);
    } else {
      const property = createSchemaProperty(value);
      if (property instanceof Error) {
        // how to handle error + should classes validate in constructors?
      } else {
        result.set(key, property);
      }
    }
  }
  return result;
};

export const transformForCsvModelMap = (data: PropertiesMap, path = "") => {
  const map: CsvSchemaMap = new Map();
  for (let [key, val] of data) {
    let schemaPath = path.length ? `${path}.${key}` : key;
    if (val.type === "object") {
      map.set(key, transformForCsvModelMap(val.properties, schemaPath));
    } else {
      if (val.type === "array") {
        map.set(key, {
          headerIndexes: [],
          schemaPropertyType: val.type,
        });
      } else {
        map.set(key, {
          headerIndexes: null,
          schemaPropertyType: val.type,
        });
      }
    }
  }
  return map;
};

export const transformCsvModelMaptoObject = (
  data: CsvSchemaMap
): Record<string, any> => {
  let obj = {} as Record<string, any>;
  for (let [key, val] of data) {
    if (val instanceof Map) {
      obj[key] = transformCsvModelMaptoObject(val);
    } else {
      let indexes = [] as number[];
      if (val.schemaPropertyType !== "array" && val.headerIndexes !== null) {
        indexes = [val.headerIndexes as number];
      } else if (val.schemaPropertyType === "array") {
        indexes = [...(val.headerIndexes as number[])];
      }
      obj[key] = { ...val, headerIndexes: indexes };
    }
  }
  return obj;
};
