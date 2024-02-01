<script lang="ts" setup>
import { SchemaValuesKey } from "./schemaEditor.util";
import { ref, inject } from "vue";
import SchemaTree from "./components/SchemaTree.vue";
import SchemaInfoForm from "./components/SchemaInfoForm.vue";
import type { SchemaValues } from "./schemaEditor.types";

const emit = defineEmits(["updateSchema"]);

const schemaValues = inject(SchemaValuesKey);
if (!schemaValues) {
  throw new Error(`Could not resolve ${SchemaValuesKey.description}`);
}
const showEditForm = ref(false);

const handleUpdateSchema = (vals: Omit<SchemaValues, "properties">) => {
  emit("updateSchema", vals);
  showEditForm.value = false;
};
</script>

<template>
  <v-card rounded="0">
    <v-card-text>
      <v-sheet border class="pa-2 mb-2">
        <div v-if="!showEditForm" class="relative">
          <h3>{{ schemaValues.title }}</h3>
          <p>{{ schemaValues.description }}</p>
          <v-btn
            position="absolute"
            size="small"
            location="top right"
            @click="showEditForm = true"
          >
            edit
          </v-btn>
        </div>

        <SchemaInfoForm
          v-else
          @close-form="showEditForm = false"
          @update-schema="handleUpdateSchema"
        />
      </v-sheet>
      <SchemaTree />
    </v-card-text>
  </v-card>
</template>

<style scoped>
p {
  white-space: pre;
}
.relative {
  position: relative;
}
</style>
