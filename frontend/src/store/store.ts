import { defineStore } from "pinia";
import {
  ImportCsvFile,
  ImportModel,
  ImportSchema,
  ExportModel,
  ExportSchema,
  ProcessCsvDescriptor,
} from "../../wailsjs/go/main/App";
import { entity } from "../../wailsjs/go/models";

export type EntityTypes = "model" | "schema" | "field";
export type DataTypes = entity.DataType;
export type Field = entity.Field;
export type Schema = entity.Schema;
export type Model = entity.Model;
export type HeaderDescriptor = entity.HeaderDescriptor;

export const fieldOptions: DataTypes[] = [
  { name: "text", value: "text" },
  { name: "number", value: "number" },
  { name: "boolean", value: "boolean" },
];

export type SelectedColumn = {
  header: string;
  schemaIdx: number | null;
  fieldIdx: number | null;
};

type StoreState = {
  csv: {
    fileName: string;
    location: string;
    headers: string[];
  } | null;
  model: Model | null;
  csvEditor: {
    curForm: "new" | "edit" | "none";
    selectedColumns: SelectedColumn[];
  };
  modelEditor: {
    curForm: "new" | "edit" | null;
  };
  schemaEditor: {
    curForm: "new" | "edit" | null;
    curSelected: number;
  };
};

export const useStore = defineStore("store", {
  state: (): StoreState => ({
    // csv: null,
    csv: {
      fileName: "test.csv",
      location: "/home/meeps/Documents/test.csv",
      headers: ["First Name", "Last Name", "Age", "Location", "Email", "LinkedIn URL", "Title", "Company"],
    },
    // model: null,
    model: new entity.Model({
      name: "Person",
      type: "model",
      schemas: [
        {
          name: "person base",
          type: "schema",
          fields: [
            { name: "first name", type: "field", dataType: { name: "text (string)", value: "string" } },
            { name: "last name", type: "field", dataType: { name: "text (string)", value: "string" } },
            { name: "email", type: "field", dataType: { name: "text (string)", value: "string" } },
            { name: "linkedin", type: "field", dataType: { name: "text (string)", value: "string" } },
          ],
        },
        {
          name: "location",
          type: "schema",
          fields: [{ name: "name", type: "field", dataType: { name: "text (string)", value: "string" } }],
        },
      ],
      baseSchemaIdx: 0,
    }),
    csvEditor: {
      curForm: "none",
      // selectedColumns: [],
      selectedColumns: [
        { header: "First Name", schemaIdx: 0, fieldIdx: 0 },
        { header: "Last Name", schemaIdx: 0, fieldIdx: 1 },
        { header: "Location", schemaIdx: 1, fieldIdx: 0 },
        { header: "Email", schemaIdx: 0, fieldIdx: 2 },
        { header: "LinkedIn URL", schemaIdx: 0, fieldIdx: 3 },
      ],
    },
    modelEditor: {
      curForm: null,
    },
    schemaEditor: {
      curForm: null,
      curSelected: 0,
    },
  }),
  actions: {
    changeModelEditorForm(val: "new" | "edit" | null) {
      this.modelEditor.curForm = val;
    },
    changeSchemaEditorForm(val: "new" | "edit" | null) {
      this.schemaEditor.curForm = val;
    },
    changeCsvEditorForm(val: "new" | "edit" | "none") {
      this.csvEditor.curForm = val;
    },
    updateLocalModel(model: Model) {
      this.model = model;
      this.schemaEditor.curSelected = model.baseSchemaIdx;
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
    updateSelectedColumns(headers: string[]) {
      const existingColumns: SelectedColumn[] = this.csvEditor.selectedColumns.filter((sc) =>
        headers.includes(sc.header)
      );
      const notExisting = headers.filter((h) => !this.selectedColumnHeaders.includes(h));
      const newColumns: SelectedColumn[] = notExisting.map((h) => {
        return {
          header: h,
          schemaIdx: null,
          fieldIdx: null,
        };
      });
      this.csvEditor.selectedColumns = [...existingColumns, ...newColumns];
    },
    runProcessCsvDescriptor() {
      if (!this.model) return;
      const validDescs = this.csvEditor.selectedColumns.filter((hd) => hd.schemaIdx !== null && hd.fieldIdx !== null);
      ProcessCsvDescriptor(this.model, validDescs as HeaderDescriptor[], this.csv?.location!);
    },
    // Go
    async exportModel() {
      try {
        await ExportModel(this.model!);
      } catch (err) {
        console.log(err);
      }
    },
    async importModel() {
      try {
        const model = await ImportModel();
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
          this.csvEditor.selectedColumns = [];
          this.csvEditor.curForm = "edit";
          this.csv = {
            fileName: csvData.fileName,
            location: csvData.location,
            headers: csvData.headers,
          };
        } else {
          console.log("canceled import");
        }
      } catch (err) {
        console.log(err);
      }
    },
    async exportSchema() {
      try {
        await ExportSchema(this.model!.schemas[this.schemaEditor.curSelected]);
      } catch (err) {
        console.log(err);
      }
    },
    async importSchema() {
      try {
        const schema = await ImportSchema();
        this.model!.schemas.push(schema);
      } catch (err) {
        console.log(err);
      }
    },
  },
  getters: {
    selectedColumnHeaders(state) {
      return state.csvEditor.selectedColumns.map((sc) => sc.header);
    },
  },
  // getters: {
  //   formatModelDisplay(state) {
  //     const display = {};
  //     const base = {};

  //     if (state.model !== null) {
  //       const baseSchemaIdx = state.model.baseSchemaIdx !== null ? state.model.schemas[state.model.baseSchemaIdx] : null;
  //       if (baseSchemaIdx) {
  //         for (let j = 0; j < baseSchemaIdx.fields.length; j++) {
  //           const field = baseSchemaIdx.fields[j];
  //           Object.assign(base, formatFieldJson(field));
  //         }
  //         Object.assign(display, base);
  //       }

  //       for (let i = 0; i < state.model.schemas.length; i++) {
  //         const schema = state.model.schemas[i];
  //         if (state.model.baseSchemaIdx !== i) {
  //           Object.assign(display, formatSchemaJson(schema));
  //         }
  //       }
  //     }

  //     return display;
  //   },
  //   formatSchemaDisplay() {
  //     const display = {};
  //     if (this.model) {
  //       const schema = this.model.schemas[this.schemaEditor.curSelected];
  //       for (let i = 0; i < schema.fields.length; i++) {
  //         const field = schema.fields[i];
  //         Object.assign(display, formatFieldJson(field));
  //       }
  //     }

  //     return display;
  //   },
  // }
});
