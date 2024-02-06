<script lang="ts" setup>
import { reactive, ref, type PropType } from "vue";
import {
  dataTypeOpts,
  type JsonSchemaDataType,
  type PropertiesMapValue,
} from "../../../types/editor.types";
import type { VForm } from "vuetify/components";

type Node = [string, PropertiesMapValue];

const props = defineProps({
  node: {
    type: Object as PropType<Node>,
    required: true,
  },
});

const emit = defineEmits(["closeForm", "updateKey", "updateValue"]);

type EditProperty = {
  key: string;
  dataType: JsonSchemaDataType;
};
type FormControl = {
  keyRules: ((val: string) => string | boolean)[];
};
const formControl: FormControl = {
  keyRules: [(val: string) => val.length > 0 || "Property Name is required."],
};

const getDataType = (nodeValue: PropertiesMapValue) => {
  return typeof nodeValue === "object" ? "object" : nodeValue;
};

const formRef = ref<VForm | null>(null);

const editProperty = reactive<EditProperty>({
  key: props.node[0],
  dataType: getDataType(props.node[1]),
});

const handleSubmit = () => {
  if (!formRef.value) return;
  // needs to account for array and null
  const dataType = getDataType(props.node[1]);
  if (
    editProperty.dataType === dataType &&
    editProperty.key === props.node[0]
  ) {
    emit("closeForm");
    return;
  }

  formRef.value.validate().then(({ valid }) => {
    if (valid) {
      if (
        editProperty.dataType === dataType &&
        editProperty.key !== props.node[0]
      ) {
        emit("updateKey", {
          newKey: editProperty.key,
          oldKey: props.node[0],
          sameValue: props.node[1],
        });
      }
      if (
        editProperty.key === props.node[0] &&
        editProperty.dataType !== dataType
      ) {
        const newValue =
          editProperty.dataType === "object"
            ? new Map<string, PropertiesMapValue>()
            : editProperty.dataType;
        emit("updateValue", { sameKey: props.node[0], newValue });
      }
      emit("closeForm");
    }
  });
};
</script>

<template>
  <VForm @submit.prevent="handleSubmit" ref="formRef">
    <div class="d-flex">
      <VTextField
        v-model="editProperty.key"
        label="Property Name"
        :rules="formControl.keyRules"
        style="width: 200px"
      />
      <VSelect
        v-model="editProperty.dataType"
        :items="dataTypeOpts"
        item-title="name"
        item-value="value"
        style="width: 200px"
      />
      <v-btn type="submit" size="x-small" class="ml-4"> save </v-btn>
      <v-btn size="x-small" class="ml-4" @click="emit('closeForm')">
        cancel
      </v-btn>
    </div>
  </VForm>
</template>
