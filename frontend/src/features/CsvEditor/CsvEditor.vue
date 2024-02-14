<script lang="ts" setup>
import { inject } from "vue";
import { CsvFileKey } from "./CsvEditor.types";
import ModelTree from "./components/ModelTree.vue";

const csvFile = inject(CsvFileKey);
if (!csvFile) {
  throw new Error(`Could not resolve ${CsvFileKey.description}`);
}
</script>

<template>
  <v-card border rounded="0" flat>
    <template #title>
      <h3>{{ csvFile.fileName }}</h3>
    </template>

    <v-divider />

    <ModelTree />
  </v-card>
</template>
