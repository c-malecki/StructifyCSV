<script lang="ts" setup>
import { ref, inject } from "vue";
import SchemaNode from "./SchemaNode/SchemaNode.vue";
import AddPropertyForm from "./AddPropertyForm.vue";
import { JsonSchemaKey } from "../SchemaEditor.types";
import { entity } from "../../../../wailsjs/go/models";

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
  value: entity.Schema;
}) => {
  delete jsonSchema.value.properties[curKey];
  jsonSchema.value.properties[editKey] = value;
};

const updateBaseValue = ({
  editKey,
  curKey,
  value,
}: {
  editKey: string;
  curKey: string;
  value: entity.Schema;
}) => {
  if (editKey !== curKey) {
    delete jsonSchema.value.properties[curKey];
  }
  jsonSchema.value.properties[editKey] = value;
};

const deleteBaseProperty = (keyToDelete: string) => {
  const property = jsonSchema.value.properties[keyToDelete];
  const isObjOrArr = property!.type === "object" || property!.type === "array";
  let message = "";
  switch (isObjOrArr) {
    case true:
      message = `Deleting "${keyToDelete}" will also delete any descendents of "${keyToDelete}." Do you wish to proceed?`;
      break;
    case false:
      message = `Deleting "${keyToDelete}" cannot be undone. Do you wish to proceed?`;
      break;
  }
  if (confirm(message)) {
    delete jsonSchema.value.properties[keyToDelete];
  }
};

const addNewProperty = ({
  key,
  value,
}: {
  key: string;
  value: entity.Schema;
}) => {
  jsonSchema.value.properties[key] = value;
};
</script>

<template>
  <div class="schema_tree pa-4">
    <SchemaNode
      v-for="([k, v], i) in Object.entries(jsonSchema.properties)"
      :key="`1-${i}-${k}`"
      :node="[k, v]"
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
