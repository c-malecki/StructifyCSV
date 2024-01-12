<script lang="ts" setup>
import { computed } from "vue";
import { useStore } from "../../store/store";
import { formatModelJson } from "../util/format";
import NewModelForm from "./components/NewModelForm.vue";
import EditModelForm from "./components/EditModelForm.vue";

const store = useStore();
const curForm = computed(() => store.modelEditor.curForm);
const header = computed(() => {
  switch (curForm.value) {
    case "new":
      return "New Model";
    case "edit":
      return "Edit Model";
    case null:
      return `Model: ${store.model ? store.model.name : "None Selected"}`;
  }
});
</script>

<template>
  <v-card rounded="0">
    <v-toolbar density="compact" color="grey-lighten-2">
      <v-menu>
        <template v-slot:activator="{ props }">
          <v-app-bar-nav-icon v-bind="props"></v-app-bar-nav-icon>
        </template>

        <v-list variant="flat" border density="compact" elevation="0">
          <v-list-item density="compact" title="New Model" @click="store.changeModelEditorForm('new')" />
          <v-list-item density="compact" title="Import Model" @click="store.importModel()" />
          <v-list-item density="compact" title="Export Model" @click="store.exportModel()" />
        </v-list>
      </v-menu>

      <v-toolbar-title>{{ header }}</v-toolbar-title>
    </v-toolbar>

    <v-card-text>
      <NewModelForm v-if="curForm === 'new'" @close-form="store.changeModelEditorForm(null)" />

      <div v-if="store.model">
        <div v-if="!curForm" class="d-flex">
          <h3>{{ store.model.name.replaceAll(" ", "_").toLowerCase() }}</h3>
          <v-btn @click="store.changeModelEditorForm('edit')" size="x-small">edit</v-btn>
        </div>

        <v-sheet v-if="!curForm" border class="pa-2">
          <p>
            <i>
              <b>base schema:</b>
              {{ store.model.schemas[store.model.baseSchema].name.replaceAll(" ", "_").toLowerCase() }}</i
            >
          </p>
          <pre>{{ JSON.stringify(formatModelJson(store.model), null, 2).replaceAll('"', "") }}</pre>
        </v-sheet>

        <EditModelForm v-if="curForm === 'edit'" @close-form="store.changeModelEditorForm(null)" />
      </div>
    </v-card-text>
  </v-card>
</template>

<style scoped></style>
