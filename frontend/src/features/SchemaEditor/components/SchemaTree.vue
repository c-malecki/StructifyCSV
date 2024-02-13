<script lang="ts" setup>
import { ref, inject } from "vue";
import SchemaNode from "./SchemaNode/SchemaNode.vue";
import AddPropertyForm from "./forms/AddPropertyForm.vue";
import EditRequiredForm from "./forms/EditRequiredForm.vue";
import { JsonSchemaKey } from "../SchemaEditor.types";
import { entity } from "../../../../wailsjs/go/models";

const jsonSchema = inject(JsonSchemaKey);
if (!jsonSchema) {
  throw new Error(`Could not resolve ${JsonSchemaKey.description}`);
}

const showAddForm = ref(false);
const showEditRequiredForm = ref(false);

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

const updateRequired = (required: string[]) => {
  jsonSchema.value.required = required;
  showEditRequiredForm.value = false;
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
      <div v-if="!showAddForm && !showEditRequiredForm" class="d-flex">
        <v-btn
          size="x-small"
          @click="showAddForm = true"
          class="ml-2"
          prepend-icon="mdi-plus"
        >
          add
        </v-btn>
        <v-btn
          size="x-small"
          @click="showEditRequiredForm = true"
          class="ml-4"
          prepend-icon="mdi-pencil-box-outline"
        >
          required
        </v-btn>
      </div>

      <AddPropertyForm
        v-if="showAddForm"
        :nodeValue="jsonSchema.properties"
        @close-form="showAddForm = false"
        @add-new-property="addNewProperty"
      />
      <EditRequiredForm
        v-if="showEditRequiredForm"
        :schema="jsonSchema"
        @close-form="showEditRequiredForm = false"
        @update-required="updateRequired"
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
