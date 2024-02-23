<script lang="ts" setup>
import { computed, type ComputedRef } from "vue";
import { useAppStore } from "../../store/app";
import { useCsvStore } from "../../store/csv";
import { type VDataTableVirtual } from "vuetify/components";
import MapperTree from "./MapperTree/MapperTree.vue";
import ProcessingReport from "./ProcessingReport/ProcessingReport.vue";

const appStore = useAppStore();
const csvStore = useCsvStore();

const headers: ComputedRef<VDataTableVirtual["headers"]> = computed(() =>
  csvStore.csvFile
    ? csvStore.csvFile.headers.map((ch) => {
        return {
          title: ch.header,
          letter: ch.column,
          align: "start",
          key: ch.column,
        };
      })
    : []
);

const refRows = computed(() =>
  csvStore.csvFile ? csvStore.csvFile.referenceRows : []
);
</script>

<template>
  <div :class="{ hide: appStore.appTab === 'schema' }">
    <v-sheet border class="pa-4 ma-2">
      <h3 v-if="csvStore.csvFile">
        CSV:
        {{ csvStore.csvFile.fileName }}
        <span class="lighten">reference rows</span>
      </h3>
      <h3 v-else>
        CSV:
        <span class="lighten"> none imported</span>
      </h3>

      <v-divider />

      <v-sheet border>
        <v-data-table-virtual :headers="headers" :items="refRows">
          <template #headers="{ columns }">
            <tr>
              <template v-for="col in columns" :key="col.key">
                <td>
                  <div class="d-flex flex-column text-center">
                    <b>{{ col["letter" as keyof typeof col] }}</b>
                    <span>{{ col.title }}</span>
                  </div>
                </td>
              </template>
            </tr>
          </template>
          <template #item="{ item }">
            <tr>
              <td v-for="cell in item">
                <span v-if="cell">
                  {{ cell }}
                </span>
                <span class="lighten" v-else>
                  <i>empty cell</i>
                </span>
              </td>
            </tr>
          </template>
        </v-data-table-virtual>
      </v-sheet>
    </v-sheet>

    <ProcessingReport />
    <MapperTree />
  </div>
</template>

<style scoped>
.hide {
  display: none;
}
.lighten {
  opacity: 0.5;
  font-weight: bold;
}
.v-table:deep(td) {
  border: 1px solid rgba(0, 0, 0, 0.06);
}
.v-table:deep(thead) {
  background-color: #eeeeee;
  /* background-color: #F5F5F5 */
}
</style>
