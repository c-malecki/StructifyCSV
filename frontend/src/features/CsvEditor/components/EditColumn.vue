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

const modelSchemaOpts = computed(() => (store.model !== null ? store.model.schemas : []));

const fieldOpts = computed(() => {
  const curColSchemaIdx = store.selectedColumns[props.columnIndex].schema;
  if (modelSchemaOpts.value.length > 0 && curColSchemaIdx !== null) {
    return modelSchemaOpts.value[curColSchemaIdx].fields;
  }
  return [];
});
</script>

<template>
  <div class="column-editor">
    <p>
      <b>{{ store.selectedColumns[props.columnIndex].name }}</b>
    </p>

    <div class="row">
      <div class="col">
        <label for="Schema">Schema</label>
        <select
          v-model="store.selectedColumns[props.columnIndex].schema"
          name="Schema"
          :disabled="!modelSchemaOpts.length"
        >
          <option v-for="(opt, sIdx) in modelSchemaOpts" :key="opt.name" :value="sIdx">
            {{ opt.name }}
          </option>
        </select>
      </div>

      <div class="col">
        <label for="Field">Field</label>
        <select
          v-model="store.selectedColumns[props.columnIndex].field"
          name="Field"
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
