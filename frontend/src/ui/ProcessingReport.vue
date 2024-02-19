<script lang="ts" setup>
import { ref, type PropType } from "vue";
import { type VDataTableVirtual } from "vuetify/components";
import { entity } from "../../wailsjs/go/models";

type RowError = {
  rowNum: number;
  property: string;
  type: string;
  value: string;
  colNum: number | null;
  error: string;
};

type PReport = {
  successes: number;
  failures: number;
  errors: RowError[];
};

const report: PReport = {
  successes: 1,
  failures: 1,
  errors: [
    {
      rowNum: 2,
      colNum: 1,
      property: "productName",
      type: "string",
      value: "name",
      error: "length is less than 5",
    },
    {
      rowNum: 2,
      colNum: 3,
      property: "totalBids",
      type: "integer",
      value: "imAString",
      error: "could not be converted to type integer",
    },
  ],
};

const headers: VDataTableVirtual["headers"] = [
  { title: "Row", align: "start", key: "row" },
  { title: "Property Name", align: "start", key: "propertyKey" },
  { title: "Property Type", align: "start", key: "propertyType" },
  { title: "Error Type", align: "start", key: "errorType" },
  { title: "Column", align: "start", key: "column" },
  { title: "Value", align: "start", key: "value" },
  { title: "Error Message", align: "start", key: "errorMessage" },
];

const props = defineProps({
  report: {
    type: Object as PropType<entity.CsvProcessingReport>,
    required: true,
  },
});
</script>

<template>
  <v-sheet border rounded class="pa-4">
    <h3>Processing Report</h3>
    <v-data-table-virtual :headers="headers" :items="props.report.rowErrors" />
  </v-sheet>
</template>

<style scoped>
p {
  white-space: pre-line;
}
</style>
