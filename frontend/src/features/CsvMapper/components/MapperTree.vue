<script lang="ts" setup>
import { useCsvStore } from "../../../store/csv";
import MapperNode from "./MapperNode.vue";

const csvStore = useCsvStore();
</script>

<template>
  <v-sheet border>
    <h3 class="pb-1 pt-1 pl-4">
      Schema:
      {{ csvStore.selectedSchema ? csvStore.selectedSchema.title : "" }}
      <span v-if="!csvStore.selectedSchema" class="empty"> none imported </span>
    </h3>

    <v-divider />

    <div v-if="csvStore.selectedSchema" class="tree pa-4">
      <MapperNode
        v-for="[k, v] in Object.entries(csvStore.selectedSchema.properties)"
        :key="`1-csv-${k}`"
        :node="[k, v]"
        :level="1"
      />
    </div>
  </v-sheet>
</template>

<style scoped>
h3 {
  background-color: #cfd8dc;
}
.empty {
  opacity: 0.5;
  font-weight: bold;
}
.tree:before {
  content: "{";
}
.tree:after {
  content: "}";
}
</style>
