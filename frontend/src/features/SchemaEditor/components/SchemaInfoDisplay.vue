<script lang="ts" setup>
import { useSchemaStore } from "../../../store/schema";

const schemaStore = useSchemaStore();

const emit = defineEmits(["showEdit"]);
</script>

<template>
  <div class="relative pa-4">
    <h3>{{ schemaStore.jsonSchema.title }}</h3>
    <p>{{ schemaStore.jsonSchema.description }}</p>
    <v-btn
      position="absolute"
      size="small"
      location="top right"
      @click="emit('showEdit')"
    >
      edit
    </v-btn>
  </div>
</template>

<style scoped>
.hide {
  display: none;
}
.relative {
  position: relative;
}
</style>
