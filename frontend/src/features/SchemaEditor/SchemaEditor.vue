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

const schemaValues = reactive<FormValues>({
  title: "Test Schema",
  description: "Just a test schema to build out the editor UI with.",
});
const formControl: FormControl = {
  titleRules: [
    (v: string) => v.length > 0 || "Schema Name is required.",
    (v: string) => [...v].length <= 150 || "Schema Name cannot be longer than 150 characters.",
  ],
  descriptionRules: [(v: string) => [...v].length <= 1000 || "Description cannot be longer than 1000 characters."],
};

const clearValues = () => {
  schemaValues.title = "";
  schemaValues.description = "";
};

const handleSubmit = () => {};
const handleImportSchema = () => {
  ImportSchema()
    .then(() => {})
    .catch((err) => {});
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
          <v-list-item density="compact" title="New Schema" @click="clearValues" />
          <v-list-item density="compact" title="Import Schema" @click="handleImportSchema" />
          <v-list-item density="compact" title="Export Schema" @click="handleExportSchema" />
        </v-list>
      </v-menu>

      <v-toolbar-title> Schema Editor</v-toolbar-title>
    </v-toolbar>

    <v-card-text>
      <v-sheet border class="pa-2 mb-2">
        <div v-if="!isEdit" class="relative">
          <h3>{{ schemaValues.title }}</h3>
          <p>{{ schemaValues.description }}</p>
          <v-btn position="absolute" size="small" location="top right" @click="isEdit = true">edit</v-btn>
        </div>

        <VForm @submit.prevent="handleSubmit" v-else>
          <VTextField
            v-model="schemaValues.title"
            label="Schema Name"
            :rules="formControl.titleRules"
            :counter="150"
            persistent-counter
          />
          <VTextarea
            v-model="schemaValues.description"
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
      <SchemaTree />
    </v-card-text>
  </v-card>
</template>

<style scoped>
.relative {
  position: relative;
}
</style>
