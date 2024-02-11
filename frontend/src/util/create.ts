import {
  StringProperty,
  IntegerProperty,
  NumberProperty,
  ArrayProperty,
  ObjectProperty,
  BooleanProperty,
  NullProperty,
  type SchemaProperty,
} from "../types/properties.types";

export type StringForm = {
  minLength: string | null;
  maxLength: string | null;
};

export type NumOrIntForm = {
  minimum: string | null;
  maximum: string | null;
};

export type ArrayForm = {
  items: "string" | "number" | "integer" | null;
  minItems: string | null;
  maxItems: string | null;
};

export type PropertyForm = {
  string: StringForm;
  numOrInt: NumOrIntForm;
  array: ArrayForm;
};

export const createSchemaProperty = (
  obj: Record<string, any>
): SchemaProperty | Error => {
  switch (obj.type) {
    case "string":
      return new StringProperty(obj);
    case "integer":
      return new IntegerProperty(obj);
    case "number":
      return new IntegerProperty(obj);
    case "array":
      return new ArrayProperty(obj);
    case "object":
      return new ObjectProperty(obj);
    case "boolean":
      return new BooleanProperty(obj);
    case "null":
      return new NullProperty(obj);
    default:
      return new Error(
        `type "${obj.type} is not an accepted json schema data type`
      );
  }
};
