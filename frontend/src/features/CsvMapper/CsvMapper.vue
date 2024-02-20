<script lang="ts" setup>
import { useAppStore } from "../../store/app";
import CsvReference from "./components/CsvReference.vue";
import MapperTree from "./components/MapperTree.vue";
import ProcessingReport from "./components/ProcessingReport.vue";

const appStore = useAppStore();
</script>

<template>
  <v-sheet rounded="0" flat :class="{ hide: appStore.appTab === 'schema' }">
    <CsvReference />
    <!-- <ProcessingReport /> -->
    <MapperTree />
  </v-sheet>
</template>

<style scoped>
.hide {
  display: none;
}
</style>
