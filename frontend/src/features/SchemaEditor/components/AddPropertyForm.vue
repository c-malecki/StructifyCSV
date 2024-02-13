<script lang="ts" setup>
import { ref, reactive } from "vue";
import type { VForm } from "vuetify/components";
import {
  schemaPropertyTypes,
  type PropertyConstructorFormValues,
} from "../SchemaEditor.types";
import { propertyFormNullToUndefined } from "../../../util/transform";
import { enforceNumOnKeyDown, enterNumOnKeyUp } from "../../../util/numInput";
import { entity } from "../../../../wailsjs/go/models";

const emit = defineEmits(["closeForm", "addNewProperty"]);

const formValues = reactive<PropertyConstructorFormValues>({
  type: "string",
  minProperties: null,
  maxProperties: null,
  minItems: null,
  maxItems: null,
  items: null,
  minLength: null,
  maxLength: null,
  numMinimum: null,
  numMaximum: null,
  intMinimum: null,
  intMaximum: null,
});
const keyName = ref("");

// todo: min max validation
type FormControl = {
  keyRules: ((val: string) => string | boolean)[];
};
const formControl = reactive<FormControl>({
  keyRules: [(val: string) => val.length > 0 || "Property Name is required."],
});

const formRef = ref<VForm | null>(null);

const handleSubmit = () => {
  if (!formRef.value) return;

  formRef.value.validate().then(({ valid }) => {
    if (valid) {
      const constructorValues = propertyFormNullToUndefined(formValues);
      const value = new entity.Schema(constructorValues);

      emit("addNewProperty", { key: keyName.value, value });
      emit("closeForm");
    }
  });
};

const resetFormOnTypeChange = () => {
  formValues.minProperties = null;
  formValues.maxProperties = null;
  formValues.minItems = null;
  formValues.maxItems = null;
  formValues.items = null;
  formValues.minLength = null;
  formValues.maxLength = null;
  formValues.numMinimum = null;
  formValues.numMaximum = null;
  formValues.intMinimum = null;
  formValues.intMaximum = null;
};
</script>

<template>
  <VForm @submit.prevent="handleSubmit" ref="formRef">
    <v-sheet border rounded class="d-flex flex-column" max-width="440">
      <v-container class="pa-2">
        <h4>Property</h4>
        <v-row>
          <v-col cols="6">
            <VTextField
              v-model="keyName"
              label="Name"
              :rules="formControl.keyRules"
              style="width: 200px"
            />
          </v-col>

          <v-col cols="6">
            <VSelect
              v-model="formValues.type"
              label="Type"
              :items="schemaPropertyTypes"
              style="width: 200px"
              @update:model-value="resetFormOnTypeChange"
            />
          </v-col>
        </v-row>
        <h4 v-if="formValues.type !== 'boolean' && formValues.type !== 'null'">
          Attributes
        </h4>
        <v-row v-if="formValues.type === 'string'">
          <v-col cols="6">
            <VTextField
              :model-value="formValues.minLength"
              type="number"
              label="Min Length"
              style="width: 200px"
              hide-details
              clearable
              persistent-clear
              pattern="[0-9]*"
              inputmode="numeric"
              @keydown="(e: KeyboardEvent) => enforceNumOnKeyDown(e)"
              @keyup="(e: KeyboardEvent) => enterNumOnKeyUp(e, formValues, 'minLength')"
            />
          </v-col>

          <v-col cols="6">
            <VTextField
              :model-value="formValues.maxLength"
              type="number"
              label="Max Length"
              style="width: 200px"
              hide-details
              clearable
              persistent-clear
              pattern="[0-9]*"
              inputmode="numeric"
              @keydown="(e: KeyboardEvent) => enforceNumOnKeyDown(e)"
              @keyup="(e: KeyboardEvent) => enterNumOnKeyUp(e, formValues, 'maxLength')"
            />
          </v-col>
        </v-row>
        <v-row v-if="formValues.type === 'integer'">
          <v-col cols="6">
            <VTextField
              :model-value="formValues.intMinimum"
              type="number"
              label="Minimum"
              style="width: 200px"
              hide-details
              clearable
              persistent-clear
              pattern="[0-9]*"
              inputmode="numeric"
              @keydown="(e: KeyboardEvent) => enforceNumOnKeyDown(e)"
              @keyup="(e: KeyboardEvent) => enterNumOnKeyUp(e, formValues, 'intMinimum')"
            />
          </v-col>

          <v-col cols="6">
            <VTextField
              :model-value="formValues.intMaximum"
              type="number"
              label="Maximum"
              style="width: 200px"
              hide-details
              clearable
              persistent-clear
              pattern="[0-9]*"
              inputmode="numeric"
              @keydown="(e: KeyboardEvent) => enforceNumOnKeyDown(e)"
              @keyup="(e: KeyboardEvent) => enterNumOnKeyUp(e, formValues, 'intMaximum')"
            />
          </v-col>
        </v-row>
        <v-row v-if="formValues.type === 'number'">
          <v-col cols="6">
            <VTextField
              :model-value="formValues.numMinimum"
              type="number"
              label="Minimum"
              style="width: 200px"
              hide-details
              clearable
              persistent-clear
              pattern="[0-9]*"
              inputmode="numeric"
              @keydown="(e: KeyboardEvent) => enforceNumOnKeyDown(e)"
              @keyup="(e: KeyboardEvent) => enterNumOnKeyUp(e, formValues, 'numMinimum')"
            />
          </v-col>

          <v-col cols="6">
            <VTextField
              :model-value="formValues.numMaximum"
              type="number"
              label="Maximum"
              style="width: 200px"
              hide-details
              clearable
              persistent-clear
              pattern="[0-9]*"
              inputmode="numeric"
              @keydown="(e: KeyboardEvent) => enforceNumOnKeyDown(e)"
              @keyup="(e: KeyboardEvent) => enterNumOnKeyUp(e, formValues, 'numMinimum')"
            />
          </v-col>
        </v-row>
        <v-row v-if="formValues.type === 'array'" class="d-flex flex-wrap">
          <v-col cols="12">
            <VSelect
              v-model="formValues.items"
              label="Item Type"
              :items="['string', 'number', 'integer']"
              style="width: 200px"
              hide-details
              clearable
              persistent-clear
            />
          </v-col>

          <v-col cols="6">
            <VTextField
              :model-value="formValues.minItems"
              type="number"
              label="Minimum Items"
              style="width: 200px"
              hide-details
              clearable
              persistent-clear
              pattern="[0-9]*"
              inputmode="numeric"
              @keydown="(e: KeyboardEvent) => enforceNumOnKeyDown(e)"
              @keyup="(e: KeyboardEvent) => enterNumOnKeyUp(e, formValues, 'minItems')"
            />
          </v-col>

          <v-col cols="6">
            <VTextField
              :model-value="formValues.maxItems"
              type="number"
              label="Maximum Items"
              style="width: 200px"
              hide-details
              clearable
              persistent-clear
              pattern="[0-9]*"
              inputmode="numeric"
              @keydown="(e: KeyboardEvent) => enforceNumOnKeyDown(e)"
              @keyup="(e: KeyboardEvent) => enterNumOnKeyUp(e, formValues, 'maxItems')"
            />
          </v-col>
        </v-row>
        <v-row v-if="formValues.type === 'object'">
          <v-col cols="6">
            <VTextField
              :model-value="formValues.minProperties"
              type="number"
              label="Min Properties"
              style="width: 200px"
              hide-details
              clearable
              persistent-clear
              pattern="[0-9]*"
              inputmode="numeric"
              @keydown="(e: KeyboardEvent) => enforceNumOnKeyDown(e)"
              @keyup="(e: KeyboardEvent) => enterNumOnKeyUp(e, formValues, 'minProperties')"
            />
          </v-col>

          <v-col cols="6">
            <VTextField
              :model-value="formValues.maxProperties"
              type="number"
              label="Max Properties"
              style="width: 200px"
              hide-details
              clearable
              persistent-clear
              pattern="[0-9]*"
              inputmode="numeric"
              @keydown="(e: KeyboardEvent) => enforceNumOnKeyDown(e)"
              @keyup="(e: KeyboardEvent) => enterNumOnKeyUp(e, formValues, 'maxProperties')"
            />
          </v-col>
        </v-row>
        <div class="d-flex justify-center mt-4">
          <v-btn size="x-small" @click="emit('closeForm')" class="mr-2">
            cancel
          </v-btn>
          <v-btn type="submit" size="x-small" class="ml-2"> save </v-btn>
        </div>
      </v-container>
    </v-sheet>
  </VForm>
</template>

<style scoped>
.v-input:deep(.v-field__field) {
  height: 36px;
}
.v-input:deep(.v-input__control) {
  height: 36px;
}
.v-input:deep(.v-field__input) {
  min-height: 36px;
  padding-top: 2px;
  padding-bottom: 2px;
}
</style>
