<script lang="ts" setup>
import ModelEditor from "./features/ModelEditor/ModelEditor.vue";
import CsvEditor from "./features/CsvEditor/CsvEditor.vue";
import SchemaEditor from "./features/SchemaEditor/SchemaEditor.vue";
import { useStore } from "./store/store";
const store = useStore();
</script>

<template>
  <v-app id="app">
    <v-main style="max-width: 1200px">
      <ModelEditor />
      <SchemaEditor v-if="store.model" />

      <CsvEditor />
    </v-main>
  </v-app>
</template>

<style></style>
