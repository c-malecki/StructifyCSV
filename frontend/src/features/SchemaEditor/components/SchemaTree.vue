<script lang="ts" setup>
import { ref, reactive } from "vue";
import SchemaNode from "./SchemaNode.vue";

type DataTypes = { name: string; value: string };

const dataTypes: DataTypes[] = [
  { name: "string (text)", value: "string" },
  { name: "float (decimal)", value: "number" },
  { name: "integer", value: "integer" },
  { name: "object", value: "object" },
  { name: "array", value: "array" },
  { name: "boolean", value: "boolean" },
  { name: "null", value: "null" },
];

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

const p = new Map<string, string | Map<any, any>>([
  ["first_name", "string"],
  ["last_name", "string"],
  ["email", "string"],
  ["location", location],
  ["user", user],
]);

const element = reactive({
  key: "",
  value: "string",
});
const showAdd = ref(false);
const properties = ref(p);
const handleAddProperty = () => {
  showAdd.value = false;
  p.set(element.key, element.value === "object" ? new Map() : element.value);
  element.key = "";
  element.value = "string";
};
</script>

<template>
  <v-sheet border class="schema_tree pa-2">
    <SchemaNode v-for="(el, i) in properties" :key="`1-${i}`" :label="el[0]" :val="el[1]" :level="1" />
    <div>
      <v-btn v-if="!showAdd" size="x-small" @click="showAdd = true" class="ml-2">add property</v-btn>
      <VForm v-else>
        <div class="d-flex">
          <VTextField v-model="element.key" label="Property Name" style="max-width: 200px" />
          <VSelect
            v-model="element.value"
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
