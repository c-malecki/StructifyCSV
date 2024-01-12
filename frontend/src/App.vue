<script lang="ts" setup>
import ModelEditor from "./features/ModelEditor/ModelEditor.vue";
import CsvEditor from "./features/CsvEditor/CsvEditor.vue";
import SchemaEditor from "./features/SchemaEditor/SchemaEditor.vue";
import { useStore } from "./store/store";
const store = useStore();
</script>

<template>
  <v-app id="app">
    <v-main>
      <ModelEditor />
      <SchemaEditor v-if="store.model" />
      <CsvEditor />
    </v-main>
  </v-app>
</template>

<style>
.error {
  font-size: 14px;
  margin-block-start: 0.2em;
  margin-block-end: 0.2em;
  color: red;
}

.is-selected {
  background-color: aquamarine;
}
</style>
