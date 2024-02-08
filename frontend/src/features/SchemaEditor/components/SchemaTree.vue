<script lang="ts" setup>
import { ref, inject } from "vue";
import SchemaNode from "./SchemaNode/SchemaNode.vue";
import AddPropertyForm from "./AddPropertyForm.vue";
import { JsonSchemaKey } from "../../../types/editor.types";
import { type SchemaProperty } from "../../../types/properties.types";

const jsonSchema = inject(JsonSchemaKey);
if (!jsonSchema) {
  throw new Error(`Could not resolve ${JsonSchemaKey.description}`);
}

const showAddForm = ref(false);

const updateBaseKey = ({
  editKey,
  curKey,
  value,
}: {
  editKey: string;
  curKey: string;
  value: SchemaProperty;
}) => {
  jsonSchema.properties.delete(curKey);
  jsonSchema.properties.set(editKey, value);
};

const updateBaseValue = ({
  editKey,
  curKey,
  value,
}: {
  editKey: string;
  curKey: string;
  value: SchemaProperty;
}) => {
  if (editKey !== curKey) {
    jsonSchema.properties.delete(curKey);
  }
  jsonSchema.properties.set(editKey, value);
};

const deleteBaseProperty = (keyToDelete: string) => {
  const property = jsonSchema.properties.get(keyToDelete);
  const isObjorArr = property!.type === "object" || property!.type === "array";
  let message = "";
  switch (isObjorArr) {
    case true:
      message = `Deleting "${keyToDelete}" will also delete any descendents of "${keyToDelete}." Do you wish to proceed?`;
      break;
    case false:
      message = `Deleting "${keyToDelete}" cannot be undone. Do you wish to proceed?`;
      break;
  }
  if (confirm(message)) {
    jsonSchema.properties.delete(keyToDelete);
  }
};

const addNewProperty = ({
  key,
  value,
}: {
  key: string;
  value: SchemaProperty;
}) => {
  jsonSchema.properties.set(key, value);
};
</script>

<template>
  <div class="schema_tree pa-4">
    <SchemaNode
      v-for="(node, i) in jsonSchema.properties"
      :key="`1-${i}-${typeof node[1]}`"
      :node="node"
      :level="1"
      @update-base-key="updateBaseKey"
      @update-base-value="updateBaseValue"
      @delete-base-property="deleteBaseProperty"
    />
    <div>
      <v-btn
        v-if="!showAddForm"
        size="x-small"
        @click="showAddForm = true"
        class="ml-2"
      >
        add
      </v-btn>
      <AddPropertyForm
        v-else
        :nodeValue="jsonSchema.properties"
        @close-form="showAddForm = false"
        @add-new-property="addNewProperty"
      />
    </div>
  </div>
</template>

<style scoped>
.schema_tree:before {
  content: "{";
}
.schema_tree:after {
  content: "}";
}
</style>
