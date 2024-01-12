<script lang="ts" setup>
import { ref, reactive } from "vue";
import { useStore, fieldOptions } from "../../../store/store";
import { main } from "../../../../wailsjs/go/models";
import { formatModelJson } from "../../util/format";
import { type VForm, type VSelect } from "vuetify/lib/components/index.mjs";

type FormValues = main.Model;
type FormControl = {
  nameRules: ((val: string) => string | boolean)[];
};

const emit = defineEmits(["closeForm"]);
const store = useStore();
const formRef = ref<VForm | null>(null);
const hiddenRef = ref<VSelect | null>(null);
const hiddenError = ref(false);
const formValues = reactive<FormValues>(
  new main.Model({
    name: "Example Model",
    type: "model",
    baseSchema: 0,
    schemas: [
      {
        name: "Example Schema",
        type: "schema",
        fields: [
          {
            name: "First Name",
            type: "field",
            dataType: { name: "text (string)", value: "string" },
          },
          {
            name: "Last Name",
            type: "field",
            dataType: { name: "text (string)", value: "string" },
          },
          {
            name: "Age",
            type: "field",
            dataType: { name: "number", value: "number" },
          },
        ],
      },
    ],
  })
);
const formControl: FormControl = {
  nameRules: [(v: string) => v.length > 0 || "Name is required."],
};

const addFieldToSchema = () => {
  formValues.schemas[0].fields.push(
    new main.Field({
      name: "",
      type: "field",
      dataType: { name: "text (string)", value: "string" },
    })
  );
  hiddenError.value = false;
};

const removeFieldFromSchema = (idx: number) => {
  formValues.schemas[0].fields = formValues.schemas[0].fields.filter((f, i) => i !== idx);
  if (!formValues.schemas[0].fields.length) {
    hiddenError.value = true;
  }
};

const handleSubmit = () => {
  if (!formRef.value) return;
  formRef.value.validate().then(({ valid }) => {
    if (valid) {
      store.updateLocalModel(formValues);
      emit("closeForm");
    }
  });
};
</script>

<template>
  <v-form @submit.prevent="handleSubmit" ref="formRef" style="max-width: 1200px">
    <v-container fluid>
      <v-row>
        <v-col>
          <VTextField v-model="formValues.name" label="Model Name" :rules="formControl.nameRules" />
          <p>
            <i><b>base schema:</b> {{ formValues.schemas[0].name.replaceAll(" ", "_").toLowerCase() }}</i>
          </p>
          <v-sheet border class="pa-2">
            <pre>{{ JSON.stringify(formatModelJson(formValues), null, 2).replaceAll('"', "") }}</pre>
          </v-sheet>
        </v-col>
        <v-col>
          <VTextField v-model="formValues.schemas[0].name" label="Base Schema Name" :rules="formControl.nameRules" />
          <v-btn @click="addFieldToSchema">Add Field</v-btn>
          <VSelect
            :error-messages="hiddenError ? 'At least one Field is required' : undefined"
            class="hidden-input"
            ref="hiddenRef"
          />
          <v-sheet v-for="(field, fIdx) in formValues.schemas[0].fields" :key="`field-${fIdx}`" border class="pa-2">
            <v-row>
              <v-col>
                <VTextField
                  v-model="formValues.schemas[0].fields[fIdx].name"
                  label="Field Name"
                  :rules="formControl.nameRules"
                />
              </v-col>
              <v-col>
                <VSelect
                  v-model="formValues.schemas[0].fields[fIdx].dataType"
                  label="Data Type"
                  :items="fieldOptions"
                  item-title="name"
                  item-value="value"
                  return-object
                />
              </v-col>
              <v-col>
                <v-btn icon="mdi-close" @click="removeFieldFromSchema(fIdx)"></v-btn>
              </v-col>
            </v-row>
          </v-sheet>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <div class="d-flex">
            <v-btn type="button" @click="emit('closeForm')"> cancel </v-btn>
            <v-btn type="submit"> create </v-btn>
          </div>
        </v-col>
      </v-row>
    </v-container>
  </v-form>
</template>

<style scoped>
.hidden-input:deep(.v-input__control) {
  display: none;
}
</style>
