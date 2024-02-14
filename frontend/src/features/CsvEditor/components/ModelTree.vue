<script lang="ts" setup>
import { ProcessCsvWithSchema } from "../../../../wailsjs/go/main/App";
import { computed, provide, inject } from "vue";
import { CsvFileKey, HeaderOptsKey } from "../CsvEditor.types";
import { JsonSchemaKey } from "../../SchemaEditor/SchemaEditor.types";
import ModelNode from "./ModelNode.vue";

const csvFile = inject(CsvFileKey);
if (!csvFile) {
  throw new Error(`Could not resolve ${CsvFileKey.description}`);
}
const jsonSchema = inject(JsonSchemaKey);
if (!jsonSchema) {
  throw new Error(`Could not resolve ${JsonSchemaKey.description}`);
}

const headerOpts = computed(() =>
  csvFile.headers.map((h, i) => {
    return {
      header: h,
      index: i,
    };
  })
);

provide(HeaderOptsKey, headerOpts);

const handleExport = () => {
  ProcessCsvWithSchema(jsonSchema.value);
};
</script>

<template>
  <div class="pl-4 pt-4">
    <v-btn @click="handleExport">test</v-btn>
    <h3>{{ jsonSchema.title }}</h3>
  </div>

  <div class="schema_tree pa-4">
    <ModelNode
      v-for="([k, v], i) in Object.entries(jsonSchema.properties)"
      :key="`1-csv-${k}`"
      :node="[k, v]"
      :level="1"
    />
  </div>
</template>

<style scoped>
.schema_tree:before {
  content: "{";
}
.schema_tree:after {
  content: "}";
}
</style>
