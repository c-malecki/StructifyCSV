<script lang="ts" setup>
import { computed, ref } from "vue";
import TitleBar from "./ui/TitleBar.vue";
import CsvEditor from "./features/CsvEditor/CsvEditor.vue";
import SchemaEditor from "./features/SchemaEditor/SchemaEditor.vue";

const schemaEditor = ref<typeof SchemaEditor | null>(null);
const csvEditor = ref<typeof CsvEditor | null>(null);

const tab = ref(0);
const curMenu = ref<"schema" | "csv">("schema");

const debounce = ref<number | undefined>(undefined);
const isHovering = ref(false);

const checkHovering = (bool: boolean, menu: "schema" | "csv") => {
  clearTimeout(debounce.value);
  if (!bool) {
    debounce.value = setTimeout(() => (isHovering.value = false), 200);
  } else {
    curMenu.value = menu;
    isHovering.value = bool;
  }
};
</script>

<template>
  <v-app id="app">
    <v-main>
      <TitleBar />

      <v-toolbar density="compact" :extended="isHovering">
        <v-tabs v-model="tab" density="compact" bg-color="blue-grey-lighten-5">
          <v-hover
            @update:model-value="(b: boolean) => checkHovering(b, 'schema')"
          >
            <template v-slot:default="{ props: hoverProps }">
              <v-tab v-bind="hoverProps" :value="0"> Schema Editor </v-tab>
            </template>
          </v-hover>
          <v-hover
            @update:model-value="(b: boolean) => checkHovering(b, 'csv')"
          >
            <template v-slot:default="{ props: hoverProps }">
              <v-tab v-bind="hoverProps" :value="1"> CSV Parser </v-tab>
            </template>
          </v-hover>
        </v-tabs>
        <template v-if="isHovering" #extension>
          <v-hover
            @update:model-value="(b: boolean) => checkHovering(b, curMenu)"
          >
            <template v-slot:default="{ props: hoverProps }">
              <v-toolbar
                v-if="curMenu === 'schema'"
                v-bind="hoverProps"
                density="compact"
                border
              >
                <v-btn
                  prepend-icon="mdi-file-plus-outline"
                  size="small"
                  @click="schemaEditor!.newSchema()"
                >
                  new
                </v-btn>
                <v-btn
                  prepend-icon="mdi-folder-open-outline"
                  size="small"
                  @click="schemaEditor!.importSchema()"
                >
                  open
                </v-btn>
                <v-btn
                  prepend-icon="mdi-content-save-edit-outline"
                  size="small"
                  @click="schemaEditor!.exportSchema()"
                >
                  save as
                </v-btn>
              </v-toolbar>
              <v-toolbar
                v-if="curMenu === 'csv'"
                v-bind="hoverProps"
                density="compact"
                border
              >
                <v-btn
                  prepend-icon="mdi-table-arrow-left"
                  size="small"
                  @click="csvEditor!.importCsv()"
                >
                  import
                </v-btn>
                <v-btn
                  prepend-icon="mdi-vector-polyline-edit"
                  size="small"
                  @click="csvEditor!.processCsv()"
                  :disabled="!csvEditor!.csvFile || !csvEditor!.selectedSchema"
                >
                  process
                </v-btn>
              </v-toolbar>
            </template>
          </v-hover>
        </template>
      </v-toolbar>

      <SchemaEditor :hidden="tab === 1" ref="schemaEditor" />

      <CsvEditor :hidden="tab === 0" ref="csvEditor" />
    </v-main>
  </v-app>
</template>
