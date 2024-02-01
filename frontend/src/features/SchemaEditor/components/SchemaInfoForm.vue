<script lang="ts" setup>
import { reactive, ref, inject } from "vue";
import type { VForm } from "vuetify/lib/components/index.mjs";
import type { SchemaValues } from "../schemaEditor.types";
import { SchemaValuesKey } from "../schemaEditor.util";

type FormControl = {
  titleRules: ((val: string) => string | boolean)[];
  descriptionRules: ((val: string) => string | boolean)[];
};

const emit = defineEmits(["closeForm", "updateSchema"]);

const schemaValues = inject(SchemaValuesKey);
if (!schemaValues) {
  throw new Error(`Could not resolve ${SchemaValuesKey.description}`);
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
const formValues = reactive<Omit<SchemaValues, "properties">>({
  title: schemaValues.title,
  description: schemaValues.description,
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
