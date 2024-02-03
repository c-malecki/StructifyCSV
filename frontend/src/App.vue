<script lang="ts" setup>
import {
  ImportSchema,
  ExportSchema,
  ImportCsvData,
} from "../wailsjs/go/main/App";
import { ref, reactive, provide } from "vue";
import { convertMaptoObject, convertObjectToMap } from "./util/transform";
import { exampleSchema, exampleCsvFile, exampleCsvModel } from "./util/example";
import {
  SchemaValuesKey,
  CsvFileKey,
  CsvModelKey,
  type SchemaValues,
  type CsvFile,
  type CsvModel,
  type PropertiesMap,
} from "./types/editor.types";
import ProgramBar from "./ui/ProgramBar.vue";
import CsvEditor from "./features/CsvEditor/CsvEditor.vue";
import SchemaEditor from "./features/SchemaEditor/SchemaEditor.vue";

// make note about mutating props directly in recursive component trees

const programBarRef = ref<typeof ProgramBar | null>(null);
const schemaEditorRef = ref<typeof SchemaEditor | null>(null);

// rename schemaData
const schemaValues = reactive<SchemaValues>(exampleSchema);
const csvFile = reactive<CsvFile>(exampleCsvFile);
const csvModel = reactive<CsvModel>(exampleCsvModel);

provide(CsvFileKey, csvFile);
provide(SchemaValuesKey, schemaValues);
provide(CsvModelKey, csvModel);

const handleCreateNewSchema = () => {
  schemaValues.title = "New Schema";
  schemaValues.description =
    "To change the name and description of this Schema, use the EDIT button to the right. \nTo begin building your Schema, click the ADD button below.";
  schemaValues.properties = new Map() as PropertiesMap;
  programBarRef.value!.menuControl.show = false;
};

const handleImportSchema = () => {
  ImportSchema()
    .then((res) => {
      schemaValues.title = res.title;
      schemaValues.description = res.description;
      schemaValues.properties = convertObjectToMap(res.properties);
      programBarRef.value!.menuControl.show = false;
    })
    .catch(() => {});
};

const handleExportSchema = () => {
  ExportSchema({
    ...schemaValues,
    properties: convertMaptoObject(schemaValues.properties),
  })
    .then(() => {
      programBarRef.value!.menuControl.show = false;
    })
    .catch((err) => {});
};

const handleUpdateSchema = ({
  title,
  description,
}: Omit<SchemaValues, "properties">) => {
  schemaValues.title = title;
  schemaValues.description = description;
};

// const importCsv = async () => {
//   try {
//     const csvFileData = await ImportCsvData();
//     if (csvFileData.headers !== null) {
//       csvData.fileName = csvFileData.fileName;
//       csvData.fileLocation = csvFileData.location;
//       csvData.headers = csvFileData.headers;
//     } else {
//       console.log("canceled import");
//     }
//   } catch (err) {
//     console.log(err);
//   }
// };

// todo: How to handle arrays? Do I need to have support for
// multi-typed properties?
// look at next steps of JSON Schema spec integration
</script>

<template>
  <v-app id="app">
    <v-main>
      <ProgramBar
        @new-schema="handleCreateNewSchema"
        @import-schema="handleImportSchema"
        @export-schema="handleExportSchema"
        ref="programBarRef"
      />
      <v-container fluid class="pa-0">
        <v-row no-gutters>
          <v-col>
            <SchemaEditor
              @update-schema="handleUpdateSchema"
              @close-menu="programBarRef!.menuControl.show = false"
              ref="schemaEditorRef"
            />
          </v-col>
          <v-col>
            <CsvEditor />
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>
