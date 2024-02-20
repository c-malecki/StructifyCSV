<script lang="ts" setup>
import { useCsvStore } from "../../../store/csv";
import { type VDataTableVirtual } from "vuetify/components";

const csvStore = useCsvStore();

const headers: VDataTableVirtual["headers"] = [
  { title: "Row", align: "start", key: "row" },
  { title: "Property Name", align: "start", key: "propertyKey" },
  { title: "Property Type", align: "start", key: "propertyType" },
  { title: "Error Type", align: "start", key: "errorType" },
  { title: "Column", align: "start", key: "column" },
  { title: "Value", align: "start", key: "value" },
  { title: "Error Message", align: "start", key: "errorMessage" },
];
</script>

<template>
  <v-sheet border class="mb-4">
    <h3 class="pb-1 pt-1 pl-4">Processing Report</h3>
    <v-data-table-virtual
      v-if="csvStore.processingReport"
      :headers="headers"
      :items="csvStore.processingReport.rowErrors"
    />
  </v-sheet>
</template>

<style scoped>
h3 {
  background-color: #cfd8dc;
}
p {
  white-space: pre-line;
}
</style>
