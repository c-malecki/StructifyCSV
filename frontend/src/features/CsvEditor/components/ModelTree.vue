<script lang="ts" setup>
import { computed, provide, inject } from "vue";
import { CsvFileKey, HeaderOptsKey } from "../CsvEditor.types";
import { JsonSchemaKey } from "../../SchemaEditor/SchemaEditor.types";
import ModelNode from "./ModelNode.vue";

const csvFile = inject(CsvFileKey);
if (!csvFile) {
  throw new Error(`Could not resolve ${CsvFileKey.description}`);
}
const selectedSchema = inject(JsonSchemaKey);
if (!selectedSchema) {
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
</script>

<template>
  <div class="pl-4 pt-4">
    <h3>{{ selectedSchema.title }}</h3>
  </div>

  <div class="schema_tree pa-4">
    <ModelNode
      v-for="([k, v], i) in Object.entries(selectedSchema.properties)"
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
