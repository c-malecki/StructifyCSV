<script lang="ts" setup>
import { ref, reactive } from "vue";
import type { VForm } from "vuetify/components";
import {
  schemaPropertyTypes,
  StringProperty,
  NumberProperty,
  IntegerProperty,
  ArrayProperty,
  ObjectProperty,
  BooleanProperty,
  NullProperty,
  type SchemaPropertyType,
  type StringConstructor,
  type NumOrIntConstructor,
  type ArrayConstructor,
} from "../../../types/properties.types";
import {
  type PropertyForm,
  type StringForm,
  type NumOrIntForm,
  type ArrayForm,
} from "../../../util/create";

const emit = defineEmits(["closeForm", "addNewProperty"]);

type NewProperty = {
  key: string;
  dataType: SchemaPropertyType;
};

type FormControl = {
  keyRules: ((val: string) => string | boolean)[];
};
const formControl = reactive<FormControl>({
  keyRules: [(val: string) => val.length > 0 || "Property Name is required."],
});

const formRef = ref<VForm | null>(null);

const newProperty = reactive<NewProperty>({
  key: "",
  dataType: "string",
});

const propertyForm = reactive<PropertyForm>({
  string: {
    minLength: null,
    maxLength: null,
  },
  numOrInt: {
    minimum: null,
    maximum: null,
  },
  array: {
    items: null,
    minItems: null,
    maxItems: null,
  },
});

const createSchemaProperty = (
  type: SchemaPropertyType,
  form: StringConstructor | NumOrIntConstructor | ArrayConstructor | {}
) => {
  switch (type) {
    case "string":
      return new StringProperty(form as StringConstructor);
    case "number":
      return new NumberProperty(form as NumOrIntConstructor);
    case "integer":
      return new IntegerProperty(form as NumOrIntConstructor);
    case "object":
      return new ObjectProperty();
    case "array":
      return new ArrayProperty(form as ArrayConstructor);
    case "boolean":
      return new BooleanProperty();
    case "null":
      return new NullProperty();
  }
};

const getPropertyForm = (
  dataType: SchemaPropertyType
): StringForm | NumOrIntForm | ArrayForm | {} => {
  switch (dataType) {
    case "string":
      const { minLength, maxLength } = propertyForm.string;
      return {
        minLength: minLength !== null ? parseInt(minLength) : undefined,
        maxLength: maxLength !== null ? parseInt(maxLength) : undefined,
      };
    case "integer":
    case "number":
      const { minimum, maximum } = propertyForm.numOrInt;
      return {
        minimum: minimum !== null ? parseInt(minimum) : undefined,
        maximum: maximum !== null ? parseInt(maximum) : undefined,
      };
    case "array":
      const { items, minItems, maxItems } = propertyForm.array;
      return {
        items: items !== null ? items : null,
        minItems: minItems !== null ? parseInt(minItems) : undefined,
        maxItems: maxItems !== null ? parseInt(maxItems) : undefined,
      };
    default:
      return {};
  }
};

const handleSubmit = () => {
  if (!formRef.value) return;

  formRef.value.validate().then(({ valid }) => {
    if (valid) {
      const form = getPropertyForm(newProperty.dataType);
      const value = createSchemaProperty(newProperty.dataType, form);
      emit("addNewProperty", { key: newProperty.key, value });
      emit("closeForm");
    }
  });
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
              v-model="newProperty.key"
              label="Name"
              :rules="formControl.keyRules"
              style="width: 200px"
            />
          </v-col>
          <v-col cols="6">
            <VSelect
              v-model="newProperty.dataType"
              label="Type"
              :items="schemaPropertyTypes"
              style="width: 200px"
            />
          </v-col>
        </v-row>
        <h4
          v-if="
            newProperty.dataType !== 'object' &&
            newProperty.dataType !== 'boolean' &&
            newProperty.dataType !== 'null'
          "
        >
          Attributes
        </h4>
        <v-row v-if="newProperty.dataType === 'string'">
          <v-col cols="6">
            <VTextField
              v-model="propertyForm.string.minLength"
              type="number"
              label="Min Length"
              style="width: 200px"
              hide-details
              clearable
              persistent-clear
            />
          </v-col>
          <v-col cols="6">
            <VTextField
              v-model="propertyForm.string.maxLength"
              type="number"
              label="Max Length"
              style="width: 200px"
              hide-details
              clearable
              persistent-clear
            />
          </v-col>
        </v-row>
        <v-row
          v-if="
            newProperty.dataType === 'number' ||
            newProperty.dataType === 'integer'
          "
        >
          <v-col cols="6">
            <VTextField
              v-model="propertyForm.numOrInt.minimum"
              type="number"
              label="Minimum"
              style="width: 200px"
              hide-details
              clearable
              persistent-clear
            />
          </v-col>
          <v-col cols="6">
            <VTextField
              v-model="propertyForm.numOrInt.maximum"
              type="number"
              label="Maximum"
              style="width: 200px"
              hide-details
              clearable
              persistent-clear
            />
          </v-col>
        </v-row>
        <v-row v-if="newProperty.dataType === 'array'" class="d-flex flex-wrap">
          <v-col cols="12">
            <VSelect
              v-model="propertyForm.array.items"
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
              v-model="propertyForm.array.minItems"
              type="number"
              label="Minimum Items"
              style="width: 200px"
              hide-details
              clearable
              persistent-clear
            />
          </v-col>

          <v-col cols="6">
            <VTextField
              v-model="propertyForm.array.maxItems"
              type="number"
              label="Maximum Items"
              style="width: 200px"
              hide-details
              clearable
              persistent-clear
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
../../../types/properties.types
