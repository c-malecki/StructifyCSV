<script lang="ts" setup>
import { ref } from "vue";
import { useAppStore } from "../../store/app";
import SchemaInfoDisplayVue from "./components/SchemaInfoDisplay.vue";
import PropertyTree from "./components/PropertyTree.vue";
import SchemaInfoForm from "./components/forms/SchemaInfoForm.vue";

const appStore = useAppStore();

const showEditForm = ref(false);
</script>

<template>
  <v-sheet rounded="0" flat :class="{ hide: appStore.appTab === 'csv' }">
    <SchemaInfoDisplayVue
      v-if="!showEditForm"
      @show-edit="showEditForm = true"
    />

    <SchemaInfoForm v-if="showEditForm" @close-form="showEditForm = false" />

    <PropertyTree />
  </v-sheet>
</template>

<style scoped>
.hide {
  display: none;
}
</style>
