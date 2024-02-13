<script lang="ts" setup>
import {
  ImportJsonSchema,
  ExportJsonSchema,
  ImportCsvFileData,
} from "../wailsjs/go/main/App";
import { ref, reactive, provide } from "vue";
import {
  exampleSchema,
  exampleCsvFile,
  // exampleCsvSchemaMap,
} from "./util/example";
import {
  CsvFileKey,
  CsvSchemaMapKey,
  type CsvSchemaMap,
} from "./features/CsvEditor/CsvEditor.types";
import { fixWailSchemaImport } from "./util/transform";
import { JsonSchemaKey } from "./features/SchemaEditor/SchemaEditor.types";
import TitleBar from "./ui/TitleBar.vue";
import CsvEditor from "./features/CsvEditor/CsvEditor.vue";
import SchemaEditor from "./features/SchemaEditor/SchemaEditor.vue";
import { entity } from "../wailsjs/go/models";

const titleBarRef = ref<typeof TitleBar | null>(null);

const jsonSchema = ref<entity.JsonSchema>(exampleSchema);
const csvFile = reactive<entity.CsvFileData>(exampleCsvFile);
// const csvSchemaMap = reactive<CsvSchemaMap>(exampleCsvSchemaMap);

provide(CsvFileKey, csvFile);
provide(JsonSchemaKey, jsonSchema);
// provide(CsvSchemaMapKey, csvSchemaMap);

const handleCreateNewSchema = () => {
  const schema = new entity.JsonSchema({
    title: "New Schema",
    description:
      "To change the name and description of this Schema, use the EDIT button to the right. \nTo begin building your Schema, click the ADD button below.",
    properties: {},
  });

  jsonSchema.value = schema;
  titleBarRef.value!.menuControl.show = false;
};

const handleImportSchema = () => {
  ImportJsonSchema()
    .then(({ schema, error }) => {
      if (error) {
        console.log(error);
        // show import error somewhere
      }
      if (schema) {
        // when Wails unmarshals the JSON file, values are null instead of undefined
        // but the generated models.ts class uses undefined
        const properties = fixWailSchemaImport(schema.properties);
        jsonSchema.value = { ...schema, properties };
      }
      titleBarRef.value!.menuControl.show = false;
    })
    .catch(() => {});
};

const handleExportSchema = () => {
  ExportJsonSchema(jsonSchema.value)
    .then(() => {
      titleBarRef.value!.menuControl.show = false;
    })
    .catch((err) => {});
};

const handleUpdateSchema = ({
  title,
  description,
}: Omit<entity.JsonSchema, "properties">) => {
  jsonSchema.value.title = title;
  jsonSchema.value.description = description;
};

const handleImportCsv = async () => {
  try {
    const csvFileData = await ImportCsvFileData();
    if (csvFileData.headers !== null) {
      csvFile.fileName = csvFileData.fileName;
      csvFile.location = csvFileData.location;
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
      <TitleBar
        @new-schema="handleCreateNewSchema"
        @import-schema="handleImportSchema"
        @export-schema="handleExportSchema"
        @import-csv="handleImportCsv"
        ref="titleBarRef"
      />
      <v-container fluid class="pa-0">
        <v-row no-gutters>
          <v-col>
            <SchemaEditor
              @update-schema="handleUpdateSchema"
              @close-menu="titleBarRef!.menuControl.show = false"
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
