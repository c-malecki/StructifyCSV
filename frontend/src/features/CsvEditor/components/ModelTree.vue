<script lang="ts" setup>
import { WriteJsonFromCsvModelMap } from "../../../../wailsjs/go/main/App";
import { inject } from "vue";
import { convertMaptoObject } from "../../../util/transform";
import { CsvSchemaKey, JsonSchemaKey } from "../../../types/editor.types";
import ModelNode from "./ModelNode.vue";

const csvSchema = inject(CsvSchemaKey);
if (!csvSchema) {
  throw new Error(`Could not resolve ${CsvSchemaKey.description}`);
}
const jsonSchema = inject(JsonSchemaKey);
if (!jsonSchema) {
  throw new Error(`Could not resolve ${JsonSchemaKey.description}`);
}

const handleExport = () => {
  WriteJsonFromCsvModelMap(convertMaptoObject(csvSchema.map));
};
</script>

<template>
  <div class="pl-4 pt-4">
    <v-btn @click="handleExport">test</v-btn>
    <h3>{{ jsonSchema.title }}</h3>
  </div>

  <div class="schema_tree pa-4">
    <ModelNode
      v-for="(node, i) in csvSchema.map"
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
