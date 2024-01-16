<script lang="ts" setup>
import { computed } from "vue";
const props = defineProps({
  label: {
    type: String,
    required: true,
  },
  val: {
    type: [String, Object],
  },
  margin: {
    type: Number,
    default: 8,
  },
});
const marginStyle = computed(() => `margin-left: ${props.margin}px`);
const valType = computed(() => (typeof props.val === "string" ? "string" : "object"));
const nextKeys = computed(() => (valType.value === "string" ? [] : Object.keys(props.val!)));
</script>

<template>
  <div :style="marginStyle" :class="{ closing_bracket: valType === 'object' }">
    <div v-if="valType === 'string'">
      <b>{{ label }}:</b> {{ val }}
    </div>
    <div v-if="valType === 'object'">
      <div class="opening_bracket">
        <b>{{ label }}:</b>
      </div>
      <SchemaNode
        v-for="(key, i) in nextKeys"
        :key="key"
        :label="key"
        :val="valType === 'object' ?  (props.val as Record<string,any>)[key] : props.val"
        :margin="props.margin + 4"
      />
    </div>
  </div>
</template>

<style scoped>
.opening_bracket:after {
  content: "{";
  margin-left: 8px;
}
.closing_bracket:after {
  content: "}";
}
</style>
