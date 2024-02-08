<script lang="ts" setup>
import { inject, computed, type PropType } from "vue";
import { getHoverColorScheme, getLeftIndent } from "../../../util/style";
import {
  HeaderOptsKey,
  type CsvSchemaMap,
  type CsvSchemaMapValue,
  type CsvSchemaProperty,
} from "../../../types/editor.types";

const headerOpts = inject(HeaderOptsKey);
if (!headerOpts) {
  throw new Error(`Could not resolve ${HeaderOptsKey.description}`);
}

// make note about directly mutating props and why
const nodeProps = defineProps({
  nodeKey: {
    type: String,
    required: true,
  },
  nodeValue: {
    type: Object as PropType<CsvSchemaMapValue>,
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

// const handleUpdateCsvHeader = (val: string | string[] | null) => {
//   const csvSchemaNodeValue = nodeProps.nodeValue as CsvSchemaNodeValue;
//   if (val !== null) {
//     if (Array.isArray(val)) {

//     }
//     const index = csvSchema.headerDescriptors.findIndex(
//       (h) => h.headerText === val
//     );
//     csvSchema.usedHeaderIndexes.push(index);
//     csvSchemaNodeValue.headerIndex = index;
//     csvSchema.headerDescriptors[index].propertyDescriptor = {
//       key: nodeProps.nodeKey,
//       path: csvSchemaNodeValue.schemaPath,
//       type: csvSchemaNodeValue.schemaPropertyType,
//     };
//   } else {
//     csvSchema.usedHeaderIndexes = csvSchema.usedHeaderIndexes.filter(
//       (h, i) => i !== csvSchemaNodeValue.headerIndex!
//     );
//     csvSchema.headerDescriptors[
//       csvSchemaNodeValue.headerIndex!
//     ].propertyDescriptor = undefined;
//     csvSchemaNodeValue.headerIndex = null;
//   }
// };
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
              v-model="(nodeProps.nodeValue as CsvSchemaProperty).headerIndexes"
              :items="headerOpts"
              label="Headers"
              item-title="header"
              item-value="index"
              style="max-width: 300px"
              hide-details
              clearable
              persistent-clear
              :multiple="(nodeProps.nodeValue as CsvSchemaProperty).schemaPropertyType === 'array'"
            />
          </div>

          <div v-else>
            <p class="opening_bracket">
              <b>{{ nodeKey }}: </b>
            </p>
            <ModelNode
              v-for="(node, i) in (nodeProps.nodeValue as CsvSchemaMap)"
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
  min-height: 36px;
  padding-top: 2px;
  padding-bottom: 2px;
}
</style>
