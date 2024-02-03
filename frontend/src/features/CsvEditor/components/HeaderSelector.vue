<script lang="ts" setup>
import { ref, computed, inject } from "vue";
import { CsvModelKey, SchemaValuesKey } from "../../../types/editor.types";
import { getAllSchemaMapProperties } from "../../../util/transform";
import type { VDataTable } from "vuetify/components";

const schemaValues = inject(SchemaValuesKey);
if (!schemaValues) {
  throw new Error(`Could not resolve ${SchemaValuesKey.description}`);
}
const csvData = inject(CsvModelKey);
if (!csvData) {
  throw new Error(`Could not resolve ${CsvModelKey.description}`);
}

const allSchemaMapProperties = computed(() =>
  getAllSchemaMapProperties(schemaValues.properties)
);

const headerSearch = ref("");
const selectedHeaders = computed(() =>
  csvData.headers.filter((h) => h.isSelected)
);

const tableHeaders: VDataTable["$props"]["headers"] = [
  {
    title: "Selected",
    align: "start",
    key: "isSelected",
    sortable: false,
  },
  {
    title: "CSV Header",
    align: "start",
    key: "name",
    sortable: false,
  },
  {
    title: "Schema Property",
    align: "start",
    key: "schemaProperty",
    sortable: false,
  },
];
</script>

<template>
  <div class="pa-4">
    <p class="mb-1">Total: {{ csvData.headers.length }}</p>
    <p class="mb-2">Selected: {{ selectedHeaders.length }}</p>

    <v-text-field
      v-model="headerSearch"
      label="Search Headers"
      prepend-inner-icon="mdi-magnify"
      single-line
      hide-details
    />
  </div>

  <v-divider />

  <v-data-table
    :headers="tableHeaders"
    :items="csvData.headers"
    :items-per-page="10"
    :search="headerSearch"
    items-per-page-text="Headers per page:"
  >
    <template #item="{ item }">
      <tr>
        <td>
          <v-checkbox-btn v-model="item.isSelected" />
        </td>
        <td>
          {{ item.header }}
        </td>
        <td>
          <!-- <VAutocomplete
            v-model="item.schemaProperty"
            :items="allSchemaMapProperties"
            label="Property"
            item-title="key"
            item-value="value"
            return-object
            hide-details
            clearable
            persistent-clear
          /> -->
          <i v-if="!item.schemaProperty">None Assigned</i>
          <span v-else
            >{{ item.schemaProperty.schemaPath }} ({{
              item.schemaProperty.value
            }})</span
          >
        </td>
      </tr>
    </template>
  </v-data-table>
</template>

<style scoped>
.v-table:deep(.v-data-table-header__content) {
  font-weight: bold;
}
</style>
