<script lang="ts" setup>
import { ref, computed, type PropType } from "vue";
import { type VForm } from "vuetify/components";
import { entity } from "../../../../../wailsjs/go/models";

const props = defineProps({
  schema: {
    type: Object as PropType<entity.Schema | entity.JsonSchema>,
    required: true,
  },
});

const emit = defineEmits(["updateRequired", "closeForm"]);
const required = ref<string[]>(props.schema.required);

const childPropertyOpts = computed(() => {
  const childPropertyKeys = Object.keys(props.schema.properties);
  return childPropertyKeys.map((k) => {
    return {
      propertyKey: k,
      type: props.schema.properties[k].type,
    };
  });
});

const handleSubmit = () => {
  emit("updateRequired", required.value);
};
</script>

<template>
  <VForm @submit.prevent="handleSubmit" ref="formRef" style="margin-left: 24px">
    <v-sheet border rounded class="d-flex flex-column" max-width="440">
      <v-data-table
        v-model="required"
        :items="childPropertyOpts"
        item-value="propertyKey"
        show-select
      />
      <v-divider />
      <div class="d-flex justify-center pa-2">
        <v-btn size="x-small" @click="emit('closeForm')" class="mr-2">
          cancel
        </v-btn>
        <v-btn type="submit" size="x-small" class="ml-2"> save </v-btn>
      </div>
    </v-sheet>
  </VForm>
</template>
