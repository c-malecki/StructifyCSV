<script lang="ts" setup>
import { reactive, ref, inject } from "vue";
import type { VForm } from "vuetify/components";
import { JsonSchemaKey, type JsonSchema } from "../../../types/editor.types";

type FormControl = {
  titleRules: ((val: string) => string | boolean)[];
  descriptionRules: ((val: string) => string | boolean)[];
};

const emit = defineEmits(["closeForm", "updateSchema"]);

const jsonSchema = inject(JsonSchemaKey);
if (!jsonSchema) {
  throw new Error(`Could not resolve ${JsonSchemaKey.description}`);
}

const formRef = ref<VForm | null>(null);
const formControl: FormControl = {
  titleRules: [
    (v: string) => v.length > 0 || "Schema Name is required.",
    (v: string) =>
      [...v].length <= 150 ||
      "Schema Name cannot be longer than 150 characters.",
  ],
  descriptionRules: [
    (v: string) =>
      [...v].length <= 1000 ||
      "Description cannot be longer than 1000 characters.",
  ],
};
const formValues = reactive<Omit<JsonSchema, "properties">>({
  title: jsonSchema.title,
  description: jsonSchema.description,
});

const handleSubmit = () => {
  if (!formRef.value) return;
  formRef.value.validate().then(({ valid }) => {
    if (valid) {
      emit("updateSchema", formValues);
    }
  });
};
</script>

<template>
  <VForm @submit.prevent="handleSubmit" ref="formRef">
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
      <v-btn
        type="button"
        size="small"
        @click="emit('closeForm')"
        class="ml-auto mr-4"
      >
        cancel
      </v-btn>
      <v-btn type="submit" size="small">save</v-btn>
    </div>
  </VForm>
</template>

<style scoped>
.v-input:not(.v-textarea):deep(.v-field__field) {
  height: 36px;
}
.v-input:not(.v-textarea):deep(.v-input__control) {
  height: 36px;
}
.v-input:not(.v-textarea):deep(.v-field__input) {
  min-height: 36px;
  padding-top: 2px;
  padding-bottom: 2px;
}
</style>
