<script lang="ts" setup>
import { useStore } from "../../../store/store";
import EditField from "./EditField.vue";
const store = useStore();
</script>

<template>
  <div>
    <p>
      Selected Schema: <b>{{ store.getSelectedSchemaLabel }}</b>
    </p>
    <pre>{{ JSON.stringify(store.formatSchemaDisplay, null, 2) }}</pre>

    <div>
      <h3>Fields</h3>
      <EditField
        v-for="(field, fIdx) in store.getSelectedSchemaFields"
        :key="`${store.getSelectedSchemaLabel}-field-${fIdx}`"
        :field-index="fIdx"
      />
    </div>

    <div class="col">
      <button @click="store.addField">add field</button>
    </div>
  </div>
</template>

<style scoped>
pre {
  padding: 0.2rem;
  border: 1px solid black;
}
</style>
