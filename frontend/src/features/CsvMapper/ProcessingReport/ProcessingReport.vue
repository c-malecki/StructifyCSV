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
  <v-dialog :model-value="csvStore.showReport" persistent max-height="100%">
    <v-sheet border class="mb-4">
      <v-toolbar density="compact" color="blue-grey-lighten-4">
        <v-toolbar-title>Processing Report</v-toolbar-title>
        <v-spacer />
        <v-btn icon="mdi-close" @click="csvStore.showReport = false" />
      </v-toolbar>
      <v-data-table-virtual
        v-if="csvStore.processingReport"
        :headers="headers"
        :items="csvStore.processingReport.rowErrors"
        height="600"
      />
    </v-sheet>
  </v-dialog>
</template>

<style scoped>
p {
  white-space: pre-line;
}
</style>
