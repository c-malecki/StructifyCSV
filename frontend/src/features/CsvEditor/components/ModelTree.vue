<script lang="ts" setup>
import { WriteJsonFromCsvModelMap } from "../../../../wailsjs/go/main/App";
import { computed, provide, inject } from "vue";
import { transformCsvModelMaptoObject } from "../../../util/transform";
import {
  CsvSchemaMapKey,
  CsvFileKey,
  JsonSchemaKey,
  HeaderOptsKey,
} from "../../../types/editor.types";
import ModelNode from "./ModelNode.vue";

const csvSchemaMap = inject(CsvSchemaMapKey);
if (!csvSchemaMap) {
  throw new Error(`Could not resolve ${CsvSchemaMapKey.description}`);
}
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
  WriteJsonFromCsvModelMap(transformCsvModelMaptoObject(csvSchemaMap));
};
</script>

<template>
  <div class="pl-4 pt-4">
    <v-btn @click="handleExport">test</v-btn>
    <h3>{{ jsonSchema.title }}</h3>
  </div>

  <div class="schema_tree pa-4">
    <ModelNode
      v-for="(node, i) in csvSchemaMap"
      :key="`1-${i}-${typeof node[1]}-csv`"
      :nodeKey="node[0]"
      :nodeValue="node[1]"
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
