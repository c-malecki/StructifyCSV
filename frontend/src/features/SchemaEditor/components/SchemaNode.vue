<script lang="ts" setup>
import { ref, computed, type PropType } from "vue";
import { getHoverColorScheme, getLeftIndent } from "../../../util/style";
import {
  type SchemaPropertiesMap,
  type PropertiesMapValue,
} from "../../../types/editor.types";
import SchemaNodeButtons from "./SchemaNodeButtons.vue";
import AddPropertyForm from "./AddPropertyForm.vue";
import EditPropertyForm from "./EditPropertyForm.vue";
import { type JsonSchemaDataType } from "../../../types/schema.types";

type Node = [string, PropertiesMapValue];

const nodeProps = defineProps({
  node: {
    type: Object as PropType<Node>,
    required: true,
  },
  level: {
    type: Number,
    default: 1,
  },
});
const emit = defineEmits([
  "updateBaseKey",
  "updateBaseValue",
  "deleteBaseProperty",
  "updateParentKey",
  "updateParentValue",
  "deleteParentProperty",
]);

const showAddForm = ref(false);
const showEditForm = ref(false);

const isMap = computed(() => nodeProps.node[1] instanceof Map);
const leftIndent = computed(() => getLeftIndent(nodeProps.level));
const colorScheme = computed(() => getHoverColorScheme(nodeProps.level));

const updateKey = (update: {
  newKey: string;
  oldKey: string;
  sameValue: Node;
}) => {
  if (nodeProps.level === 1) {
    emit("updateBaseKey", update);
  } else {
    emit("updateParentKey", update);
  }
};

const updateParentKey = ({
  newKey,
  oldKey,
  sameValue,
}: {
  newKey: string;
  oldKey: string;
  sameValue: PropertiesMapValue;
}) => {
  (nodeProps.node[1] as SchemaPropertiesMap).delete(oldKey);
  (nodeProps.node[1] as SchemaPropertiesMap).set(newKey, sameValue);
};

const updateValue = (update: {
  sameKey: string;
  newValue: JsonSchemaDataType | SchemaPropertiesMap;
}) => {
  if (nodeProps.level === 1) {
    emit("updateBaseValue", update);
  } else {
    emit("updateParentValue", update);
  }
};

const updateParentValue = ({
  sameKey,
  newValue,
}: {
  sameKey: string;
  newValue: JsonSchemaDataType | SchemaPropertiesMap;
}) => {
  (nodeProps.node[1] as SchemaPropertiesMap).set(sameKey, newValue);
};

const deleteProperty = (keyToDelete: string) => {
  if (nodeProps.level === 1) {
    emit("deleteBaseProperty", keyToDelete);
  } else {
    emit("deleteParentProperty", keyToDelete);
  }
};

const deleteParentProperty = (keyToDelete: string) => {
  // will need to account for array and null?
  const isObject =
    typeof (nodeProps.node[1] as SchemaPropertiesMap).get(keyToDelete) ===
    "object";
  let message = "";
  switch (isObject) {
    case true:
      message = `Deleting "${keyToDelete}" will also delete any descendents of "${keyToDelete}." Do you wish to proceed?`;
      break;
    case false:
      message = `Deleting "${keyToDelete}" cannot be undone. Do you wish to proceed?`;
      break;
  }
  if (confirm(message)) {
    (nodeProps.node[1] as SchemaPropertiesMap).delete(keyToDelete);
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
          class="pa-1"
        >
          <div v-if="!isMap" class="d-flex">
            <p v-if="!showEditForm">
              <b>{{ node[0] }}:</b> {{ node[1] }}
            </p>
            <SchemaNodeButtons
              v-if="isHovering && !showEditForm"
              :show-add-button="false"
              @show-edit-form="showEditForm = true"
              @delete-property="deleteProperty(node[0])"
            />
            <EditPropertyForm
              v-if="showEditForm"
              :node="node"
              @update-key="updateKey"
              @update-value="updateValue"
              @close-form="showEditForm = false"
            />
          </div>

          <div v-else>
            <div class="d-flex opening_bracket pb-1">
              <p v-if="!showEditForm">
                <b>{{ node[0] }}:</b>
              </p>
              <SchemaNodeButtons
                v-if="isHovering && !showEditForm"
                @show-add-form="showAddForm = true"
                @show-edit-form="showEditForm = true"
                @delete-property="deleteProperty(node[0])"
              />
              <EditPropertyForm
                v-if="showEditForm"
                :node="node"
                @update-key="updateKey"
                @update-value="updateValue"
                @close-form="showEditForm = false"
              />
            </div>
            <AddPropertyForm
              v-if="showAddForm"
              :nodeValue="(node[1] as Map<string, PropertiesMapValue>)"
              @hideForm="showAddForm = false"
            />
            <SchemaNode
              v-for="(n, i) in node[1]"
              :key="`${level + 1}-${i}-${typeof node[1]}`"
              :node="(n as Node)"
              :level="level + 1"
              @update-parent-key="updateParentKey"
              @update-parent-value="updateParentValue"
              @delete-parent-property="deleteParentProperty"
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
</style>
