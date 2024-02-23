import { entity } from "../../../wailsjs/go/models";
import { type PropertyType } from "../../store/schema";

export type ArrayItemType = "string" | "number" | "integer";

export type PropertyNode = [string, entity.PropertySchema];

export type PropertyConstructorFormValues = {
  type: PropertyType;
  minProperties: string | null;
  maxProperties: string | null;
  minItems: string | null;
  maxItems: string | null;
  items: ArrayItemType | null;
  minLength: string | null;
  maxLength: string | null;
  numMinimum: string | null;
  numMaximum: string | null;
  intMinimum: string | null;
  intMaximum: string | null;
};

export type PropertyConstructor = {
  type: PropertyType;
  minProperties?: number;
  maxProperties?: number;
  minItems?: number;
  maxItems?: number;
  items?: { type: ArrayItemType };
  minLength?: number;
  maxLength?: number;
  numMinimum?: number;
  numMaximum?: number;
  intMinimum?: number;
  intMaximum?: number;
};
