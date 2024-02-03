<script lang="ts" setup>
import { ref, inject, computed, type PropType } from "vue";
import { getHoverColorScheme, getLeftIndent } from "../../../util/style";
import {
  CsvModelKey,
  type CsvModelMap,
  type CsvModelProperty,
} from "../../../types/editor.types";

// make note about directly mutating props and why
const nodeProps = defineProps({
  nodeKey: {
    type: String,
    required: true,
  },
  nodeValue: {
    type: Object as PropType<CsvModelProperty | CsvModelMap>,
    required: true,
  },
  level: {
    type: Number,
    default: 1,
  },
});

const isMap = computed(() => nodeProps.nodeValue instanceof Map);
const leftIndent = computed(() => getLeftIndent(nodeProps.level));
const colorScheme = computed(() => getHoverColorScheme(nodeProps.level));

const csvModel = inject(CsvModelKey);
if (!csvModel) {
  throw new Error(`Could not resolve ${CsvModelKey.description}`);
}

const handleUpdateCsvHeader = (val: string | null) => {
  const csvModelProperty = nodeProps.nodeValue as CsvModelProperty;
  if (val !== null) {
    const index = csvModel.headers.findIndex((h) => h.header === val);
    csvModelProperty.headerIdx = index;
    csvModel.headers[index].schemaProperty = {
      key: nodeProps.nodeKey,
      schemaPath: csvModelProperty.schemaPath,
      value: csvModelProperty.dataType,
    };
  } else {
    csvModel.headers[csvModelProperty.headerIdx!].schemaProperty = null;
    csvModelProperty.headerIdx = null;
  }
};
</script>

<template>
  <v-hover>
    <template v-slot:default="{ isHovering, props }">
      <v-card
        v-bind="props"
        :color="isHovering ? colorScheme.hover : undefined"
        variant="flat"
        :style="leftIndent"
      >
        <div
          :class="{ closing_bracket: isMap }"
          :style="`color: ${isHovering ? colorScheme.font : 'black'}`"
          class="pl-1 pt-1 pb-1"
        >
          <div v-if="!isMap" class="d-flex align-center">
            <p class="mr-2">
              <b>{{ nodeKey }}: </b>
            </p>

            <VAutocomplete
              v-model="(nodeProps.nodeValue as CsvModelProperty).header"
              :items="csvModel.headers"
              label="Headers"
              item-title="header"
              item-value="header"
              style="max-width: 200px"
              hide-details
              clearable
              persistent-clear
              @update:model-value="handleUpdateCsvHeader"
            />
          </div>

          <div v-else>
            <p class="opening_bracket">
              <b>{{ nodeKey }}: </b>
            </p>
            <ModelNode
              v-for="(node, i) in (nodeProps.nodeValue as CsvModelMap)"
              :key="`${nodeProps.level + 1}-${i}-${typeof node[1]}-csv`"
              :nodeKey="node[0]"
              :nodeValue="node[1]"
              :level="nodeProps.level + 1"
            />
          </div>
        </div>
      </v-card>
    </template>
  </v-hover>
</template>

<style scoped>
.closing_bracket:after {
  content: "}";
  margin-left: 16px;
}
.opening_bracket:after {
  content: "{";
  margin-left: 16px;
}
.v-input:deep(.v-field__field) {
  height: 36px;
}
.v-input:deep(.v-input__control) {
  height: 36px;
}
.v-input:deep(.v-field__input) {
  padding-top: 2px;
  padding-bottom: 2px;
}
</style>
