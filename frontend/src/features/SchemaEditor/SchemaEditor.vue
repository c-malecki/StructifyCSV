<script lang="ts" setup>
import { ImportSchema, ExportSchema } from "../../../wailsjs/go/main/App";
import { reactive, ref } from "vue";
import SchemaTree from "./components/SchemaTree.vue";

export type FormValues = {
  title: string;
  description: string;
};
type FormControl = {
  titleRules: ((val: string) => string | boolean)[];
  descriptionRules: ((val: string) => string | boolean)[];
};

const isEdit = ref(false);
const schemaTree = ref<typeof SchemaTree | null>(null);

const formValues = reactive<FormValues>({
  title: "Product Example Schema",
  description: "Just a test schema to build out the editor UI with.",
});
const formControl: FormControl = {
  titleRules: [
    (v: string) => v.length > 0 || "Schema Name is required.",
    (v: string) => [...v].length <= 150 || "Schema Name cannot be longer than 150 characters.",
  ],
  descriptionRules: [(v: string) => [...v].length <= 1000 || "Description cannot be longer than 1000 characters."],
};

const handleNewSchema = () => {
  if (schemaTree.value) {
    schemaTree.value.nodeMap = new Map<string, string | Map<any, any>>();
    formValues.title = "New Schema";
    formValues.description =
      "To change the name and description of this Schema, use the EDIT button to the right. \nTo begin building your Schema, click the ADD button below.";
  }
};

const handleSubmit = () => {};
const handleImportSchema = () => {
  ImportSchema()
    .then(() => {})
    .catch(() => {});
};
const handleExportSchema = () => {
  // ExportSchema()
  // .then(() => {})
  // .catch((err) => {})
};
</script>

<template>
  <v-card rounded="0">
    <v-toolbar density="compact" color="grey-lighten-2">
      <v-menu>
        <template v-slot:activator="{ props }">
          <v-app-bar-nav-icon v-bind="props"></v-app-bar-nav-icon>
        </template>

        <v-list variant="flat" border density="compact" elevation="0">
          <v-list-item density="compact" title="New Schema" @click="handleNewSchema" />
          <v-list-item density="compact" title="Import Schema" @click="handleImportSchema" />
          <v-list-item density="compact" title="Export Schema" @click="handleExportSchema" />
        </v-list>
      </v-menu>

      <v-toolbar-title>Schema Editor</v-toolbar-title>
    </v-toolbar>

    <v-card-text>
      <v-sheet border class="pa-2 mb-2">
        <div v-if="!isEdit" class="relative">
          <h3>{{ formValues.title }}</h3>
          <p>{{ formValues.description }}</p>
          <v-btn position="absolute" size="small" location="top right" @click="isEdit = true">edit</v-btn>
        </div>

        <VForm @submit.prevent="handleSubmit" v-else>
          <VTextField
            v-model="formValues.title"
            label="Schema Name"
            :rules="formControl.titleRules"
            :counter="150"
            persistent-counter
          />
          <VTextarea
            v-model="formValues.description"
            label="Description"
            :rules="formControl.descriptionRules"
            :counter="1000"
            persistent-counter
          />
          <div class="d-flex mt-2">
            <v-btn type="button" size="small" @click="isEdit = false" class="ml-auto mr-4">cancel</v-btn>
            <v-btn type="submit" size="small">save</v-btn>
          </div>
        </VForm>
      </v-sheet>
      <SchemaTree ref="schemaTree" />
    </v-card-text>
  </v-card>
</template>

<style scoped>
p {
  white-space: pre;
}
.relative {
  position: relative;
}
</style>
