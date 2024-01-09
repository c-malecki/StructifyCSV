<script lang="ts" setup>
import { useStore } from "../../../store/store";
import { type DataTypes } from "../../../store/store";

const store = useStore();

const fieldOptions: DataTypes[] = [
  { name: "text (string)", value: "string" },
  { name: "int (whole number)", value: "int" },
  { name: "boolean (true/false)", value: "bool" },
];

const props = defineProps({
  fieldIndex: {
    type: Number,
    required: true,
  },
});
</script>

<template>
  <div class="row field-editor">
    <div class="col">
      <label for="Field Name">Field Name</label>
      <input v-model="store.model.schemas[store.selectedSchema].fields[props.fieldIndex].name" name="Field Name" />
    </div>

    <div class="col">
      <label for="Field Type">Field Type</label>
      <select name="Field Type" v-model="store.model.schemas[store.selectedSchema].fields[props.fieldIndex].dataType">
        <option v-for="opt in fieldOptions" :key="opt.name" :value="opt">
          {{ opt.name }}
        </option>
      </select>
    </div>
  </div>
</template>

<style scoped>
.field-editor {
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
