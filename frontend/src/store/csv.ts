import { defineStore } from "pinia";
import { entity } from "../../wailsjs/go/models";
import {
  ImportCsvFileData,
  ImportJsonSchema,
  ProcessCsvWithSchema,
} from "../../wailsjs/go/main/App";
import { exampleSchema, exampleCsvFile } from "../util/example";
import { fixWailsJsonSchemaImport } from "../util/transform";

type CsvStore = {
  csvFile: entity.CsvFileData | null;
  selectedSchema: entity.JsonSchema | null;
  processingReport: entity.CsvProcessingReport | null;
};

export const useCsvStore = defineStore("csv", {
  state: (): CsvStore => ({
    csvFile: exampleCsvFile,
    selectedSchema: exampleSchema,
    processingReport: null,
  }),
  getters: {
    csvHeaderOpts(state) {
      return state.csvFile
        ? state.csvFile.headers.map((h, i) => {
            return {
              header: h,
              index: i,
            };
          })
        : [];
    },
    canRunCsvProcessor(state) {
      return !!state.selectedSchema && !!state.csvFile;
    },
  },
  actions: {
    importCsvFileData() {
      ImportCsvFileData().then((res) => {
        if (res.headers !== null) {
          this.csvFile = res;
        } else {
          console.log("canceled import");
        }
      });
    },
    selectJsonSchema() {
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
            this.selectedSchema = imported;
          }
        })
        .catch(() => {});
    },
    processCsvWithSelectedSchema() {
      if (!this.csvFile || !this.selectedSchema) return;
      ProcessCsvWithSchema(this.csvFile, this.selectedSchema).then((res) => {
        this.processingReport = res;
      });
    },
  },
});
