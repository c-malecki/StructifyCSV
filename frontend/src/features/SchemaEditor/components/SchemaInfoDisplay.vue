<script lang="ts" setup>
import { useSchemaStore } from "../../../store/schema";

const schemaStore = useSchemaStore();

const emit = defineEmits(["showEdit"]);
</script>

<template>
  <v-sheet border class="mb-4">
    <h3>
      {{ schemaStore.jsonSchema.title }}
      <v-btn size="x-small" class="ml-2" @click="emit('showEdit')">
        edit
      </v-btn>
    </h3>

    <p class="pa-4">{{ schemaStore.jsonSchema.description }}</p>
  </v-sheet>
</template>

<style scoped>
h3 {
  display: flex;
  flex-direction: row;
  align-items: center;
  padding: 4px 4px 4px 16px;
  background-color: #cfd8dc;
}
.hide {
  display: none;
}

p {
  white-space: pre-line;
}
</style>
