import type { PropertiesMap } from "../types/editor.types";

export const convertMaptoObject = (
  data: PropertiesMap
): Record<string, any> => {
  let obj = {} as Record<string, any>;
  for (let [k, v] of data) {
    if (v instanceof Map) {
      obj[k] = convertMaptoObject(v);
    } else {
      obj[k] = v;
    }
  }
  return obj;
};

export const convertObjectToMap = (obj: Record<string, any>) => {
  let map = new Map();
  for (let key of Object.keys(obj)) {
    if (obj[key] instanceof Object) {
      map.set(key, convertObjectToMap(obj[key]));
    } else {
      map.set(key, obj[key]);
    }
  }
  return map;
};
