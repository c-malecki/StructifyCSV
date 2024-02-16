<script lang="ts" setup>
import { ProcessCsvWithSchema } from "../../../wailsjs/go/main/App";
import { inject } from "vue";
import { CsvFileKey } from "./CsvEditor.types";
import { JsonSchemaKey } from "../SchemaEditor/SchemaEditor.types";
import ModelTree from "./components/ModelTree.vue";

const csvFile = inject(CsvFileKey);
if (!csvFile) {
  throw new Error(`Could not resolve ${CsvFileKey.description}`);
}
const jsonSchema = inject(JsonSchemaKey);
if (!jsonSchema) {
  throw new Error(`Could not resolve ${JsonSchemaKey.description}`);
}

const handleExport = () => {
  ProcessCsvWithSchema(jsonSchema.value).then((res) => {
    console.log(res);
  });
};
</script>

<template>
  <v-card border rounded="0" flat>
    <template #title>
      <h3>{{ csvFile.fileName }}</h3>
      <v-btn @click="handleExport">test</v-btn>
    </template>

    <v-divider />

    <ModelTree />
  </v-card>
</template>
