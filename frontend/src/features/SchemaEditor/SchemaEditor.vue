<script lang="ts" setup>
import { SchemaValuesKey, type SchemaValues } from "../../types/editor.types";
import { ref, inject } from "vue";
import SchemaTree from "./components/SchemaTree.vue";
import SchemaInfoForm from "./components/SchemaInfoForm.vue";

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
  <v-card border rounded="0" flat>
    <template v-if="!showEditForm" #title>
      <div class="relative">
        <h3>{{ schemaValues.title }}</h3>
        <v-btn
          position="absolute"
          size="small"
          location="top right"
          @click="showEditForm = true"
        >
          edit
        </v-btn>
      </div>
    </template>

    <template v-if="!showEditForm" #subtitle>
      <p>{{ schemaValues.description }}</p>
    </template>

    <template v-if="showEditForm" #text>
      <SchemaInfoForm
        @close-form="showEditForm = false"
        @update-schema="handleUpdateSchema"
      />
    </template>

    <v-divider />

    <SchemaTree />
  </v-card>
</template>

<style scoped>
.relative {
  position: relative;
}

.v-card:deep(.v-card-subtitle) {
  opacity: 1;
  white-space: pre;
}
</style>
