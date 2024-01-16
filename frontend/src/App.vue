<script lang="ts" setup>
import CsvEditor from "./features/CsvEditor/CsvEditor.vue";
import SchemaEditor from "./features/SchemaEditor/SchemaEditor.vue";
</script>

<template>
  <v-app id="app">
    <v-main style="max-width: 1200px">
      <SchemaEditor />
      <CsvEditor />
    </v-main>
  </v-app>
</template>

<style></style>
