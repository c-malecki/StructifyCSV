<script lang="ts" setup>
// import CsvEditor from "./features/CsvEditor/CsvEditor.vue";
import SchemaEditor from "./features/SchemaEditor/SchemaEditor.vue";

// todo: How to handle arrays? Do I need to have support for
// multi-typed properties?
// look at next stesp of JSON Schema spec integration
</script>

<template>
  <v-app id="app">
    <v-main style="max-width: 1200px">
      <SchemaEditor />
      <!-- <CsvEditor /> -->
    </v-main>
  </v-app>
</template>

<style></style>
