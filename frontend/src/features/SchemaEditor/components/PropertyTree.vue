<script lang="ts" setup>
import { ref } from "vue";
import { useSchemaStore } from "../../../store/schema";
import PropertyNode from "./PropertyNode/PropertyNode.vue";
import AddPropertyForm from "./forms/AddPropertyForm.vue";
import EditObjectRequiredForm from "./forms/EditObjectRequiredForm.vue";
import { entity } from "../../../../wailsjs/go/models";

const schemaStore = useSchemaStore();

type CurForm = "add" | "required" | null;
const curForm = ref<CurForm>(null);

const updateBaseKey = ({
  editKey,
  curKey,
  value,
}: {
  editKey: string;
  curKey: string;
  value: entity.PropertySchema;
}) => {
  delete schemaStore.jsonSchema.properties[curKey];
  schemaStore.jsonSchema.properties[editKey] = value;
};

const updateBaseValue = ({
  editKey,
  curKey,
  value,
}: {
  editKey: string;
  curKey: string;
  value: entity.PropertySchema;
}) => {
  if (editKey !== curKey) {
    delete schemaStore.jsonSchema.properties[curKey];
  }
  schemaStore.jsonSchema.properties[editKey] = value;
};

const deleteBaseProperty = (keyToDelete: string) => {
  const property = schemaStore.jsonSchema.properties[keyToDelete];
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
    delete schemaStore.jsonSchema.properties[keyToDelete];
  }
};

const addNewProperty = ({
  key,
  value,
}: {
  key: string;
  value: entity.PropertySchema;
}) => {
  schemaStore.jsonSchema.properties[key] = value;
};

const updateRequired = (required: string[]) => {
  schemaStore.jsonSchema.required = required;
  curForm.value = null;
};
</script>

<template>
  <v-sheet border>
    <div class="tree pa-4">
      <PropertyNode
        v-for="[k, v] in Object.entries(schemaStore.jsonSchema.properties)"
        :key="`1-json-${k}`"
        :node="[k, v]"
        :level="1"
        @update-base-key="updateBaseKey"
        @update-base-value="updateBaseValue"
        @delete-base-property="deleteBaseProperty"
      />
      <div>
        <div v-if="!curForm" class="d-flex">
          <v-btn
            size="x-small"
            @click="curForm = 'add'"
            class="ml-2"
            prepend-icon="mdi-plus"
          >
            add
          </v-btn>
          <v-btn
            size="x-small"
            @click="curForm = 'required'"
            class="ml-4"
            prepend-icon="mdi-pencil-box-outline"
          >
            required
          </v-btn>
        </div>

        <AddPropertyForm
          v-if="curForm === 'add'"
          :nodeValue="schemaStore.jsonSchema.properties"
          @close-form="curForm = null"
          @add-new-property="addNewProperty"
        />
        <EditObjectRequiredForm
          v-if="curForm === 'required'"
          :objectProperty="schemaStore.jsonSchema"
          @close-form="curForm = null"
          @update-required="updateRequired"
        />
      </div>
    </div>
  </v-sheet>
</template>

<style scoped>
.tree:before {
  content: "{";
}
.tree:after {
  content: "}";
}
</style>
