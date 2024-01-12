<script lang="ts" setup>
import { computed } from "vue";
import { useStore } from "../../store/store";
import { formatSchemaJson } from "../util/format";
import NewSchemaForm from "./components/NewSchemaForm.vue";
import EditSchemaForm from "./components/EditSchemaForm.vue";

const store = useStore();
const curForm = computed(() => store.schemaEditor.curForm);
const header = computed(() => {
  switch (curForm.value) {
    case "new":
      return "New Schema";
    case "edit":
      return "Edit Schema";
    case null:
      return `Schema: ${store.model!.schemas[store.schemaEditor.curSelected].name}`;
  }
});
const schemaOpts = computed(() =>
  store.model?.schemas.map((s, i) => {
    return {
      name: s.name,
      value: i,
    };
  })
);
</script>

<template>
  <v-card rounded="0">
    <v-toolbar density="compact" color="grey-lighten-2">
      <v-menu>
        <template v-slot:activator="{ props }">
          <v-app-bar-nav-icon v-bind="props"></v-app-bar-nav-icon>
        </template>

        <v-list variant="flat" border density="compact" elevation="0">
          <v-list-item density="compact" title="New Schema" @click="store.changeSchemaEditorForm('new')" />
          <v-list-item density="compact" title="Import Schema" @click="store.importModel()" />
          <v-list-item density="compact" title="Export Schema" @click="store.exportModel()" />
        </v-list>
      </v-menu>

      <v-toolbar-title> {{ header }}</v-toolbar-title>
    </v-toolbar>

    <v-card-text>
      <NewSchemaForm v-if="curForm === 'new'" @close-form="store.changeSchemaEditorForm(null)" />

      <div v-if="!curForm">
        <VSelect
          v-model="store.schemaEditor.curSelected"
          label="Schemas"
          :items="schemaOpts"
          item-title="name"
          item-value="value"
        />

        <div class="d-flex">
          <h3>{{ store.model!.schemas[store.schemaEditor.curSelected!].name.replaceAll(" ", "_").toLowerCase() }}</h3>
          <v-btn @click="store.changeSchemaEditorForm('edit')" size="x-small">edit</v-btn>
        </div>

        <v-sheet border class="pa-2">
          <pre>{{
            JSON.stringify(formatSchemaJson(store.model!.schemas[store.schemaEditor.curSelected!]), null, 2).replaceAll(
              '"',
              ""
            )
          }}</pre>
        </v-sheet>
      </div>

      <EditSchemaForm v-if="curForm === 'edit'" @close-form="store.changeSchemaEditorForm(null)" />
    </v-card-text>
  </v-card>
</template>

<style scoped>
pre {
  padding: 0.2rem;
  border: 1px solid black;
}
</style>
