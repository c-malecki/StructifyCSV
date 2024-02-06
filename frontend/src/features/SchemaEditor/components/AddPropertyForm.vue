<script lang="ts" setup>
import { ref, reactive, type PropType } from "vue";
import type { VForm } from "vuetify/components";
import { type PropertiesMapValue } from "../../../types/editor.types";
import {
  dataTypeOpts,
  type JsonSchemaDataType,
} from "../../../types/schema.types";

const props = defineProps({
  nodeValue: {
    type: Object as PropType<Map<string, PropertiesMapValue>>,
    required: true,
  },
});

const emit = defineEmits(["hideForm"]);

type FormControl = {
  showAdd: boolean;
  keyRules: ((val: string) => string | boolean)[];
};

type NewProperty = {
  key: string;
  value: JsonSchemaDataType;
};

const formRef = ref<VForm | null>(null);
const formControl = reactive<FormControl>({
  showAdd: false,
  keyRules: [(val: string) => val.length > 0 || "Property Name is required."],
});

const newProperty = reactive<NewProperty>({
  key: "",
  value: "string",
});

const resetAddProperty = () => {
  newProperty.key = "";
  newProperty.value = "string";
  emit("hideForm");
};

const addProperty = () => {
  formRef.value!.validate().then(({ valid }) => {
    if (valid) {
      props.nodeValue.set(
        newProperty.key,
        newProperty.value === "object"
          ? new Map<string, PropertiesMapValue>()
          : newProperty.value
      );
      resetAddProperty();
    }
  });
};
</script>

<template>
  <VForm @submit.prevent="addProperty" ref="formRef">
    <div class="d-flex">
      <VTextField
        v-model="newProperty.key"
        label="Property Name"
        :rules="formControl.keyRules"
        style="max-width: 200px"
      />
      <VSelect
        v-model="newProperty.value"
        :items="dataTypeOpts"
        style="max-width: 200px"
      />
      <v-btn
        type="button"
        size="x-small"
        class="ml-4"
        @click="resetAddProperty"
      >
        cancel
      </v-btn>
      <v-btn type="submit" size="x-small" class="ml-4">save</v-btn>
    </div>
  </VForm>
</template>
