<script lang="ts" setup>
import { computed, type ComputedRef } from "vue";
import { useCsvStore } from "../../../store/csv";
import { type VDataTableVirtual } from "vuetify/components";

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
</script>

<template>
  <v-sheet border class="mb-4">
    <h3 v-if="csvStore.csvFile" class="pb-1 pt-1 pl-4">
      File:
      {{ csvStore.csvFile.fileName }} <span class="empty">reference rows</span>
    </h3>
    <h3 v-else class="pb-1 pt-1 pl-4">
      File:
      <span class="empty"> none imported</span>
    </h3>

    <v-divider />

    <v-data-table-virtual
      v-if="csvStore.csvFile"
      :headers="headers"
      :items="csvStore.csvFile.referenceRows"
    >
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
            <span class="empty" v-else>
              <i>empty cell</i>
            </span>
          </td>
        </tr>
      </template>
    </v-data-table-virtual>
  </v-sheet>
</template>

<style scoped>
h3 {
  background-color: #cfd8dc;
}
.v-table:deep(td) {
  border: 1px solid rgba(0, 0, 0, 0.06);
}
.v-table:deep(thead) {
  background-color: #eeeeee;
  /* background-color: #F5F5F5 */
}
.empty {
  opacity: 0.5;
  font-weight: bold;
}
</style>
