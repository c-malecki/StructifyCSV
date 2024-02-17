<script lang="ts" setup>
import { inject, computed, type PropType } from "vue";
import { getHoverColorScheme } from "../../../util/style";
import { HeaderOptsKey } from "../CsvEditor.types";
import { type PropertyNode } from "../../SchemaEditor/SchemaEditor.types";

const headerOpts = inject(HeaderOptsKey);
if (!headerOpts) {
  throw new Error(`Could not resolve ${HeaderOptsKey.description}`);
}

// make note about directly mutating props and why
const props = defineProps({
  node: {
    type: Object as PropType<PropertyNode>,
    required: true,
  },
  level: {
    type: Number,
    default: 1,
  },
});

const colorScheme = computed(() => getHoverColorScheme(props.level));
</script>

<template>
  <v-hover>
    <template v-slot:default="{ isHovering, props: hoverProps }">
      <v-card
        v-bind="hoverProps"
        :color="isHovering ? colorScheme.hover : undefined"
        variant="flat"
        style="margin-left: 24px; padding: 6px"
      >
        <div
          :class="{ closing_bracket: props.node[1].type === 'object' }"
          :style="`color: ${isHovering ? colorScheme.font : 'black'}`"
        >
          <div
            :class="`d-flex align-center ${
              props.node[1].type === 'object' ? 'opening_bracket pb-1' : ''
            }`"
          >
            <p class="mr-2">
              <b>{{ props.node[0] }}: </b>
            </p>

            <VAutocomplete
              v-if="props.node[1].type !== 'object'"
              v-model="props.node[1].csvHeaderIndex"
              :items="headerOpts"
              label="Headers"
              item-title="header"
              item-value="index"
              style="max-width: 300px"
              hide-details
              clearable
              persistent-clear
              :multiple="props.node[1].type === 'array'"
            />
          </div>

          <ModelNode
            v-if="node[1].properties"
            v-for="([k, v], i) in Object.entries(node[1].properties)"
            :key="`${level + 1}-csv-${k}`"
            :node="[k, v]"
            :level="props.level + 1"
          />
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
