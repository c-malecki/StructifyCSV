<script lang="ts" setup>
import { ref, provide } from "vue";
import { entity } from "../../../wailsjs/go/models";
import {
  ImportJsonSchema,
  ExportJsonSchema,
} from "../../../wailsjs/go/main/App";
import { JsonSchemaKey } from "./SchemaEditor.types";
import SchemaInfoDisplayVue from "./components/SchemaInfoDisplay.vue";
import PropertyTree from "./components/PropertyTree.vue";
import SchemaInfoForm from "./components/forms/SchemaInfoForm.vue";
import { fixWailsJsonSchemaImport } from "../../util/transform";
import { exampleSchema } from "../../util/example";

const props = defineProps({
  hidden: {
    type: Boolean,
    required: true,
  },
});

const jsonSchema = ref<entity.JsonSchema>(exampleSchema);
provide(JsonSchemaKey, jsonSchema);

const emit = defineEmits(["changeTab"]);

const showEditForm = ref(false);

const newSchema = () => {
  const schema = new entity.JsonSchema({
    title: "New Schema",
    description:
      "To change the name and description of this Schema, use the EDIT button to the right. \nTo begin building your Schema, click the ADD button below.",
    properties: {},
  });

  jsonSchema.value = schema;
  emit("changeTab", "schema");
};

const importSchema = () => {
  ImportJsonSchema()
    .then(({ schema, error }) => {
      if (error) {
        console.log(error);
        // show import error somewhere
      }
      if (schema) {
        // when Wails unmarshals the JSON file, values are null instead of undefined
        // but the generated models.ts class uses undefined
        const properties = fixWailsJsonSchemaImport(schema.properties);
        jsonSchema.value = { ...schema, properties };
      }
      emit("changeTab", "schema");
    })
    .catch(() => {});
};

const exportSchema = () => {
  ExportJsonSchema(jsonSchema.value)
    .then(() => {
      emit("changeTab", "schema");
    })
    .catch((err) => {});
};

defineExpose({ newSchema, importSchema, exportSchema });

const handleUpdateSchema = ({
  title,
  description,
}: Omit<entity.JsonSchema, "properties">) => {
  jsonSchema.value.title = title;
  jsonSchema.value.description = description;
};
</script>

<template>
  <v-sheet rounded="0" flat :class="{ hide: props.hidden }">
    <SchemaInfoDisplayVue
      v-if="!showEditForm"
      :schema="jsonSchema"
      @show-edit="showEditForm = true"
    />

    <SchemaInfoForm
      v-if="showEditForm"
      @close-form="showEditForm = false"
      @update-schema="handleUpdateSchema"
    />

    <v-divider />

    <PropertyTree />
  </v-sheet>
</template>

<style scoped>
.hide {
  display: none;
}
</style>
