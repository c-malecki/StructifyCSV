export type JsonSchemaDataType =
  | "string"
  | "number"
  | "integer"
  | "object"
  | "array"
  | "boolean"
  | "null";

export const dataTypeOpts: JsonSchemaDataType[] = [
  "string",
  "number",
  "integer",
  "object",
  "array",
  "boolean",
  "null",
];

export type SchemaPropertyTypes =
  | StringProperty
  | IntegerProperty
  | NumberProperty
  | ArrayProperty
  | ObjectProperty
  | BooleanProperty
  | NullProperty;

export class StringProperty {
  type: "string";
  minLength?: number;
  maxLength?: number;

  static createFrom(source: any = {}) {
    return new StringProperty(source);
  }

  constructor(source: any = {}) {
    if ("string" === typeof source) source = JSON.parse(source);
    this.type = "string";
    this.minLength = source["minLength"];
    this.maxLength = source["maxLength"];
  }
}

export class IntegerProperty {
  type: "integer";
  minimum?: number;
  maximum?: number;

  static createFrom(source: any = {}) {
    return new IntegerProperty(source);
  }

  constructor(source: any = {}) {
    if ("string" === typeof source) source = JSON.parse(source);
    this.type = "integer";
    this.minimum = source["minimum"];
    this.maximum = source["maximum"];
  }
}

export class NumberProperty {
  type: "number";
  minimum?: number;
  maximum?: number;

  static createFrom(source: any = {}) {
    return new NumberProperty(source);
  }

  constructor(source: any = {}) {
    if ("string" === typeof source) source = JSON.parse(source);
    this.type = "number";
    this.minimum = source["minimum"];
    this.maximum = source["maximum"];
  }
}

export class ArrayProperty {
  type: "array";
  items: { type: JsonSchemaDataType };
  minItems?: number;
  maxItems?: number;

  static createFrom(source: any = {}) {
    return new ArrayProperty(source);
  }

  constructor(source: any = {}) {
    if ("string" === typeof source) source = JSON.parse(source);
    this.type = "array";
    this.items = source["items"];
    this.minItems = source["minItems"];
    this.maxItems = source["maxItems"];
  }
}

export class ObjectProperty {
  type: "object";
  properties: Record<string, { type: JsonSchemaDataType } | ObjectProperty>;
  required?: string[];

  static createFrom(source: any = {}) {
    return new ObjectProperty(source);
  }

  constructor(source: any = {}) {
    if ("string" === typeof source) source = JSON.parse(source);
    this.type = "object";
    this.properties = source["properties"];
    this.required = source["required"];
  }
}

export class BooleanProperty {
  type: "boolean";

  static createFrom(source: any = {}) {
    return new ObjectProperty(source);
  }

  constructor(source: any = {}) {
    if ("string" === typeof source) source = JSON.parse(source);
    this.type = "boolean";
  }
}

export class NullProperty {
  type: "null";

  static createFrom(source: any = {}) {
    return new ObjectProperty(source);
  }

  constructor(source: any = {}) {
    if ("string" === typeof source) source = JSON.parse(source);
    this.type = "null";
  }
}
