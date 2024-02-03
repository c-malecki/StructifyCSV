<script lang="ts" setup>
import { ref, inject } from "vue";
import { CsvFileKey } from "../../types/editor.types";
import HeaderSelector from "./components/HeaderSelector.vue";
import ModelTree from "./components/ModelTree.vue";

const csvFile = inject(CsvFileKey);
if (!csvFile) {
  throw new Error(`Could not resolve ${CsvFileKey.description}`);
}

const tab = ref(1);
</script>

<template>
  <v-card border rounded="0" flat>
    <template #title>
      <h3>{{ csvFile.fileName }}</h3>
    </template>

    <v-divider />

    <v-tabs v-model="tab">
      <v-tab :value="0">Headers</v-tab>
      <v-tab :value="1">Model</v-tab>
    </v-tabs>

    <v-divider />

    <HeaderSelector v-if="tab === 0" />
    <ModelTree v-else />
  </v-card>
</template>
