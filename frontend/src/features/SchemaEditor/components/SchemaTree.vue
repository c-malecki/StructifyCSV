<script lang="ts" setup>
import { ref, reactive } from "vue";
import SchemaNode from "./SchemaNode.vue";
import { dataTypes } from "../schemaEditor.util";
import type { MapValue } from "../schemaEditor.types";
import type { VForm } from "vuetify/lib/components/index.mjs";

type FormControl = {
  showAdd: boolean;
  keyRules: ((val: string) => string | boolean)[];
};
const formControl = reactive<FormControl>({
  showAdd: false,
  keyRules: [(val: string) => val.length > 0 || "Key name is required."],
});
const formRef = ref<VForm | null>(null);

const location = new Map<string, string | Map<any, any>>([
  ["name", "string"],
  ["longitude", "number"],
  ["latitude", "number"],
]);

const member = new Map<string, string | Map<any, any>>([["id", "string"]]);

const user = new Map<string, string | Map<any, any>>([
  ["id", "string"],
  ["member", member],
]);

const testMap = new Map<string, string | Map<any, any>>([
  ["first_name", "string"],
  ["last_name", "string"],
  ["email", "string"],
  ["location", location],
  ["user", user],
]);

const newProperty = reactive({
  key: "",
  value: "string",
});

// const nodeMap = ref(new Map<string, string | Map<any, any>>());
const nodeMap = ref(testMap);

const resetAddProperty = () => {
  formControl.showAdd = false;
  newProperty.key = "";
  newProperty.value = "string";
};

const addProperty = () => {
  formRef.value!.validate().then(({ valid }) => {
    if (valid) {
      nodeMap.value.set(newProperty.key, newProperty.value === "object" ? new Map() : newProperty.value);
      resetAddProperty();
    }
  });
};

const handleDeleteProperty = (keyToDelete: string) => {
  const propertyType = typeof nodeMap.value.get(keyToDelete);
  if (propertyType === "object") {
    if (
      confirm(`Deleting "${keyToDelete}" will also delete any descendents of "${keyToDelete}." Do you wish to proceed?`)
    ) {
      nodeMap.value.delete(keyToDelete);
    }
  } else {
    if (confirm(`Deleting "${keyToDelete}" cannot be undone. Do you wish to proceed?`)) {
      nodeMap.value.delete(keyToDelete);
    }
  }
};

const handleAddPropertyToNode = (key: string, value: MapValue) => {
  nodeMap.value.set(key, value);
};

const updateAfterDelete = (keyToUpdate: string, value: MapValue) => {
  nodeMap.value.set(keyToUpdate, value);
};

const handleUpdateProperty = (key: string, value: string, originalKey: string) => {
  if (key !== originalKey) {
    nodeMap.value.delete(originalKey);
  }
  nodeMap.value.set(key, value === "object" ? new Map() : value);
};
</script>

<template>
  <v-sheet border class="schema_tree pa-2">
    <SchemaNode
      v-for="(node, i) in nodeMap"
      :key="`1-${i}-${typeof node[1]}`"
      :nodeKey="node[0]"
      :nodeValue="node[1]"
      :level="1"
      @updateParentNode="handleUpdateProperty"
      @addPropertyToNode="handleAddPropertyToNode"
      @deletePropertyFromNode="handleDeleteProperty"
      @updateAfterDelete="updateAfterDelete"
    />
    <div>
      <v-btn v-if="!formControl.showAdd" size="x-small" @click="formControl.showAdd = true" class="ml-2"> add </v-btn>
      <VForm v-else @submit.prevent="addProperty" ref="formRef">
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
    </div>
  </v-sheet>
</template>

<style scoped>
.schema_tree {
  font-size: 1.2em;
}
.schema_tree:before {
  content: "{";
}
.schema_tree:after {
  content: "}";
}
</style>
