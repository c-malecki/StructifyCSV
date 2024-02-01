<script lang="ts" setup>
import { ref, reactive, inject } from "vue";
import SchemaNode from "./SchemaNode.vue";
import { dataTypeOpts, SchemaValuesKey } from "../schemaEditor.util";
import type { PropertiesMap } from "../schemaEditor.types";
import type { VForm } from "vuetify/lib/components/index.mjs";

type FormControl = {
  showAdd: boolean;
  keyRules: ((val: string) => string | boolean)[];
};

const schemaValues = inject(SchemaValuesKey);
if (!schemaValues) {
  throw new Error(`Could not resolve ${SchemaValuesKey.description}`);
}

const formRef = ref<VForm | null>(null);
const formControl = reactive<FormControl>({
  showAdd: false,
  keyRules: [(val: string) => val.length > 0 || "Property Name is required."],
});

const newProperty = reactive({
  key: "",
  value: "string",
});

const resetAddProperty = () => {
  formControl.showAdd = false;
  newProperty.key = "";
  newProperty.value = "string";
};

const addProperty = () => {
  formRef.value!.validate().then(({ valid }) => {
    if (valid) {
      schemaValues.properties.set(
        newProperty.key,
        newProperty.value === "object" ? new Map() : newProperty.value
      );
      resetAddProperty();
    }
  });
};

const handleDeleteProperty = (keyToDelete: string) => {
  const message =
    typeof schemaValues.properties.get(keyToDelete) === "object"
      ? `Deleting "${keyToDelete}" will also delete any descendents of "${keyToDelete}." Do you wish to proceed?`
      : `Deleting "${keyToDelete}" cannot be undone. Do you wish to proceed?`;
  if (confirm(message)) {
    schemaValues.properties.delete(keyToDelete);
  }
};

const handleAddPropertyToNode = (key: string, value: PropertiesMap) => {
  schemaValues.properties.set(key, value);
};

const updateAfterDelete = (keyToUpdate: string, value: PropertiesMap) => {
  schemaValues.properties.set(keyToUpdate, value);
};

const handleUpdateProperty = (
  key: string,
  value: string,
  originalKey: string
) => {
  if (key !== originalKey) {
    schemaValues.properties.delete(originalKey);
  }
  schemaValues.properties.set(key, value === "object" ? new Map() : value);
};
</script>

<template>
  <v-sheet border class="schema_tree pa-2">
    <SchemaNode
      v-for="(node, i) in schemaValues.properties"
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
      <v-btn
        v-if="!formControl.showAdd"
        size="x-small"
        @click="formControl.showAdd = true"
        class="ml-2"
      >
        add
      </v-btn>
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
            :items="dataTypeOpts"
            item-title="name"
            item-value="value"
            style="max-width: 200px"
          />
          <v-btn
            type="button"
            size="x-small"
            class="ml-4"
            @click="resetAddProperty"
          >
            cancel
          </v-btn>
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
