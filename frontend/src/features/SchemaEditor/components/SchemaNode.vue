<script lang="ts" setup>
import { reactive, ref, toRef, inject, type PropType } from "vue";
import { computed } from "vue";
import type { DataTypes, Node, UpdateValueArgs } from "./SchemaTree.vue";

const emit = defineEmits(["updateParentNode", "updateNodeValue"]);
const dataTypes: DataTypes[] = inject("dataTypes")!;

const nodeProps = defineProps({
  label: {
    type: String,
    required: true,
  },
  val: {
    type: [String, Object],
    required: true,
  },
  level: {
    type: Number,
    default: 1,
  },
});

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

const showAdd = ref(false);
const showEdit = ref(false);

const copyVal = ref(nodeProps.val);
const editProperty = reactive({
  key: nodeProps.label,
  value: typeof nodeProps.val === "object" ? "object" : nodeProps.val,
});

const cancelEdit = () => {
  showEdit.value = false;
  editProperty.key = nodeProps.label;
  editProperty.value = typeof nodeProps.val;
};

const handleUpdateProperty = () => {
  // no change
  if (editProperty.value === copyVal.value && editProperty.key === nodeProps.label) {
    cancelEdit();
  }
  // if value is object and copyVal is object, don't change the properties
  if (editProperty.value === "object" && typeof copyVal.value === "object") {
    emit("updateParentNode", editProperty.key, copyVal.value, nodeProps.label);
    showEdit.value = false;
  } else {
    emit("updateParentNode", editProperty.key, editProperty.value, nodeProps.label);
    showEdit.value = false;
  }
};

const handleUpdateParentNode = (key: string, value: string, originalKey: string) => {
  if (typeof copyVal.value === "object") {
    if (key !== originalKey) {
      (copyVal.value as Node).delete(originalKey);
    }
    (copyVal.value as Node).set(key, value === "object" ? new Map() : value);
    emit("updateParentNode", nodeProps.label, copyVal.value, nodeProps.label);
  } else {
    emit("updateParentNode", editProperty.key, editProperty.value, nodeProps.label);
  }
};
// const handleAddProperty = () => {
//   showAdd.value = false;
//   p.set(editProperty.key, editProperty.value === "object" ? new Map() : editProperty.value);
//   editProperty.key = "";
//   editProperty.value = "string";
// };
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
            <p v-if="!showEdit">
              <b>{{ label }}:</b> {{ val }}
            </p>
            <div v-if="isHovering && !showEdit" class="d-flex">
              <v-btn size="x-small" class="ml-4" @click="showEdit = true">edit</v-btn>
              <v-btn size="x-small" class="ml-4">delete</v-btn>
            </div>

            <VForm v-if="showEdit">
              <div class="d-flex">
                <VTextField v-model="editProperty.key" label="Property Name" style="width: 200px" />
                <VSelect
                  v-model="editProperty.value"
                  :items="dataTypes"
                  item-title="name"
                  item-value="value"
                  style="width: 200px"
                />
                <v-btn size="x-small" class="ml-4" @click="handleUpdateProperty">save</v-btn>
                <v-btn size="x-small" class="ml-4" @click="cancelEdit">cancel</v-btn>
              </div>
            </VForm>
          </div>

          <div v-else>
            <div class="d-flex opening_bracket pb-1">
              <p>
                <b>{{ label }}:</b>
              </p>
              <div v-if="isHovering && !showEdit" class="d-flex">
                <v-btn size="x-small" class="ml-4">add property</v-btn>
                <v-btn size="x-small" class="ml-4" @click="showEdit = true">edit</v-btn>
                <v-btn size="x-small" class="ml-4">delete</v-btn>
              </div>
              <VForm v-if="showEdit">
                <div class="d-flex">
                  <VTextField v-model="editProperty.key" label="Property Name" style="width: 200px" />
                  <VSelect
                    v-model="editProperty.value"
                    :items="dataTypes"
                    item-title="name"
                    item-value="value"
                    style="width: 200px"
                  />
                  <v-btn size="x-small" class="ml-4" @click="handleUpdateProperty">save</v-btn>
                  <v-btn size="x-small" class="ml-4" @click="cancelEdit">cancel</v-btn>
                </div>
              </VForm>
            </div>
            <SchemaNode
              v-for="(el, i) in nodeProps.val"
              :key="`${nodeProps.level + 1}-${i}-${typeof el[1]}`"
              :label="el[0]"
              :val="el[1]"
              :level="nodeProps.level + 1"
              @updateParentNode="handleUpdateParentNode"
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
