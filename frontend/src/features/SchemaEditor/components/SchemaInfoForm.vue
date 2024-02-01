<script lang="ts" setup>
import { reactive, ref } from "vue";
import type { PropType } from "vue";
import type { VForm } from "vuetify/lib/components/index.mjs";
import type { SchemaInfo } from "../schemaEditor.types";

type FormControl = {
  titleRules: ((val: string) => string | boolean)[];
  descriptionRules: ((val: string) => string | boolean)[];
};

const props = defineProps({
  schemaInfo: {
    type: Object as PropType<SchemaInfo>,
    required: true,
  },
});
const emit = defineEmits(["closeForm", "updateSchema"]);

const formRef = ref<VForm | null>(null);
const formValues = reactive<SchemaInfo>({
  title: props.schemaInfo.title,
  description: props.schemaInfo.description,
});
const formControl: FormControl = {
  titleRules: [
    (v: string) => v.length > 0 || "Schema Name is required.",
    (v: string) => [...v].length <= 150 || "Schema Name cannot be longer than 150 characters.",
  ],
  descriptionRules: [(v: string) => [...v].length <= 1000 || "Description cannot be longer than 1000 characters."],
};

const handleSubmit = () => {
  formRef.value!.validate().then(({ valid }) => {
    if (valid) {
      emit("updateSchema", formValues);
    }
  });
};

const closeForm = () => {
  emit("closeForm");
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
      <v-btn type="button" size="small" @click="closeForm" class="ml-auto mr-4">cancel</v-btn>
      <v-btn type="submit" size="small">save</v-btn>
    </div>
  </VForm>
</template>
