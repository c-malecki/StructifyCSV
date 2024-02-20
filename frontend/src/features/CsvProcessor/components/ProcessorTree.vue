<script lang="ts" setup>
import { useCsvStore } from "../../../store/csv";
import ProcessorNode from "./ProcessorNode.vue";

const csvStore = useCsvStore();
</script>

<template>
  <div class="pl-4 pt-4">
    <h3>{{ csvStore.selectedSchema ? csvStore.selectedSchema.title : "" }}</h3>
  </div>

  <div v-if="csvStore.selectedSchema" class="tree pa-4">
    <ProcessorNode
      v-for="[k, v] in Object.entries(csvStore.selectedSchema.properties)"
      :key="`1-csv-${k}`"
      :node="[k, v]"
      :level="1"
    />
  </div>
</template>

<style scoped>
.tree:before {
  content: "{";
}
.tree:after {
  content: "}";
}
</style>
