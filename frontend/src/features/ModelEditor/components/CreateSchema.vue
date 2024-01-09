<script lang="ts" setup>
import { reactive } from "vue";
import { useStore } from "../../../store/store";
const store = useStore();

type FormData = {
  schemaName: string;
};
type FormControl = {
  schemaNameError: string | null;
};

const formData = reactive<FormData>({
  schemaName: "",
});
const formControl = reactive<FormControl>({
  schemaNameError: null,
});

const reset = () => {
  formData.schemaName = "";
  formControl.schemaNameError = null;
};

const validate = () => {
  return formData.schemaName.length > 0 || "Model Name is required.";
};

function handleSubmit() {
  const valid = validate();
  if (valid === true) {
    store.createSchema(formData.schemaName);
    reset();
  } else {
    formControl.schemaNameError = valid;
  }
}
</script>

<template>
  <div>
    <h3>New Schema</h3>
    <form @submit.prevent="handleSubmit">
      <div class="col">
        <label for="Schema Name">Schema Name</label>
        <input v-model="formData.schemaName" name="Schema Name" />
        <p class="error">{{ formControl.schemaNameError }}</p>
      </div>

      <button type="submit">Create</button>
    </form>
  </div>
</template>

<style scoped></style>
