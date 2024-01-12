<script lang="ts" setup>
import { useStore } from "../../store/store";
import ColumnSelection from "./components/ColumnSelection.vue";
import EditColumn from "./components/EditColumn.vue";
const store = useStore();
</script>

<template>
  <div class="inner-container">
    <button @click="store.importCsv">Import CSV</button>
    <div v-if="store.csv !== null">
      <h3>{{ store.csv.fileName }}</h3>
      <p>CSV Headers</p>
      <ColumnSelection />
      <p>Selected Headers: {{ store.selectedColumns.length }}</p>
      <EditColumn v-for="(column, cIdx) in store.selectedColumns" :column-index="cIdx" />
    </div>
  </div>
</template>

<style scoped>
.column-editor {
  border: 1px solid black;
  padding: 0.5rem;
  max-width: 300px;
}

p {
  margin: 0;
}

button {
  margin-left: auto;
}
</style>
