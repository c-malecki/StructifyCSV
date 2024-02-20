<script lang="ts" setup>
import { ref, computed } from "vue";
import { useAppStore, type AppTab } from "../store/app";
import { useSchemaStore } from "../store/schema";
import { useCsvStore } from "../store/csv";

const appStore = useAppStore();
const schemaStore = useSchemaStore();
const csvStore = useCsvStore();

const appTab = computed({
  get: () => appStore.appTab,
  set: (tab: AppTab) => appStore.setAppTab(tab),
});

const isHovering = ref(false);
</script>

<template>
  <v-hover v-model="isHovering">
    <template v-slot:default="{ props: hoverProps }">
      <v-toolbar v-bind="hoverProps" density="compact" :extended="isHovering">
        <v-tabs
          v-model="appTab"
          density="compact"
          bg-color="blue-grey-lighten-5"
        >
          <v-tab value="schema"> Schema Editor </v-tab>
          <v-tab value="csv"> CSV Processor </v-tab>
        </v-tabs>
        <template v-if="isHovering" #extension>
          <v-toolbar v-if="appTab === 'schema'" density="compact" border>
            <v-btn
              prepend-icon="mdi-file-plus-outline"
              size="small"
              @click="schemaStore.newJsonSchema()"
            >
              new
            </v-btn>
            <v-btn
              prepend-icon="mdi-folder-open-outline"
              size="small"
              @click="schemaStore.importJsonSchema()"
            >
              open
            </v-btn>
            <v-btn
              prepend-icon="mdi-content-save-edit-outline"
              size="small"
              @click="schemaStore.exportJsonSchema()"
            >
              save as
            </v-btn>
          </v-toolbar>
          <v-toolbar v-if="appTab === 'csv'" density="compact" border>
            <v-btn
              prepend-icon="mdi-table-arrow-left"
              size="small"
              @click="csvStore.importCsvFileData()"
            >
              import
            </v-btn>
            <v-btn
              prepend-icon="mdi-vector-polyline-edit"
              size="small"
              @click="csvStore.processCsvWithSelectedSchema()"
              :disabled="!csvStore.canRunCsvProcessor"
            >
              process
            </v-btn>
          </v-toolbar>
        </template>
      </v-toolbar>
    </template>
  </v-hover>
</template>
