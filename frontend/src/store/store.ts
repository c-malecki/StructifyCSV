import { defineStore } from "pinia";
import { ImportCsv, WriteModel, WriteSchema } from "../../wailsjs/go/main/App";
import { main } from "../../wailsjs/go/models";

export type EntityTypes = "model" | "schema" | "field";

export type DataTypes = main.DataType;
export type Field = main.Field;
export type Schema = main.Schema;
export type Model = main.Model;

export type CsvColumn = {
  header: string;
  isSelected: boolean;
};

export type SelectedColumn = {
  name: string;
  schema: number | null;
  field: number | null;
};

type StoreState = {
  csv: {
    fileName: string | null;
    columns: CsvColumn[];
  };
  selectedColumns: SelectedColumn[];
  model: Model;
  selectedSchema: number;
};

const formatFieldJson = (field: Field, acc?: Record<string, any>) => {
  const obj = acc ?? {};
  return Object.assign(obj, {
    [field.name]: field.dataType.value,
  });
};

const formatSchemaJson = (schema: Schema, acc?: Record<string, any>) => {
  const obj = acc ?? {};
  const reformatName = schema.name.replace(" ", "_").toLowerCase();
  Object.assign(obj, {
    [reformatName]: {},
  });
  for (let i = 0; i < schema.fields.length; i++) {
    const field = schema.fields[i];
    const fieldObj = formatFieldJson(field);
    Object.assign(obj[reformatName], fieldObj);
  }
  return obj;
};

export const useStore = defineStore("store", {
  state: (): StoreState => ({
    csv: {
      fileName: "test.csv",
      columns: [
        { header: "First Name", isSelected: true },
        { header: "Last Name", isSelected: true },
        { header: "Email", isSelected: true },
        { header: "Place of Work", isSelected: true },
        { header: "Age", isSelected: false },
      ],
    },
    selectedColumns: [
      {
        name: "First Name",
        schema: 0,
        field: 1,
      },
      {
        name: "Last Name",
        schema: 0,
        field: 1,
      },
      {
        name: "Email",
        schema: 1,
        field: 0,
      },
      {
        name: "Place of Work",
        schema: 2,
        field: 0,
      },
    ],
    model: new main.Model({
      name: "Example Model",
      type: "model",
      baseSchema: 0,
      schemas: [
        {
          name: "Person",
          type: "schema",
          fields: [
            { name: "firstName", type: "field", dataType: { name: "text (string)", value: "string" } },
            { name: "lastName", type: "field", dataType: { name: "text (string)", value: "string" } },
          ],
        },
        {
          name: "Person",
          type: "schema",
          fields: [
            { name: "firstName", type: "field", dataType: { name: "text (string)", value: "string" } },
            { name: "lastName", type: "field", dataType: { name: "text (string)", value: "string" } },
          ],
        },
        {
          name: "User",
          type: "schema",
          fields: [{ name: "email", type: "field", dataType: { name: "text (string)", value: "string" } }],
        },
        {
          name: "Organization",
          type: "schema",
          fields: [{ name: "name", type: "field", dataType: { name: "text (string)", value: "string" } }],
        },
      ],
    }),
    selectedSchema: 0,
  }),
  getters: {
    getCurModelLabel(state) {
      return state.model.name;
    },
    getCurModelSchemas(state) {
      return state.model.schemas;
    },
    getSelectedSchemaLabel(state) {
      return state.model.schemas[state.selectedSchema].name;
    },
    getSelectedSchemaFields(state) {
      return state.model.schemas[state.selectedSchema].fields;
    },
    formatModelDisplay(state) {
      const display = {};
      const base = {};
      const baseSchema = state.model.schemas[state.model.baseSchema];
      for (let j = 0; j < baseSchema.fields.length; j++) {
        const field = baseSchema.fields[j];
        Object.assign(base, formatFieldJson(field));
      }
      Object.assign(display, base);

      for (let i = 0; i < this.getCurModelSchemas.length; i++) {
        const schema = this.getCurModelSchemas[i];
        if (state.model.baseSchema !== i) {
          Object.assign(display, formatSchemaJson(schema));
        }
      }
      return display;
    },
    formatSchemaDisplay() {
      const display = {};
      for (let i = 0; i < this.getSelectedSchemaFields.length; i++) {
        const field = this.getSelectedSchemaFields[i];
        Object.assign(display, formatFieldJson(field));
      }
      return display;
    },
  },
  actions: {
    async saveSchema() {
      try {
        await WriteSchema(this.model.schemas[this.selectedSchema]);
      } catch (err) {
        console.log(err);
      }
    },
    async saveModel() {
      try {
        await WriteModel(this.model);
      } catch (err) {
        console.log(err);
      }
    },
    async setColumns() {
      try {
        const file = await ImportCsv();
        this.csv.columns = file.headers.map((header) => {
          return {
            header,
            isSelected: false,
          };
        });
      } catch (err) {
        console.log(err);
      }
    },
    toggleSelectColumn(index: number) {
      // Column headers should be unique. Will have to ensure no duplicates when parsing CSV
      const isSelected = !this.csv.columns[index].isSelected;
      this.csv.columns[index].isSelected = isSelected;
      const name = this.csv.columns[index].header;
      if (isSelected) {
        this.selectedColumns.push({
          name,
          schema: null,
          field: null,
        });
      } else {
        this.selectedColumns = this.selectedColumns.filter((col) => col.name !== name);
      }
    },
    createModel(name: string) {
      this.model = new main.Model({ name, type: "model", baseSchema: 0, schemas: [] });
    },
    createSchema(name: string) {
      this.model?.schemas.push(new main.Schema({ name, type: "schema", fields: [] }));
    },
    addField() {
      this.model?.schemas[this.selectedSchema!].fields.push(
        new main.Field({
          name: "",
          type: "field",
          dataType: { name: "text (string)", value: "string" },
        })
      );
    },
  },
  // selectSchema(index: number) {
  //   this.selectedSchema = this.model!.schemas[index];
  // },
  // createSchema(schemaName: string) {
  //   this.schema = { name: schemaName, fields: [] };
  // },
});
