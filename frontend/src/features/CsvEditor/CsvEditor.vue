<script lang="ts" setup>
import { ref, inject } from "vue";
import { CsvFileKey } from "../../types/editor.types";
// import HeaderDescriptions from "./components/HeaderDescriptions.vue";
import ModelTree from "./components/ModelTree.vue";

const csvFile = inject(CsvFileKey);
if (!csvFile) {
  throw new Error(`Could not resolve ${CsvFileKey.description}`);
}

// const tab = ref(1);
</script>

<template>
  <v-card border rounded="0" flat>
    <template #title>
      <h3>{{ csvFile.fileName }}</h3>
    </template>

    <v-divider />

    <!-- <v-tabs v-model="tab">
      <v-tab :value="0">csv headers</v-tab>
      <v-tab :value="1">csv model</v-tab>
    </v-tabs>

    <v-divider /> -->

    <ModelTree />

    <!-- <v-window v-model="tab">
      <v-window-item :value="0">
        <HeaderDescriptions />
      </v-window-item>
      <v-window-item :value="1">
        <ModelTree />
      </v-window-item>
    </v-window> -->
  </v-card>
</template>
