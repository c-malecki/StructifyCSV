<script lang="ts" setup>
import { reactive, computed } from "vue";
import { useStore } from "../../../store/store";
import { type VForm } from "vuetify/lib/components/index.mjs";

const emit = defineEmits(["closeForm"]);
const store = useStore();
const formValues = reactive({
  selectedHeaders: [],
});

const headerOpts = computed(() => (store.csv ? store.csv.headers : []));

// const handleSubmit = () => {
//   store.updateSelectedColumns(formValues.selectedHeaders);
//   store.changeCsvEditorForm("none");
// };
</script>

<template>
  <!-- <v-form @submit.prevent="handleSubmit">
    <p><b>Total Headers:</b> {{ headerOpts.length }}</p>
    <VAutocomplete
      v-model="formValues.selectedHeaders"
      label="Select CSV Headers"
      :items="headerOpts"
      multiple
      counter
      persistent-counter
    />
    <div class="d-flex">
      <v-btn type="button" size="small" @click="store.changeCsvEditorForm('none')" class="ml-auto mr-8"> cancel </v-btn>
      <v-btn type="submit" size="small">save</v-btn>
    </div>
  </v-form> -->
</template>

<style scoped></style>
