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
    if (val instanceof Map) {
      map.set(key, transformForCsvModelMap(val, schemaPath));
    } else {
      map.set(key, {
        schemaPath,
        header: null,
        headerIdx: null,
        dataType: val.type,
      });
    }
  }
  return map;
};
