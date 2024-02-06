import type { SchemaPropertiesMap, CsvModelMap } from "../types/editor.types";

export const convertMaptoObject = (
  data: SchemaPropertiesMap | CsvModelMap
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

export const transformForCsvModelMap = (
  data: SchemaPropertiesMap,
  path = ""
) => {
  const map: CsvModelMap = new Map();
  for (let [key, val] of data) {
    let schemaPath = path.length ? `${path}.${key}` : key;
    if (val instanceof Map) {
      map.set(key, transformForCsvModelMap(val, schemaPath));
    } else {
      map.set(key, {
        schemaPath,
        header: null,
        headerIdx: null,
        dataType: val,
      });
    }
  }
  return map;
};
