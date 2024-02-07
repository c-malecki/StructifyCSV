<script lang="ts" setup>
import { ref, computed, type PropType } from "vue";
import { getHoverColorScheme, getLeftIndent } from "../../../../util/style";
import SchemaNodeButtons from "./SchemaNodeButtons.vue";
import AddPropertyForm from "../AddPropertyForm.vue";
import EditPropertyForm from "../EditPropertyForm.vue";
import {
  type SchemaProperty,
  ObjectProperty,
  ArrayProperty,
} from "../../../../types/properties.types";
import { type SchemaNode } from "../../../../types/editor.types";

const props = defineProps({
  node: {
    type: Object as PropType<SchemaNode>,
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

const isObjectProperty = computed(
  () => props.node[1] instanceof ObjectProperty
);
const isArrayProperty = computed(() => props.node[1] instanceof ArrayProperty);

const leftIndent = computed(() => getLeftIndent(props.level));
const colorScheme = computed(() => getHoverColorScheme(props.level));

const propertyAttributes = computed(() => props.node[1].getAttributeDisplay());

const updateKey = (update: {
  editKey: string;
  curKey: string;
  value: SchemaNode;
}) => {
  if (props.level === 1) {
    emit("updateBaseKey", update);
  } else {
    emit("updateParentKey", update);
  }
};

const updateParentKey = ({
  editKey,
  curKey,
  value,
}: {
  editKey: string;
  curKey: string;
  value: SchemaProperty;
}) => {
  (props.node[1] as ObjectProperty).properties.delete(curKey);
  (props.node[1] as ObjectProperty).properties.set(editKey, value);
};

const updateValue = (update: {
  editKey: string;
  curKey: string;
  value: SchemaNode;
}) => {
  if (props.level === 1) {
    emit("updateBaseValue", update);
  } else {
    emit("updateParentValue", update);
  }
};

const updateParentValue = ({
  editKey,
  curKey,
  value,
}: {
  editKey: string;
  curKey: string;
  value: SchemaProperty;
}) => {
  if (editKey !== curKey) {
    (props.node[1] as ObjectProperty).properties.delete(curKey);
  }
  (props.node[1] as ObjectProperty).properties.set(editKey, value);
};

const deleteProperty = (keyToDelete: string) => {
  if (props.level === 1) {
    emit("deleteBaseProperty", keyToDelete);
  } else {
    emit("deleteParentProperty", keyToDelete);
  }
};

const deleteParentProperty = (keyToDelete: string) => {
  const isObjorArr =
    props.node[1].type === "object" || props.node[1].type === "array";
  let message = "";
  switch (isObjorArr) {
    case true:
      message = `Deleting "${keyToDelete}" will also delete any descendents of "${keyToDelete}." Do you wish to proceed?`;
      break;
    case false:
      message = `Deleting "${keyToDelete}" cannot be undone. Do you wish to proceed?`;
      break;
  }
  if (confirm(message)) {
    (props.node[1] as ObjectProperty).properties.delete(keyToDelete);
  }
};

const addNewProperty = ({
  key,
  value,
}: {
  key: string;
  value: SchemaProperty;
}) => {
  (props.node[1] as ObjectProperty).properties.set(key, value);
};
</script>

<template>
  <v-hover>
    <template v-slot:default="{ isHovering, props: hoverProps }">
      <v-card
        v-bind="hoverProps"
        :color="isHovering ? colorScheme.hover : undefined"
        variant="flat"
        :style="leftIndent"
      >
        <div
          :class="{ closing_bracket: isObjectProperty }"
          :style="`color: ${isHovering ? colorScheme.font : 'black'}`"
        >
          <div
            :class="`d-flex pb-1 ${isObjectProperty ? 'opening_bracket' : ''}`"
          >
            <p v-if="!showEditForm && !isObjectProperty && !isArrayProperty">
              <b>{{ node[0] }}:</b> {{ node[1].type }}
            </p>

            <p v-if="!showEditForm && isArrayProperty">
              <b>{{ node[0] }}:</b> {{ node[1].type }}
              {{
                (node[1] as ArrayProperty).items
                  ? `[ ${(node[1] as ArrayProperty).items!.type} ]`
                  : ""
              }}
            </p>

            <p v-if="!showEditForm && isObjectProperty">
              <b>{{ node[0] }}:</b>
            </p>

            <SchemaNodeButtons
              v-if="isHovering && !showEditForm"
              :show-add-button="props.node[1] instanceof ObjectProperty"
              @show-edit-form="showEditForm = true"
              @show-add-form="showAddForm = true"
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

          <v-expand-transition>
            <div v-show="isHovering">
              <v-chip
                v-for="(attr, idx) in propertyAttributes"
                size="small"
                class="ma-1"
              >
                {{ attr[0] }}: {{ attr[1] }}
              </v-chip>
            </div>
          </v-expand-transition>

          <AddPropertyForm
            v-if="props.node[1] instanceof ObjectProperty && showAddForm"
            @close-form="showAddForm = false"
            @add-new-property="addNewProperty"
          />

          <SchemaNode
            v-if="props.node[1] instanceof ObjectProperty"
            v-for="(n, i) in (node[1] as ObjectProperty).properties"
            :key="`${level + 1}-${i}-${typeof node[1]}`"
            :node="(n as SchemaNode)"
            :level="level + 1"
            @update-parent-key="updateParentKey"
            @update-parent-value="updateParentValue"
            @delete-parent-property="deleteParentProperty"
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
</style>
