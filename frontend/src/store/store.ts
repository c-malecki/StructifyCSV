import { defineStore } from "pinia";
import {
  ImportCsvFile,
  ImportModelJson,
  ImportSchemaJson,
  ExportModelToJson,
  ExportSchemaToJson,
} from "../../wailsjs/go/main/App";
import { main } from "../../wailsjs/go/models";

export type EntityTypes = "model" | "schema" | "field";

export type NewModelValues = {
  name: string;
  baseSchema: number;
};

export type DataTypes = main.DataType;
export type Field = main.Field;
export type Schema = main.Schema;
export type Model = main.Model;

export const fieldOptions: DataTypes[] = [
  { name: "text", value: "text" },
  { name: "number", value: "number" },
  { name: "boolean", value: "boolean" },
];

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
    fileName: string;
    columns: CsvColumn[];
  } | null;
  modelEditor: {
    curForm: "new" | "edit" | null;
  };
  schemaEditor: {
    curForm: "new" | "edit" | null;
    curSelected: number;
  };
  selectedColumns: SelectedColumn[];
  model: Model | null;
  schemas: Schema[];
  selectedSchema: number | null;
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
    csv: null,
    selectedColumns: [],
    modelEditor: {
      curForm: null,
    },
    schemaEditor: {
      curForm: null,
      curSelected: 0,
    },
    model: null,
    schemas: [],
    selectedSchema: null,
  }),
  getters: {
    formatModelDisplay(state) {
      const display = {};
      const base = {};

      if (state.model !== null) {
        const baseSchema = state.model.baseSchema !== null ? state.model.schemas[state.model.baseSchema] : null;
        if (baseSchema) {
          for (let j = 0; j < baseSchema.fields.length; j++) {
            const field = baseSchema.fields[j];
            Object.assign(base, formatFieldJson(field));
          }
          Object.assign(display, base);
        }

        for (let i = 0; i < state.model.schemas.length; i++) {
          const schema = state.model.schemas[i];
          if (state.model.baseSchema !== i) {
            Object.assign(display, formatSchemaJson(schema));
          }
        }
      }

      return display;
    },
    formatSchemaDisplay() {
      const display = {};
      if (this.model && this.selectedSchema !== null) {
        const schema = this.model.schemas[this.selectedSchema];
        for (let i = 0; i < schema.fields.length; i++) {
          const field = schema.fields[i];
          Object.assign(display, formatFieldJson(field));
        }
      }

      return display;
    },
  },
  actions: {
    changeModelEditorForm(val: "new" | "edit" | null) {
      this.modelEditor.curForm = val;
    },
    changeSchemaEditorForm(val: "new" | "edit" | null) {
      this.schemaEditor.curForm = val;
    },
    updateLocalModel(model: Model) {
      this.model = model;
      this.schemaEditor.curSelected = model.baseSchema;
    },
    createLocalSchema(schema: Schema) {
      this.model!.schemas.push(schema);
      this.schemaEditor.curSelected = this.model!.schemas.length - 1;
    },
    updateLocalSchema(schema: Schema) {
      if (this.schemaEditor.curSelected === 0 && this.model!.schemas.length === 1) {
        this.model!.schemas = [schema];
      } else {
        this.model!.schemas[this.schemaEditor.curSelected] = schema;
      }
    },
    toggleSelectColumn(index: number) {
      // Column headers should be unique. Will have to ensure no duplicates when parsing CSV
      const isSelected = !this.csv!.columns[index].isSelected;
      this.csv!.columns[index].isSelected = isSelected;
      const name = this.csv!.columns[index].header;
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
    async exportSchema() {
      try {
        await ExportSchemaToJson(this.model!.schemas[this.selectedSchema!]);
      } catch (err) {
        console.log(err);
      }
    },
    async importSchema() {
      try {
        const schema = await ImportSchemaJson();
        this.model!.schemas.push(schema);
      } catch (err) {
        console.log(err);
      }
    },
    async exportModel() {
      try {
        await ExportModelToJson(this.model!);
      } catch (err) {
        console.log(err);
      }
    },
    async importModel() {
      try {
        const model = await ImportModelJson();
        this.model = model;
        this.modelEditor.curForm = null;
      } catch (err) {
        console.log(err);
      }
    },
    async importCsv() {
      try {
        const csvData = await ImportCsvFile();
        if (csvData.headers !== null) {
          this.selectedColumns = [];
          this.csv = {
            fileName: csvData.fileName,
            columns: csvData.headers.map((header) => {
              return {
                header,
                isSelected: false,
              };
            }),
          };
        } else {
          console.log("canceled import");
        }
      } catch (err) {
        console.log(err);
      }
    },
  },
});
