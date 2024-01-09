<script lang="ts" setup>
import { useStore } from "../../../store/store";
const store = useStore();
</script>

<template>
  <div class="row wrap-row">
    <button
      v-for="(col, idx) in store.csv.columns"
      :key="`header-${idx}`"
      :class="col.isSelected ? `is-selected` : ``"
      @click="store.toggleSelectColumn(idx)"
    >
      <b>{{ col.header }}</b>
    </button>
  </div>
</template>

<style scoped>
.wrap-row {
  flex-wrap: wrap;
}

.column-editor {
  border: 1px solid black;
  padding: 0.5rem;
  max-width: 300px;
}

.is-selected {
  background-color: aquamarine;
}
</style>
