<script lang="ts" setup>
import { ImportSchema, ExportSchema } from "../wailsjs/go/main/App";
import { ref, reactive, provide } from "vue";
import { convertMaptoObject, convertObjectToMap } from "./util/transform";
import { exampleSchema } from "./util/example";
import { SchemaValuesKey, type SchemaValues } from "./types/editor.types";
import ProgramBar from "./ui/ProgramBar.vue";
import CsvEditor from "./features/CsvEditor/CsvEditor.vue";
import SchemaEditor from "./features/SchemaEditor/SchemaEditor.vue";

const programBarRef = ref<typeof ProgramBar | null>(null);
const schemaEditorRef = ref<typeof SchemaEditor | null>(null);

const schemaValues = reactive<SchemaValues>(exampleSchema);
provide(SchemaValuesKey, schemaValues);

const handleCreateNewSchema = () => {
  schemaValues.title = "New Schema";
  schemaValues.description =
    "To change the name and description of this Schema, use the EDIT button to the right. \nTo begin building your Schema, click the ADD button below.";
  schemaValues.properties = new Map<string, string | Map<string, any>>();
  programBarRef.value!.showMenu = false;
};
const handleImportSchema = () => {
  ImportSchema()
    .then((res) => {
      schemaValues.title = res.title;
      schemaValues.description = res.description;
      schemaValues.properties = convertObjectToMap(res.properties);
      programBarRef.value!.showMenu = false;
    })
    .catch(() => {});
};
const handleExportSchema = () => {
  ExportSchema({
    ...schemaValues,
    properties: convertMaptoObject(schemaValues.properties),
  })
    .then(() => {
      programBarRef.value!.showMenu = false;
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

// todo: How to handle arrays? Do I need to have support for
// multi-typed properties?
// look at next steps of JSON Schema spec integration
</script>

<template>
  <v-app id="app">
    <v-main>
      <ProgramBar
        @new="handleCreateNewSchema"
        @import="handleImportSchema"
        @export="handleExportSchema"
        ref="programBarRef"
      />

      <SchemaEditor
        @update-schema="handleUpdateSchema"
        @close-menu="programBarRef!.showMenu = false"
        ref="schemaEditorRef"
      />

      <CsvEditor />
    </v-main>
  </v-app>
</template>
