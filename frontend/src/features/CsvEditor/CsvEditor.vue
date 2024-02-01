<script lang="ts" setup>
import { ImportCsvData } from "../../../wailsjs/go/main/App";
import { reactive, computed, inject } from "vue";
import { SchemaValuesKey, type CsvFileData } from "../../types/editor.types";

const schemaValues = inject(SchemaValuesKey);
if (!schemaValues) {
  throw new Error(`Could not resolve ${SchemaValuesKey.description}`);
}

const curCsv = reactive<CsvFileData>({
  fileName: "test.csv",
  fileLocation: "/home/meeps/Documents/test.csv",
  headers: [
    { name: "First Name", selected: false },
    { name: "Last Name", selected: false },
    { name: "Age", selected: false },
    { name: "Location", selected: false },
    { name: "Email", selected: false },
    { name: "LinkedIn URL", selected: false },
    { name: "Title", selected: false },
    { name: "Company", selected: false },
  ],
});

const importCsv = async () => {
  try {
    const csvData = await ImportCsvData();
    if (csvData.headers !== null) {
      curCsv.fileName = csvData.fileName;
      curCsv.fileLocation = csvData.location;
      curCsv.headers = csvData.headers.map((s) => {
        return {
          name: s,
          selected: false,
        };
      });
    } else {
      console.log("canceled import");
    }
  } catch (err) {
    console.log(err);
  }
};

const totalSelected = computed(
  () => curCsv.headers.filter((h) => h.selected).length
);

const tableHeaders = [
  {
    title: "Selected",
  },
  {
    title: "Column Header",
  },
  {
    title: "Schema Attribute",
  },
];
</script>

<template>
  <v-card rounded="0">
    <v-card-text>
      <v-sheet border class="pa-2 mb-2">
        <h3>{{ curCsv.fileName }}</h3>
        <p>Total Headers: {{ curCsv.headers.length }}</p>
        <p>Selected Headers: {{ totalSelected }}</p>
      </v-sheet>
      <v-sheet border class="pa-2">
        <v-data-table
          :headers="tableHeaders"
          :header-props="{ align: 'start' }"
          :items="curCsv.headers"
          :items-per-page="25"
        >
          <template #item="{ item }">
            <tr>
              <td>
                <v-checkbox-btn v-model="item.selected" />
              </td>
              <td>{{ item.name }}</td>
              <td></td>
            </tr>
          </template>
        </v-data-table>
      </v-sheet>
    </v-card-text>
  </v-card>
</template>

<style scoped></style>
