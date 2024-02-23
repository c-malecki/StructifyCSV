<script lang="ts" setup>
import { computed } from "vue";
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
</script>

<template>
  <v-toolbar height="36px" extended extension-height="26">
    <v-btn text="Schema Editor" @click="appStore.setAppTab('schema')" />
    <v-btn text="CSV Mapper" @click="appStore.setAppTab('csv')" />
    <template #extension>
      <div class="d-flex flex-column w-100">
        <v-divider />
        <div v-if="appTab === 'schema'" class="d-flex">
          <v-btn
            prepend-icon="mdi-file-plus-outline"
            size="small"
            @click="schemaStore.newJsonSchema()"
            class="ml-4"
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
        </div>
        <div v-if="appTab === 'csv'" class="d-flex">
          <v-btn
            prepend-icon="mdi-table-arrow-left"
            size="small"
            @click="csvStore.importCsvFileData()"
            class="ml-4"
          >
            import csv
          </v-btn>
          <v-btn
            prepend-icon="mdi-folder-open-outline"
            size="small"
            @click="csvStore.selectJsonSchema()"
          >
            import schema
          </v-btn>
          <v-btn
            prepend-icon="mdi-vector-polyline-edit"
            size="small"
            @click="csvStore.processCsvWithSelectedSchema()"
            :disabled="!csvStore.canRunCsvProcessor"
          >
            structify csv
          </v-btn>
          <v-btn
            prepend-icon="mdi-chart-timeline"
            size="small"
            @click="csvStore.showReport = true"
            :disabled="!csvStore.processingReport"
          >
            error report
          </v-btn>
        </div>
      </div>

      <!-- <v-toolbar v-if="appTab === 'schema'" density="compact" border>
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
      </v-toolbar> -->
      <!-- <v-toolbar v-if="appTab === 'csv'" density="compact" border>
        <v-btn
          prepend-icon="mdi-table-arrow-left"
          size="small"
          @click="csvStore.importCsvFileData()"
        >
          import csv
        </v-btn>
        <v-btn
          prepend-icon="mdi-folder-open-outline"
          size="small"
          class="mr-4"
          @click="csvStore.selectJsonSchema()"
        >
          import schema
        </v-btn>
        <v-btn
          prepend-icon="mdi-vector-polyline-edit"
          size="small"
          @click="csvStore.processCsvWithSelectedSchema()"
          :disabled="!csvStore.canRunCsvProcessor"
        >
          structify
        </v-btn>
      </v-toolbar> -->
    </template>
  </v-toolbar>
</template>
