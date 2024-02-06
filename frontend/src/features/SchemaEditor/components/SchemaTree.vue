<script lang="ts" setup>
import { ref, inject } from "vue";
import SchemaNode from "./SchemaNode.vue";
import AddPropertyForm from "./AddPropertyForm.vue";
import {
  JsonSchemaKey,
  type JsonSchemaDataType,
  type SchemaPropertiesMap,
  type PropertiesMapValue,
} from "../../../types/editor.types";

const jsonSchema = inject(JsonSchemaKey);
if (!jsonSchema) {
  throw new Error(`Could not resolve ${JsonSchemaKey.description}`);
}

const showAddForm = ref(false);

const updateBaseKey = ({
  newKey,
  oldKey,
  sameValue,
}: {
  newKey: string;
  oldKey: string;
  sameValue: PropertiesMapValue;
}) => {
  jsonSchema.properties.delete(oldKey);
  jsonSchema.properties.set(newKey, sameValue);
};

const updateBaseValue = ({
  sameKey,
  newValue,
}: {
  sameKey: string;
  newValue: JsonSchemaDataType | SchemaPropertiesMap;
}) => {
  jsonSchema.properties.set(sameKey, newValue);
};

const deleteBaseProperty = (keyToDelete: string) => {
  // will need to account for array and null?
  const isObject = typeof jsonSchema.properties.get(keyToDelete) === "object";
  let message = "";
  switch (isObject) {
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
</script>

<template>
  <div class="schema_tree pa-4">
    <SchemaNode
      v-for="(node, i) in jsonSchema.properties"
      :key="`1-${i}-${typeof node[1]}`"
      :node="node"
      :nodeKey="node[0]"
      :nodeValue="node[1]"
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
        @hideForm="showAddForm = false"
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
