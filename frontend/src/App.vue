<script lang="ts" setup>
import {
  ImportJsonSchema,
  ExportJsonSchema,
  ImportCsvData,
} from "../wailsjs/go/main/App";
import { ref, reactive, provide } from "vue";
import { convertMaptoObject, convertObjectToMap } from "./util/transform";
import {
  exampleSchema,
  exampleCsvFile,
  exampleCsvSchemaMap,
} from "./util/example";
import {
  JsonSchemaKey,
  CsvFileKey,
  CsvSchemaMapKey,
  type JsonSchema,
  type CsvFile,
  type CsvSchemaMap,
} from "./types/editor.types";
import { type PropertiesMap } from "./types/properties.types";
import ProgramBar from "./ui/ProgramBar.vue";
import CsvEditor from "./features/CsvEditor/CsvEditor.vue";
import SchemaEditor from "./features/SchemaEditor/SchemaEditor.vue";

// make note about mutating props directly in recursive component trees

const programBarRef = ref<typeof ProgramBar | null>(null);
const schemaEditorRef = ref<typeof SchemaEditor | null>(null);

const jsonSchema = reactive<JsonSchema>(exampleSchema);
const csvFile = reactive<CsvFile>(exampleCsvFile);
const csvSchemaMap = reactive<CsvSchemaMap>(exampleCsvSchemaMap);

provide(CsvFileKey, csvFile);
provide(JsonSchemaKey, jsonSchema);
provide(CsvSchemaMapKey, csvSchemaMap);

const handleCreateNewSchema = () => {
  jsonSchema.title = "New Schema";
  jsonSchema.description =
    "To change the name and description of this Schema, use the EDIT button to the right. \nTo begin building your Schema, click the ADD button below.";
  jsonSchema.properties = new Map() as PropertiesMap;
  programBarRef.value!.menuControl.show = false;
};

const handleImportSchema = () => {
  ImportJsonSchema()
    .then((res) => {
      jsonSchema.title = res.title;
      jsonSchema.description = res.description;
      jsonSchema.properties = convertObjectToMap(res.properties);
      programBarRef.value!.menuControl.show = false;
    })
    .catch(() => {});
};

const handleExportSchema = () => {
  ExportJsonSchema({
    ...jsonSchema,
    properties: convertMaptoObject(jsonSchema.properties),
  })
    .then(() => {
      programBarRef.value!.menuControl.show = false;
    })
    .catch((err) => {});
};

const handleUpdateSchema = ({
  title,
  description,
}: Omit<JsonSchema, "properties">) => {
  jsonSchema.title = title;
  jsonSchema.description = description;
};

const handleImportCsv = async () => {
  try {
    const csvFileData = await ImportCsvData();
    if (csvFileData.headers !== null) {
      csvFile.fileName = csvFileData.fileName;
      csvFile.fileLocation = csvFileData.location;
      csvFile.headers = csvFileData.headers;
    } else {
      console.log("canceled import");
    }
  } catch (err) {
    console.log(err);
  }
};
</script>

<template>
  <v-app id="app">
    <v-main>
      <ProgramBar
        @new-schema="handleCreateNewSchema"
        @import-schema="handleImportSchema"
        @export-schema="handleExportSchema"
        @import-csv="handleImportCsv"
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
