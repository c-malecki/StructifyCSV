<script lang="ts" setup>
import TitleBar from "./ui/TitleBar.vue";
import AppBar from "./ui/AppBar.vue";
import CsvMapper from "./features/CsvMapper/CsvMapper.vue";
import SchemaEditor from "./features/SchemaEditor/SchemaEditor.vue";
</script>

<template>
  <v-app id="app">
    <v-main>
      <TitleBar />
      <AppBar />
      <v-sheet class="pa-4">
        <SchemaEditor />
        <CsvMapper />
      </v-sheet>
    </v-main>
  </v-app>
</template>
