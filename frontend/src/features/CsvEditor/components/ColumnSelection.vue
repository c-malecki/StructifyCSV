<script lang="ts" setup>
import { computed } from "vue";
import { useStore } from "../../../store/store";
const store = useStore();

const csvCols = computed(() => (store.csv ? store.csv.columns : []));
</script>

<template>
  <div class="row wrap-row">
    <button
      v-for="(col, idx) in csvCols"
      :key="`header-${idx}`"
      :class="col.isSelected ? `is-selected` : ``"
      @click="store.toggleSelectColumn(idx)"
    >
      <b>{{ col.header }}</b>
    </button>
  </div>
</template>

<style scoped>
.column-editor {
  border: 1px solid black;
  padding: 0.5rem;
  max-width: 300px;
}
</style>
