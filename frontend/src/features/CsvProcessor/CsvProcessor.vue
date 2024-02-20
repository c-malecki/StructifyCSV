<script lang="ts" setup>
import { useAppStore } from "../../store/app";
import { useCsvStore } from "../../store/csv";
import ProcessorTree from "./components/ProcessorTree.vue";
import ProcessingReport from "./components/ProcessingReport.vue";

const appStore = useAppStore();
const csvStore = useCsvStore();
</script>

<template>
  <v-sheet
    border
    rounded="0"
    flat
    :class="{ hide: appStore.appTab === 'schema' }"
  >
    <h3 v-if="csvStore.csvFile">{{ csvStore.csvFile.fileName }}</h3>

    <v-divider />
    <ProcessingReport />
    <ProcessorTree />
  </v-sheet>
</template>

<style scoped>
.hide {
  display: none;
}
</style>
