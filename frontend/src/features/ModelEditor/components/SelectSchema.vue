<script lang="ts" setup>
import { useStore } from "../../../store/store";
const store = useStore();
</script>

<template>
  <div>
    <h3>Select Schema</h3>
    <div class="col">
      <label for="Schema">Schema</label>
      <select name="Schema" v-model="store.selectedSchema">
        <option v-for="(opt, idx) in store.model.schemas" :key="opt.name" :value="idx">
          {{ opt.name }}
        </option>
      </select>
    </div>
  </div>
</template>
