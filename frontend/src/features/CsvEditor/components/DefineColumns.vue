<script lang="ts" setup>
import { computed } from "vue";
import { useStore } from "../../../store/store";
import EditableColumn from "./EditableColumn.vue";
const store = useStore();
const headers = [
  {
    title: "Header Name",
    key: "name",
    sortable: false,
  },
  {
    title: "Assigned Schema",
    key: "schemaIdx",
    sortable: false,
  },
  {
    title: "Assigned Field",
    key: "fieldIdx",
    sortable: false,
  },
];
</script>

<template>
  <v-sheet border>
    <VDataTable :headers="headers" :items="store.csvEditor.selectedColumns" :items-per-page="25" density="compact">
      <template #item="{ item, index }">
        <EditableColumn :item="item" :index="index" />
      </template>
    </VDataTable>
  </v-sheet>
</template>
