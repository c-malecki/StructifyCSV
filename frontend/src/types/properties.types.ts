export type PropertiesMap = Map<string, SchemaProperty>;

export type SchemaPropertyType =
  | "string"
  | "number"
  | "integer"
  | "object"
  | "array"
  | "boolean"
  | "null";

export const schemaPropertyTypes: SchemaPropertyType[] = [
  "string",
  "number",
  "integer",
  "object",
  "array",
  "boolean",
  "null",
];

export type SchemaProperty =
  | StringProperty
  | IntegerProperty
  | NumberProperty
  | ArrayProperty
  | ObjectProperty
  | BooleanProperty
  | NullProperty;

export type StringConstructor = { minLength?: number; maxLength?: number };

export class StringProperty {
  type: "string";
  minLength?: number;
  maxLength?: number;

  constructor({ minLength, maxLength }: StringConstructor = {}) {
    this.type = "string";
    this.minLength = minLength ? minLength : undefined;
    this.maxLength = maxLength ? maxLength : undefined;
  }

  /**
   * getAttributeDisplay
   */
  public getAttributeDisplay() {
    const entries = Object.entries(this);
    return entries.filter(([k, v]) => k !== "type" && v !== undefined);
  }
}

export type NumOrIntConstructor = { minimum?: number; maximum?: number };

export class IntegerProperty {
  type: "integer";
  minimum?: number;
  maximum?: number;

  constructor({ minimum, maximum }: NumOrIntConstructor = {}) {
    this.type = "integer";
    this.minimum = minimum;
    this.maximum = maximum;
  }

  /**
   * getAttributeDisplay
   */
  public getAttributeDisplay() {
    const entries = Object.entries(this);
    return entries.filter(([k, v]) => k !== "type" && v !== undefined);
  }
}

export class NumberProperty {
  type: "number";
  minimum?: number;
  maximum?: number;

  constructor({ minimum, maximum }: NumOrIntConstructor = {}) {
    this.type = "number";
    this.minimum = minimum;
    this.maximum = maximum;
  }

  /**
   * getAttributeDisplay
   */
  public getAttributeDisplay() {
    const entries = Object.entries(this);
    return entries.filter(([k, v]) => k !== "type" && v !== undefined);
  }
}
export type ArrayConstructor = {
  items?: "string" | "number" | "integer";
  minItems?: number;
  maxItems?: number;
};
export class ArrayProperty {
  type: "array";
  items?: { type: "string" | "number" | "integer" };
  minItems?: number;
  maxItems?: number;

  constructor({ items, minItems, maxItems }: ArrayConstructor = {}) {
    this.type = "array";
    this.items = items ? { type: items } : undefined;
    this.minItems = minItems;
    this.maxItems = maxItems;
  }

  /**
   * getAttributeDisplay
   */
  public getAttributeDisplay() {
    const entries = Object.entries(this);
    return entries.filter(
      ([k, v]) => k !== "type" && k !== "items" && v !== undefined
    );
  }
}

export class ObjectProperty {
  type: "object";
  properties: PropertiesMap;
  required?: string[];

  constructor({
    properties,
    required,
  }: {
    properties?: Map<string, SchemaProperty> | null;
    required?: string[] | null;
  } = {}) {
    this.type = "object";
    this.properties = properties
      ? properties
      : new Map<string, SchemaProperty>();
    this.required = required ? required : undefined;
  }

  /**
   * getAttributeDisplay
   */
  public getAttributeDisplay() {
    const entries = Object.entries(this);
    return entries.filter(
      ([k, v]) => k !== "type" && k !== "properties" && v !== undefined
    );
  }
}

export class BooleanProperty {
  type: "boolean";

  constructor(source: any = {}) {
    if ("string" === typeof source) source = JSON.parse(source);
    this.type = "boolean";
  }

  /**
   * getAttributeDisplay
   */
  public getAttributeDisplay() {
    return [] as [string, any][];
  }
}

export class NullProperty {
  type: "null";

  constructor(source: any = {}) {
    if ("string" === typeof source) source = JSON.parse(source);
    this.type = "null";
  }

  /**
   * getAttributeDisplay
   */
  public getAttributeDisplay() {
    return [] as [string, any][];
  }
}
