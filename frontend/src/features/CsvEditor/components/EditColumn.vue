<script lang="ts" setup>
import { computed } from "vue";
import { useStore } from "../../../store/store";

const props = defineProps({
  columnIndex: {
    type: Number,
    required: true,
  },
});

const store = useStore();
const fieldOpts = computed(() =>
  store.selectedColumns[props.columnIndex].schema !== null
    ? store.model.schemas[store.selectedColumns[props.columnIndex].schema!].fields
    : []
);
</script>

<template>
  <div class="column-editor">
    <p>
      <b>{{ store.selectedColumns[props.columnIndex].name }}</b>
    </p>

    <div class="row">
      <div class="col">
        <label for="Schema">Schema</label>
        <select name="Schema" v-model="store.selectedColumns[props.columnIndex].schema">
          <option v-for="(opt, sIdx) in store.model.schemas" :key="opt.name" :value="sIdx">
            {{ opt.name }}
          </option>
        </select>
      </div>

      <div class="col">
        <label for="Field">Field</label>
        <select
          name="Field"
          v-model="store.selectedColumns[props.columnIndex].field"
          :disabled="store.selectedColumns[props.columnIndex].schema === null"
        >
          <option v-for="(opt, fIdx) in fieldOpts" :key="opt.name" :value="fIdx">
            {{ opt.name }}
          </option>
        </select>
      </div>
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
