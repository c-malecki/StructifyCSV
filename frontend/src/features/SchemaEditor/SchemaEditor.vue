<script lang="ts" setup>
import { reactive, ref } from "vue";
import { useAppStore } from "../../store/app";
import { useSchemaStore } from "../../store/schema";
import type { VForm } from "vuetify/components";
import { entity } from "../../../wailsjs/go/models";
import PropertyTree from "./PropertyTree/PropertyTree.vue";

const appStore = useAppStore();
const schemaStore = useSchemaStore();

type FormControl = {
  titleRules: ((val: string) => string | boolean)[];
  descriptionRules: ((val: string) => string | boolean)[];
};

const showEditForm = ref(false);

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
const formValues = reactive<
  Omit<entity.JsonSchema, "properties" | "required" | "type">
>({
  title: schemaStore.jsonSchema.title,
  description: schemaStore.jsonSchema.description,
});

const handleSubmit = () => {
  if (!formRef.value) return;
  formRef.value.validate().then(({ valid }) => {
    if (valid) {
      schemaStore.updateJsonSchemaNameDesc(formValues);
      showEditForm.value = false;
    }
  });
};

const cancelEdit = () => {
  showEditForm.value = false;
  (formValues.title = schemaStore.jsonSchema.title),
    (formValues.description = schemaStore.jsonSchema.description);
};
</script>

<template>
  <div :class="{ hide: appStore.appTab === 'csv' }">
    <v-sheet class="pa-4 ma-2">
      <div v-if="!showEditForm">
        <div class="d-flex align-center">
          <h3>
            {{ schemaStore.jsonSchema.title }}
          </h3>
          <v-btn
            color="primary"
            size="x-small"
            class="ml-2"
            @click="showEditForm = true"
          >
            edit
          </v-btn>
        </div>

        <p>{{ schemaStore.jsonSchema.description }}</p>
      </div>

      <VForm v-else @submit.prevent="handleSubmit" ref="formRef">
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
          rows="4"
          persistent-counter
        />
        <div class="d-flex mt-2">
          <v-btn
            type="button"
            size="small"
            @click="cancelEdit"
            class="ml-auto mr-4"
          >
            cancel
          </v-btn>
          <v-btn type="submit" size="small">save</v-btn>
        </div>
      </VForm>
    </v-sheet>

    <PropertyTree />
  </div>
</template>

<style scoped>
.hide {
  display: none;
}
p {
  white-space: pre-line;
}
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
