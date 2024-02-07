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
