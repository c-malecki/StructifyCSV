<script lang="ts" setup>
import { ImportSchema, ExportSchema } from "../../../wailsjs/go/main/App";
import { convertMaptoObject, convertObjectToMap } from "./schemaEditor.util";
import { reactive, ref } from "vue";
import SchemaTree from "./components/SchemaTree.vue";
import SchemaInfoForm from "./components/SchemaInfoForm.vue";
import type { SchemaInfo } from "./schemaEditor.types";

const isEdit = ref(false);
const schemaTreeRef = ref<typeof SchemaTree | null>(null);

const schemaInfo = reactive<SchemaInfo>({
  title: "Product Example Schema",
  description: "Just a test schema to build out the editor UI with.",
});

const handleNewSchema = () => {
  if (schemaTreeRef.value) {
    schemaTreeRef.value.nodeMap = new Map<string, string | Map<string, any>>();
    schemaInfo.title = "New Schema";
    schemaInfo.description =
      "To change the name and description of this Schema, use the EDIT button to the right. \nTo begin building your Schema, click the ADD button below.";
  }
};

const handleUpdateSchema = ({ title, description }: SchemaInfo) => {
  schemaInfo.title = title;
  schemaInfo.description = description;
  isEdit.value = false;
};

const handleImportSchema = () => {
  if (!schemaTreeRef.value) return;
  ImportSchema()
    .then((res) => {
      schemaInfo.title = res.title;
      schemaInfo.description = res.description;
      schemaTreeRef.value!.nodeMap = convertObjectToMap(res.properties);
    })
    .catch(() => {});
};

const handleExportSchema = () => {
  if (!schemaTreeRef.value) return;
  ExportSchema({
    ...schemaInfo,
    properties: convertMaptoObject(schemaTreeRef.value!.nodeMap),
  })
    .then(() => {})
    .catch((err) => {});
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
          <h3>{{ schemaInfo.title }}</h3>
          <p>{{ schemaInfo.description }}</p>
          <v-btn position="absolute" size="small" location="top right" @click="isEdit = true">edit</v-btn>
        </div>

        <SchemaInfoForm
          v-else
          :schema-info="schemaInfo"
          @close-form="isEdit = false"
          @update-schema="handleUpdateSchema"
        />
      </v-sheet>
      <SchemaTree ref="schemaTreeRef" />
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
