<script lang="ts" setup>
import { useStore } from "../../../store/store";
const store = useStore();
</script>

<template>
  <div class="container">
    <h3>{{ store.getCurModelLabel }}</h3>
    <div class="col">
      <label for="Schema">Model Base Schema</label>
      <select name="Schema" v-model="store.model.baseSchema">
        <option v-for="(opt, idx) in store.model.schemas" :key="opt.name" :value="idx">
          {{ opt.name }}
        </option>
      </select>
    </div>
    <pre>{{ JSON.stringify(store.formatModelDisplay, null, 2) }}</pre>
  </div>
</template>

<style scoped>
pre {
  padding: 0.2rem;
  border: 1px solid black;
}
</style>
