<script lang="ts" setup>
import { reactive, ref } from "vue";
import { computed } from "vue";
const nodeProps = defineProps({
  label: {
    type: String,
    required: true,
  },
  val: {
    type: [String, Object],
  },
  level: {
    type: Number,
    default: 1,
  },
});

type DataTypes = { name: string; value: string };

const dataTypes: DataTypes[] = [
  { name: "string (text)", value: "string" },
  { name: "float (decimal)", value: "number" },
  { name: "integer", value: "integer" },
  { name: "object", value: "object" },
  { name: "array", value: "array" },
  { name: "boolean", value: "boolean" },
  { name: "null", value: "null" },
];

const colorScale = {
  1: { hover: "blue-lighten-5", font: "black" },
  2: { hover: "blue-lighten-4", font: "black" },
  3: { hover: "blue-lighten-3", font: "black" },
  4: { hover: "blue-lighten-2", font: "black" },
  5: { hover: "blue-lighten-1", font: "white" },
  6: { hover: "blue-darken-1", font: "white" },
  7: { hover: "blue-darken-2", font: "white" },
  8: { hover: "blue-darken-3", font: "white" },
  9: { hover: "blue-darken-4", font: "white" },
};
const marginStyle = computed(() => `margin-left: ${nodeProps.level * 16}px`);
const color = computed(() => colorScale[nodeProps.level as keyof typeof colorScale]);

const isEdit = ref(false);
const element = reactive({
  name: nodeProps.label,
  value: typeof nodeProps.val,
});
</script>

<template>
  <v-hover>
    <template v-slot:default="{ isHovering, props }">
      <v-card v-bind="props" :color="isHovering ? color.hover : undefined" variant="flat" :style="marginStyle">
        <div
          :class="{ closing_bracket: typeof nodeProps.val === 'object' }"
          :style="`color: ${isHovering ? color.font : 'black'}`"
          class="pt-1 pb-1"
        >
          <div v-if="typeof nodeProps.val === 'string'" class="d-flex">
            <p v-if="!isEdit">
              <b>{{ label }}:</b> {{ val }}
            </p>
            <div v-if="isHovering && !isEdit" class="d-flex">
              <v-btn size="x-small" class="ml-4" @click="isEdit = true">edit</v-btn>
              <v-btn size="x-small" class="ml-4">delete</v-btn>
            </div>

            <VForm v-if="isEdit">
              <div class="d-flex">
                <VTextField v-model="element.name" label="Property Name" style="width: 200px" />
                <VSelect
                  v-model="element.value"
                  :items="dataTypes"
                  item-title="name"
                  item-value="value"
                  style="width: 200px"
                />
                <v-btn size="x-small" class="ml-4">save</v-btn>
                <v-btn size="x-small" class="ml-4" @click="isEdit = false">cancel</v-btn>
              </div>
            </VForm>
          </div>
          <div v-else>
            <div class="d-flex opening_bracket pb-1">
              <p>
                <b>{{ label }}:</b>
              </p>
              <div v-if="isHovering && !isEdit" class="d-flex">
                <v-btn size="x-small" class="ml-4">add property</v-btn>
                <v-btn size="x-small" class="ml-4" @click="isEdit = true">edit</v-btn>
                <v-btn size="x-small" class="ml-4">delete</v-btn>
              </div>
              <VForm v-if="isEdit">
                <div class="d-flex">
                  <VTextField v-model="element.name" label="Property Name" style="width: 200px" />
                  <VSelect
                    v-model="element.value"
                    :items="dataTypes"
                    item-title="name"
                    item-value="value"
                    style="width: 200px"
                  />
                  <v-btn size="x-small" class="ml-4">save</v-btn>
                  <v-btn size="x-small" class="ml-4" @click="isEdit = false">cancel</v-btn>
                </div>
              </VForm>
            </div>
            <SchemaNode
              v-for="(el, i) in nodeProps.val"
              :key="`${nodeProps.level + 1}`"
              :label="el[0]"
              :val="el[1]"
              :level="nodeProps.level + 1"
            />
          </div>
        </div>
      </v-card>
    </template>
  </v-hover>
</template>

<style scoped>
.closing_bracket:after {
  content: "}";
  margin-left: 16px;
}
.opening_bracket:after {
  content: "{";
  margin-left: 16px;
}
</style>
