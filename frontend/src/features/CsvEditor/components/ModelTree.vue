<script lang="ts" setup>
import { inject } from "vue";
import { CsvModelKey } from "../../../types/editor.types";
import ModelNode from "./ModelNode.vue";

const csvModel = inject(CsvModelKey);
if (!csvModel) {
  throw new Error(`Could not resolve ${CsvModelKey.description}`);
}
</script>

<template>
  <div class="schema_tree pa-4">
    <ModelNode
      v-for="(node, i) in csvModel.map"
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
