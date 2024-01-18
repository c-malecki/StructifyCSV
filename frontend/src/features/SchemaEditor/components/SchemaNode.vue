<script lang="ts" setup>
import { reactive, ref, type PropType } from "vue";
import { computed } from "vue";
import { getHoverColorScheme, getLeftIndent, dataTypes } from "../../../util";
import { type VForm } from "vuetify/lib/components/index.mjs";

type MapValue = Map<string, string | Map<any, any>>;

type FormControl = {
  showAdd: boolean;
  showEdit: boolean;
  keyRules: ((val: string) => string | boolean)[];
};

const nodeProps = defineProps({
  nodeKey: {
    type: String,
    required: true,
  },
  nodeValue: {
    type: [String, Object] as PropType<string | MapValue>,
    required: true,
  },
  level: {
    type: Number,
    default: 1,
  },
});
const emit = defineEmits(["updateParentNode", "addPropertyToNode"]);

const leftIndent = computed(() => getLeftIndent(nodeProps.level));
const colorScheme = computed(() => getHoverColorScheme(nodeProps.level));

const addFormRef = ref<VForm | null>(null);
const editFormRef = ref<VForm | null>(null);

const formControl = reactive<FormControl>({
  showAdd: false,
  showEdit: false,
  keyRules: [(val: string) => val.length > 0 || "Key name is required."],
});

const copyNodeValue = ref(nodeProps.nodeValue);
const newProperty = reactive({
  key: "",
  value: "string",
});

const editProperty = reactive({
  key: nodeProps.nodeKey,
  value: typeof nodeProps.nodeValue === "object" ? "object" : nodeProps.nodeValue,
});

const resetEditProperty = () => {
  formControl.showEdit = false;
  editProperty.key = nodeProps.nodeKey;
  editProperty.value = typeof nodeProps.nodeValue;
};

const resetAddProperty = () => {
  formControl.showAdd = false;
  newProperty.key = "";
  newProperty.value = "string";
};

const handleUpdateProperty = () => {
  // no change
  if (editProperty.value === copyNodeValue.value && editProperty.key === nodeProps.nodeKey) {
    resetEditProperty();
  }
  // if value is object and copyNodeValue is object, don't change the properties
  if (editProperty.value === "object" && typeof copyNodeValue.value === "object") {
    emit("updateParentNode", editProperty.key, copyNodeValue.value, nodeProps.nodeKey);
    formControl.showEdit = false;
  } else {
    emit("updateParentNode", editProperty.key, editProperty.value, nodeProps.nodeKey);
    formControl.showEdit = false;
  }
};

const addProperty = () => {
  (copyNodeValue.value as MapValue).set(
    newProperty.key,
    newProperty.value === "object" ? new Map() : newProperty.value
  );
  emit("addPropertyToNode", nodeProps.nodeKey, copyNodeValue.value);
  resetAddProperty();
};

const handleAddProperty = (key: string, value: MapValue) => {
  (copyNodeValue.value as MapValue).set(key, value);
  emit("addPropertyToNode", nodeProps.nodeKey, copyNodeValue.value);
};

const handleUpdateParentNode = (key: string, value: string, originalKey: string) => {
  if (typeof copyNodeValue.value === "object") {
    if (key !== originalKey) {
      copyNodeValue.value.delete(originalKey);
    }
    copyNodeValue.value.set(key, value === "object" ? new Map() : value);
    emit("updateParentNode", nodeProps.nodeKey, copyNodeValue.value, nodeProps.nodeKey);
  } else {
    emit("updateParentNode", editProperty.key, editProperty.value, nodeProps.nodeKey);
  }
};
</script>

<template>
  <v-hover>
    <template v-slot:default="{ isHovering, props }">
      <v-card v-bind="props" :color="isHovering ? colorScheme.hover : undefined" variant="flat" :style="leftIndent">
        <div
          :class="{ closing_bracket: typeof nodeValue === 'object' }"
          :style="`color: ${isHovering ? colorScheme.font : 'black'}`"
          class="pt-1 pb-1"
        >
          <div v-if="typeof nodeValue === 'string'" class="d-flex">
            <p v-if="!formControl.showEdit">
              <b>{{ nodeKey }}:</b> {{ nodeValue }}
            </p>
            <div v-if="isHovering && !formControl.showEdit" class="d-flex">
              <v-btn size="x-small" class="ml-4" @click="formControl.showEdit = true">edit</v-btn>
              <v-btn size="x-small" class="ml-4">delete</v-btn>
            </div>

            <VForm v-if="formControl.showEdit" ref="editFormRef">
              <div class="d-flex">
                <VTextField
                  v-model="editProperty.key"
                  label="Property Name"
                  :rules="formControl.keyRules"
                  style="width: 200px"
                />
                <VSelect
                  v-model="editProperty.value"
                  :items="dataTypes"
                  item-title="name"
                  item-value="value"
                  style="width: 200px"
                />
                <v-btn size="x-small" class="ml-4" @click="handleUpdateProperty">save</v-btn>
                <v-btn size="x-small" class="ml-4" @click="resetEditProperty">cancel</v-btn>
              </div>
            </VForm>
          </div>

          <div v-else>
            <div class="d-flex opening_bracket pb-1">
              <p v-if="!formControl.showEdit">
                <b>{{ nodeKey }}:</b>
              </p>
              <div v-if="isHovering && !formControl.showEdit" class="d-flex">
                <v-btn size="x-small" class="ml-4" @click="formControl.showAdd = true">add</v-btn>
                <v-btn size="x-small" class="ml-4" @click="formControl.showEdit = true">edit</v-btn>
                <v-btn size="x-small" class="ml-4">delete</v-btn>
              </div>
              <VForm v-if="formControl.showEdit" ref="editFormRef">
                <div class="d-flex">
                  <VTextField
                    v-model="editProperty.key"
                    label="Property Name"
                    :rules="formControl.keyRules"
                    style="width: 200px"
                  />
                  <VSelect
                    v-model="editProperty.value"
                    :items="dataTypes"
                    item-title="name"
                    item-value="value"
                    style="width: 200px"
                  />
                  <v-btn size="x-small" class="ml-4" @click="handleUpdateProperty">save</v-btn>
                  <v-btn size="x-small" class="ml-4" @click="resetEditProperty">cancel</v-btn>
                </div>
              </VForm>
            </div>
            <VForm v-if="formControl.showAdd" @submit.prevent="addProperty" ref="addFormRef">
              <div class="d-flex">
                <VTextField
                  v-model="newProperty.key"
                  label="Property Name"
                  :rules="formControl.keyRules"
                  style="max-width: 200px"
                />
                <VSelect
                  v-model="newProperty.value"
                  :items="dataTypes"
                  item-title="name"
                  item-value="value"
                  style="max-width: 200px"
                />
                <v-btn type="button" size="x-small" class="ml-4" @click="resetAddProperty">cancel</v-btn>
                <v-btn type="submit" size="x-small" class="ml-4">save</v-btn>
              </div>
            </VForm>
            <SchemaNode
              v-for="(node, i) in nodeValue"
              :key="`${nodeProps.level + 1}-${i}-${typeof node[1]}`"
              :nodeKey="node[0]"
              :nodeValue="node[1]"
              :level="nodeProps.level + 1"
              @updateParentNode="handleUpdateParentNode"
              @addPropertyToNode="handleAddProperty"
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
