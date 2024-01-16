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

// const propSelection = computed(() => []);

// const mirrorCsvEdtior = reactive({
//   schemaIdx: store.csvEditor.selectedColumns[index].schemaIdx,
//   fieldIdx: store.csvEditor.selectedColumns[index].fieldIdx,
// });

// const updateStoreColSchema = (v: number) => {
//   store.csvEditor.selectedColumns[index].schemaIdx = v;
//   store.csvEditor.selectedColumns[index].fieldIdx = null;
//   mirrorCsvEdtior.fieldIdx = null;
// };

// const updateStoreColField = (v: number) => {
//   store.csvEditor.selectedColumns[index].fieldIdx = v;
// };
</script>

<template>
  <tr>
    <td>{{ item.header }}</td>
    <!-- <td>
      <VAutocomplete
        :items="propSelection"
        item-title="first_name"
        item-value="value"
        density="compact"
        hide-details
        style="max-width: 300px"
      />
    </td> -->
  </tr>
</template>
