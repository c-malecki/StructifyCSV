<script lang="ts" setup>
import { reactive, ref, type PropType } from "vue";
import { useSchemaStore, type PropertyType } from "../../../../store/schema";
import type { VForm } from "vuetify/components";
import {
  type PropertyNode,
  type ArrayItemType,
  type PropertyConstructorFormValues,
} from "../../SchemaEditor.types";
import { propertyFormNullToUndefined } from "../../../../util/transform";
import {
  enforceNumOnKeyDown,
  enterNumOnKeyUp,
} from "../../../../util/numInput";
import { entity } from "../../../../../wailsjs/go/models";

const schemaStore = useSchemaStore();

const props = defineProps({
  node: {
    type: Object as PropType<PropertyNode>,
    required: true,
  },
});

const emit = defineEmits(["closeForm", "updateKey", "updateValue"]);

const formValues = reactive<PropertyConstructorFormValues>({
  type: props.node[1].type as PropertyType,
  minProperties: props.node[1].minProperties?.toString() ?? null,
  maxProperties: props.node[1].maxProperties?.toString() ?? null,
  minItems: props.node[1].minItems?.toString() ?? null,
  maxItems: props.node[1].maxItems?.toString() ?? null,
  items: props.node[1].items
    ? (props.node[1].items.type as ArrayItemType)
    : null,
  minLength: props.node[1].minLength?.toString() ?? null,
  maxLength: props.node[1].maxLength?.toString() ?? null,
  numMinimum: props.node[1].numMinimum?.toString() ?? null,
  numMaximum: props.node[1].numMaximum?.toString() ?? null,
  intMinimum: props.node[1].intMinimum?.toString() ?? null,
  intMaximum: props.node[1].intMaximum?.toString() ?? null,
});
const keyName = ref(props.node[0]);

// todo: min max validation
type FormControl = {
  keyRules: ((val: string) => string | boolean)[];
};
const formControl: FormControl = {
  keyRules: [(val: string) => val.length > 0 || "Property Name is required."],
};

const formRef = ref<VForm | null>(null);

const handleSubmit = () => {
  if (!formRef.value) return;

  formRef.value.validate().then(({ valid }) => {
    if (valid) {
      if (
        formValues.type === "object" &&
        props.node[1].type === "object" &&
        keyName.value !== props.node[0]
      ) {
        emit("updateKey", {
          editKey: keyName.value,
          curKey: props.node[0],
          value: props.node[1],
        });
      } else {
        const constructorValues = propertyFormNullToUndefined(formValues);
        const properties =
          props.node[1].type === "object" && formValues.type === "object"
            ? props.node[1].properties
            : {};

        const value = new entity.PropertySchema({
          ...constructorValues,
          properties,
        });

        emit("updateValue", {
          editKey: keyName.value,
          curKey: props.node[0],
          value,
        });
      }
      emit("closeForm");
    }
  });
};

const resetFormOnTypeChange = (typeVal: PropertyType) => {
  if (typeVal === props.node[1].type) {
    formValues.minProperties = props.node[1].minProperties?.toString() ?? null;
    formValues.maxProperties = props.node[1].maxProperties?.toString() ?? null;
    formValues.minItems = props.node[1].minItems?.toString() ?? null;
    formValues.maxItems = props.node[1].maxItems?.toString() ?? null;
    formValues.items = props.node[1].items
      ? (props.node[1].items.type as ArrayItemType)
      : null;
    formValues.minLength = props.node[1].minLength?.toString() ?? null;
    formValues.maxLength = props.node[1].maxLength?.toString() ?? null;
    formValues.numMinimum = props.node[1].numMinimum?.toString() ?? null;
    formValues.numMaximum = props.node[1].numMaximum?.toString() ?? null;
    formValues.intMinimum = props.node[1].intMinimum?.toString() ?? null;
    formValues.intMaximum = props.node[1].intMaximum?.toString() ?? null;
  } else {
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
  }
};
</script>

<template>
  <VForm @submit.prevent="handleSubmit" ref="formRef">
    <v-sheet border rounded class="d-flex flex-column" max-width="440">
      <v-container class="pa-2">
        <h5 class="mb-2">"{{ props.node[0] }}" Property</h5>
        <v-row>
          <v-col cols="6">
            <VTextField
              v-model="keyName"
              label="Key"
              :rules="formControl.keyRules"
              style="width: 200px"
            />
          </v-col>

          <v-col cols="6">
            <VSelect
              v-model="formValues.type"
              label="Type"
              :items="schemaStore.propertyTypes"
              style="width: 200px"
              @update:model-value="resetFormOnTypeChange"
            />
          </v-col>
        </v-row>
        <h5
          v-if="formValues.type !== 'boolean' && formValues.type !== 'null'"
          class="mb-2"
        >
          {{ formValues.type }} validators
        </h5>
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

h5:not(:first-of-type) {
  text-transform: capitalize;
}
</style>
