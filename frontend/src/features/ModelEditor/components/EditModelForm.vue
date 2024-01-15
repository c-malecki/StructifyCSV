<script lang="ts" setup>
import { ref, computed, reactive } from "vue";
import { useStore } from "../../../store/store";
import { entity } from "../../../../wailsjs/go/models";
import { formatModelJson } from "../../util/format";
import { type VForm } from "vuetify/lib/components/index.mjs";

type FormValues = {
  model: entity.Model;
};
type FormControl = {
  nameRules: ((val: string) => string | boolean)[];
};

const emit = defineEmits(["closeForm"]);
const store = useStore();
const formRef = ref<VForm | null>(null);
const formValues = reactive<FormValues>({
  model: store.model!,
});
const formControl: FormControl = {
  nameRules: [(v: string) => v.length > 0 || "Model Name is required."],
};

const modelBaseOptions = computed(() =>
  store.model!.schemas.map((s, i) => {
    return {
      name: s.name,
      value: i,
    };
  })
);

const handleSubmit = () => {
  if (!formRef.value) return;
  formRef.value.validate().then(({ valid }) => {
    if (valid) {
      store.updateLocalModel(formValues.model);
      emit("closeForm");
    }
  });
};
</script>

<template>
  <v-form @submit.prevent="handleSubmit" ref="formRef">
    <v-container fluid>
      <v-row>
        <v-col>
          <VTextField v-model="formValues.model.name" label="Model Name" :rules="formControl.nameRules" />
          <p>
            <i
              ><b>base schema:</b>
              {{ formValues.model.schemas[formValues.model.baseSchemaIdx].name.replaceAll(" ", "_").toLowerCase() }}</i
            >
          </p>
          <v-sheet border class="pa-2">
            <pre>{{ JSON.stringify(formatModelJson(formValues.model), null, 2).replaceAll('"', "") }}</pre>
          </v-sheet>
        </v-col>
        <v-col>
          <VSelect
            v-model="formValues.model.baseSchemaIdx"
            label="Model Base Schema"
            :items="modelBaseOptions"
            item-title="name"
            item-value="value"
          />
        </v-col>
      </v-row>
    </v-container>

    <div class="d-flex">
      <v-btn type="button" @click="emit('closeForm')"> cancel </v-btn>
      <v-btn type="submit">save</v-btn>
    </div>
  </v-form>
</template>
