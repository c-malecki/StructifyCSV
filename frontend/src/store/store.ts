import { defineStore } from "pinia";
import { ImportCsvData, ProcessCsvDescriptor } from "../../wailsjs/go/main/App";
import { entity } from "../../wailsjs/go/models";

export type DataTypes = { name: string; value: string };
export type Schema = entity.JsonSchema;
export type HeaderDescriptor = entity.HeaderDescriptor;

export const dataTypes: DataTypes[] = [
  { name: "string (text)", value: "string" },
  { name: "float (decimal)", value: "number" },
  { name: "integer", value: "integer" },
  { name: "object ", value: "object" },
  { name: "array", value: "array" },
  { name: "boolean", value: "boolean" },
  { name: "null", value: "null" },
];

export type SelectedColumn = {
  header: string;
  schemaIdx: number | null;
  fieldIdx: number | null;
};

export type SchemaValues = {
  title: string;
  description: string;
  properties: Record<string, any>;
};

type StoreState = {
  csv: {
    fileName: string;
    location: string;
    headers: string[];
  } | null;
  csvEditor: {
    curForm: "new" | "edit" | "none";
    selectedColumns: SelectedColumn[];
  };
  schema: SchemaValues;
};

export const useStore = defineStore("store", {
  state: (): StoreState => ({
    // csv: null,
    csv: {
      fileName: "test.csv",
      location: "/home/meeps/Documents/test.csv",
      headers: ["First Name", "Last Name", "Age", "Location", "Email", "LinkedIn URL", "Title", "Company"],
    },
    csvEditor: {
      curForm: "none",
      selectedColumns: [],
    },
    schema: {
      title: "",
      description: "",
      properties: {
        first_name: "string",
        last_name: "string",
        email: "string",
        location: {
          name: "string",
          longitude: "number",
          latitude: "number",
        },
        user: {
          id: "string",
          member: {
            id: "string",
          },
        },
      },
    },
  }),
  actions: {
    changeCsvEditorForm(val: "new" | "edit" | "none") {
      this.csvEditor.curForm = val;
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
      if (!this.schema) return;
      const validDescs = this.csvEditor.selectedColumns.filter((hd) => hd.schemaIdx !== null && hd.fieldIdx !== null);
      ProcessCsvDescriptor(this.schema, validDescs as HeaderDescriptor[], this.csv?.location!);
    },
    // Go
    async importCsv() {
      try {
        const csvData = await ImportCsvData();
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
  },
  getters: {
    selectedColumnHeaders(state) {
      return state.csvEditor.selectedColumns.map((sc) => sc.header);
    },
  },
});
