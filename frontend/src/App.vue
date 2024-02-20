<script lang="ts" setup>
import TitleBar from "./ui/TitleBar.vue";
import AppBar from "./ui/AppBar.vue";
import CsvProcessor from "./features/CsvProcessor/CsvProcessor.vue";
import SchemaEditor from "./features/SchemaEditor/SchemaEditor.vue";
</script>

<template>
  <v-app id="app">
    <v-main>
      <TitleBar />
      <AppBar />
      <SchemaEditor />
      <CsvProcessor />
    </v-main>
  </v-app>
</template>
