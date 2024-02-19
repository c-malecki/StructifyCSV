<script lang="ts" setup>
import {
  ImportCsvFileData,
  ProcessCsvWithSchema,
} from "../../../wailsjs/go/main/App";
import { reactive, ref, provide } from "vue";
import { CsvFileKey } from "./CsvEditor.types";
import { JsonSchemaKey } from "../SchemaEditor/SchemaEditor.types";
import ModelTree from "./components/ModelTree.vue";
import ProcessingReport from "../../ui/ProcessingReport.vue";
import { exampleSchema, exampleCsvFile } from "../../util/example";
import { entity } from "../../../wailsjs/go/models";

const props = defineProps({
  hidden: {
    type: Boolean,
    required: true,
  },
});

const csvFile = reactive<entity.CsvFileData>(exampleCsvFile);
provide(CsvFileKey, csvFile);

const selectedSchema = ref<entity.JsonSchema>(exampleSchema);
provide(JsonSchemaKey, selectedSchema);

const report = ref<entity.CsvProcessingReport | null>(null);

const importCsv = async () => {
  try {
    const csvFileData = await ImportCsvFileData();
    if (csvFileData.headers !== null) {
      csvFile.fileName = csvFileData.fileName;
      csvFile.location = csvFileData.location;
      csvFile.headers = csvFileData.headers;
    } else {
      console.log("canceled import");
    }
  } catch (err) {
    console.log(err);
  }
};

const processCsv = () => {
  ProcessCsvWithSchema(selectedSchema.value).then((res) => {
    report.value = res;
  });
};

defineExpose({ importCsv, processCsv, csvFile, selectedSchema });
</script>

<template>
  <v-sheet border rounded="0" flat :class="{ hide: props.hidden }">
    <h3 v-if="csvFile">{{ csvFile.fileName }}</h3>

    <v-divider />
    <ProcessingReport v-if="report" :report="report" />
    <ModelTree />
  </v-sheet>
</template>

<style scoped>
.hide {
  display: none;
}
</style>
