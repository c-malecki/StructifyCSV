<script lang="ts" setup>
import { useStore } from "./store/store";
import ModelEditor from "./features/ModelEditor/ModelEditor.vue";
import CsvEditor from "./features/CsvEditor/CsvEditor.vue";
const store = useStore();
</script>

<template>
  <div class="main">
    <div class="container">
      <button @click="store.setColumns">select csv file</button>
    </div>

    <CsvEditor />
    <ModelEditor />
  </div>
</template>

<style>
.main {
  display: flex;
  flex-direction: column;
  padding: 1rem;
}

.tabs {
  background-color: #fff;
  margin-bottom: 1rem;
}

.row {
  display: flex;
  flex-direction: row;
}

.col {
  display: flex;
  flex-direction: column;
}

.container {
  background-color: #fff;
  height: 100%;
  width: 100%;
  padding: 0.5rem;
}

.error {
  font-size: 14px;
  margin-block-start: 0.2em;
  margin-block-end: 0.2em;
  color: red;
}

h1,
h2,
h3,
h4,
h5 {
  margin-block-start: 0.5rem;
  margin-block-end: 0.5rem;
}
</style>
