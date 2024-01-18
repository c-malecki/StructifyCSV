<script lang="ts" setup>
import { computed } from "vue";
import { useStore } from "../../store/store";
import EditHeaders from "./components/EditHeaders.vue";
import DefineColumns from "./components/DefineColumns.vue";
const store = useStore();
const curForm = computed(() => store.csvEditor.curForm);
const header = computed(() => {
  switch (curForm.value) {
    case "new":
      return "New Model";
    case "edit":
      return "Select Headers";
    case "none":
      return `CSV: ${store.csv ? store.csv.fileName : "None Selected"}`;
  }
});
const headerOpts = computed(() => (store.csv ? store.csv.headers : []));
</script>

<template>
  <v-card rounded="0">
    <v-toolbar density="compact" color="grey-lighten-2">
      <v-menu>
        <template v-slot:activator="{ props }">
          <v-app-bar-nav-icon v-bind="props"></v-app-bar-nav-icon>
        </template>

        <v-list variant="flat" border density="compact" elevation="0">
          <v-list-item density="compact" title="Import CSV" @click="store.importCsv()" />
        </v-list>
      </v-menu>

      <v-toolbar-title>{{ header }}</v-toolbar-title>
    </v-toolbar>

    <v-card-text>
      <div v-if="store.csv">
        <!-- <v-btn @click="store.runProcessCsvDescriptor()">test</v-btn> -->
        <EditHeaders v-if="curForm === 'edit'" />
        <v-sheet border class="pa-2">
          <div class="d-flex">
            <v-btn size="x-small" @click="store.changeCsvEditorForm('edit')" class="mr-4">edit</v-btn>
            <h3 class="mr-4">Total Headers: {{ headerOpts.length }}</h3>
            <h3>Used Headers: {{ store.csvEditor.selectedColumns.length }}</h3>
          </div>
        </v-sheet>
        <DefineColumns v-if="curForm === 'none'" />
      </div>
    </v-card-text>
  </v-card>
</template>

<style scoped></style>
