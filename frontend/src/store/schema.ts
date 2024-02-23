import { defineStore } from "pinia";
import { entity } from "../../wailsjs/go/models";
import { ImportJsonSchema, ExportJsonSchema } from "../../wailsjs/go/main/App";
// import { testSchema } from "../util/testData";
import { fixWailsJsonSchemaImport } from "../util/transform";

export type AppTab = "schema" | "csv";

export type PropertyType =
  | "string"
  | "number"
  | "integer"
  | "object"
  | "array"
  | "boolean"
  | "null";

type SchemaStore = {
  jsonSchema: entity.JsonSchema;
  propertyTypes: PropertyType[];
};

export const useSchemaStore = defineStore("schema", {
  state: (): SchemaStore => ({
    jsonSchema: new entity.JsonSchema({
      title: "New Schema",
      description:
        "To change the name and description of this Schema, use the EDIT button. \nTo begin building your Schema, click the ADD button below.",
      properties: {},
    }),
    propertyTypes: [
      "string",
      "number",
      "integer",
      "object",
      "array",
      "boolean",
      "null",
    ],
  }),
  actions: {
    newJsonSchema() {
      this.jsonSchema = new entity.JsonSchema({
        title: "New Schema",
        description:
          "To change the name and description of this Schema, use the EDIT button. \nTo begin building your Schema, click the ADD button below.",
        properties: {},
      });
    },
    exportJsonSchema() {
      if (!this.jsonSchema) return;
      ExportJsonSchema(this.jsonSchema)
        .then(() => {
          // emit("changeTab", "schema");
        })
        .catch((err) => {});
    },
    importJsonSchema() {
      ImportJsonSchema()
        .then(({ schema, error }) => {
          if (error) {
            console.log(error);
            // show import error somewhere
          }
          if (schema) {
            const imported = {
              ...schema,
              // when Wails unmarshals the JSON file, values are null instead of undefined
              // but the generated models.ts class uses undefined
              properties: fixWailsJsonSchemaImport(schema.properties),
            };
            this.jsonSchema = imported;
          }
        })
        .catch(() => {});
    },
    updateJsonSchemaNameDesc({
      title,
      description,
    }: Omit<entity.JsonSchema, "properties" | "required" | "type">) {
      if (!this.jsonSchema) return;
      this.jsonSchema.title = title;
      this.jsonSchema.description = description;
    },
  },
});
