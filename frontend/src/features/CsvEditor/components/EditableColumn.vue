<script lang="ts" setup>
import { reactive, computed } from "vue";
import { useStore } from "../../../store/store";
const store = useStore();
const props = defineProps({
  item: {
    type: Object,
    required: true,
  },
  index: {
    type: Number,
    required: true,
  },
});
const { item, index } = props;

const mirrorCsvEdtior = reactive({
  schemaIdx: store.csvEditor.selectedColumns[index].schemaIdx,
  fieldIdx: store.csvEditor.selectedColumns[index].fieldIdx,
});

const updateStoreColSchema = (v: number) => {
  store.csvEditor.selectedColumns[index].schemaIdx = v;
  store.csvEditor.selectedColumns[index].fieldIdx = null;
  mirrorCsvEdtior.fieldIdx = null;
};

const updateStoreColField = (v: number) => {
  store.csvEditor.selectedColumns[index].fieldIdx = v;
};

const schemaOpts = computed(() =>
  store.model
    ? store.model.schemas.map((s, i) => {
        return { name: s.name, value: i };
      })
    : []
);
const fieldOpts = computed(() => {
  const schemaIdx = store.csvEditor.selectedColumns[index].schemaIdx;
  return store.model && schemaIdx !== null
    ? store.model.schemas[schemaIdx].fields.map((f, i) => {
        return { name: f.name, value: i };
      })
    : [];
});
</script>

<template>
  <tr>
    <td>{{ item.header }}</td>
    <td>
      <VAutocomplete
        v-model="mirrorCsvEdtior.schemaIdx"
        :items="schemaOpts"
        item-title="name"
        item-value="value"
        density="compact"
        hide-details
        :disabled="!store.model || !store.model.schemas.length"
        style="max-width: 300px"
        @update:model-value="updateStoreColSchema"
      />
    </td>
    <td>
      <VAutocomplete
        v-model="mirrorCsvEdtior.fieldIdx"
        :items="fieldOpts"
        item-title="name"
        item-value="value"
        density="compact"
        hide-details
        :disabled="store.csvEditor.selectedColumns[index].schemaIdx === null"
        style="max-width: 300px"
        @update:model-value="updateStoreColField"
      />
    </td>
  </tr>
</template>
