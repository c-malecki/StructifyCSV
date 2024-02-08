import type { CsvSchemaMap } from "../types/editor.types";
import type { PropertiesMap } from "../types/properties.types";

export const convertMaptoObject = (
  data: PropertiesMap | CsvSchemaMap
): Record<string, any> => {
  let obj = {} as Record<string, any>;
  for (let [key, val] of data) {
    if (val instanceof Map) {
      obj[key] = convertMaptoObject(val);
    } else {
      obj[key] = val;
    }
  }
  return obj;
};

export const convertObjectToMap = (obj: Record<string, any>) => {
  const map = new Map();
  for (let key of Object.keys(obj)) {
    if (obj[key] instanceof Object) {
      map.set(key, convertObjectToMap(obj[key]));
    } else {
      map.set(key, obj[key]);
    }
  }
  return map;
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
