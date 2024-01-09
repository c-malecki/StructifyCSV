<script lang="ts" setup>
import { reactive } from "vue";
import { useStore } from "../../../store/store";
const store = useStore();

type FormData = {
  modelName: string;
};
type FormControl = {
  modelNameError: string | null;
};

const formData = reactive<FormData>({
  modelName: "",
});
const formControl = reactive<FormControl>({
  modelNameError: null,
});

const reset = () => {
  formData.modelName = "";
  formControl.modelNameError = null;
};

const validate = () => {
  return formData.modelName.length > 0 || "Model Name is required.";
};

function handleSubmit() {
  const valid = validate();
  if (valid === true) {
    store.createModel(formData.modelName);
    reset();
  } else {
    formControl.modelNameError = valid;
  }
}
</script>

<template>
  <div>
    <form @submit.prevent="handleSubmit">
      <div class="col">
        <label for="Model Name">Model Name</label>
        <input v-model="formData.modelName" name="Model Name" />
        <p class="error">{{ formControl.modelNameError }}</p>
      </div>
      <button type="submit">Create</button>
    </form>
  </div>
</template>

<style scoped></style>
