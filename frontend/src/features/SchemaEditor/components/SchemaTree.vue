<script lang="ts" setup>
import { ref, reactive, provide } from "vue";
import SchemaNode from "./SchemaNode.vue";

export type DataTypes = { name: string; value: string };
export type Node = Map<string, string | Map<any, any>>;
export type UpdateValueArgs = { key: string; value: string; originalKey: string };
export type UpdateNode = ({ key, value, originalKey }: UpdateValueArgs) => void;

const dataTypes: DataTypes[] = [
  { name: "string (text)", value: "string" },
  { name: "float (decimal)", value: "number" },
  { name: "integer", value: "integer" },
  { name: "object", value: "object" },
  // { name: "array", value: "array" },
  { name: "boolean", value: "boolean" },
  { name: "null", value: "null" },
];
provide("dataTypes", dataTypes);

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

const editNode = reactive({
  key: "",
  value: "string",
});
const showAdd = ref(false);
// const nodeMap = ref(new Map<string, string | Map<any, any>>());
const nodeMap = ref(testMap);

const handleAddProperty = () => {
  showAdd.value = false;
  nodeMap.value.set(editNode.key, editNode.value === "object" ? new Map() : editNode.value);
  editNode.key = "";
  editNode.value = "string";
};
const updateNode = (key: string, value: string, originalKey: string) => {
  if (key !== originalKey) {
    nodeMap.value.delete(originalKey);
  }
  console.log(value);
  nodeMap.value.set(key, value === "object" ? new Map() : value);
};
// const updateNode: UpdateNode = ({ key, value, originalKey }) => {
//   if (key !== originalKey) {
//     nodeMap.value.delete(originalKey);
//   }
//   nodeMap.value.set(key, value === "object" ? new Map() : value);
// };
</script>

<template>
  <v-sheet border class="schema_tree pa-2">
    <SchemaNode
      v-for="(el, i) in nodeMap"
      :key="`1-${i}-${typeof el[1]}`"
      :label="el[0]"
      :val="el[1]"
      :level="1"
      @updateParentNode="updateNode"
    />
    <div>
      <v-btn v-if="!showAdd" size="x-small" @click="showAdd = true" class="ml-2">add property</v-btn>
      <VForm v-else>
        <div class="d-flex">
          <VTextField v-model="editNode.key" label="Property Name" style="max-width: 200px" />
          <VSelect
            v-model="editNode.value"
            :items="dataTypes"
            item-title="name"
            item-value="value"
            style="max-width: 200px"
          />
          <v-btn size="x-small" class="ml-4" @click="showAdd = false">cancel</v-btn>
          <v-btn size="x-small" class="ml-4" @click="handleAddProperty">save</v-btn>
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
