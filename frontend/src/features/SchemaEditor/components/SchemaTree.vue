<script lang="ts" setup>
import { computed } from "vue";
import SchemaNode from "./SchemaNode.vue";
const props = defineProps({
  schema: {
    type: Object,
    required: true,
  },
});
const keys = computed(() => Object.keys(props.schema));
</script>

<template>
  <v-sheet border class="schema_tree pa-2">
    <SchemaNode v-for="(key, i) in keys" :label="key" :val="schema[key]" :key="key" :margin="16" />
  </v-sheet>
</template>

<style scoped>
.schema_tree:before {
  content: "{";
}
.schema_tree:after {
  content: "}";
}
</style>
