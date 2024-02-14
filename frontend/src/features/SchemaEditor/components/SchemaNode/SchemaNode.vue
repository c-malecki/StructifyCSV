<script lang="ts" setup>
import { ref, computed, type PropType } from "vue";
import { getHoverColorScheme } from "../../../../util/style";
import { getSchemaAttributesDisplay } from "../../../../util/transform";
import { type SchemaNode } from "../../SchemaEditor.types";
import { entity } from "../../../../../wailsjs/go/models";
import SchemaNodeButtons from "./SchemaNodeButtons.vue";
import AddPropertyForm from "../forms/AddPropertyForm.vue";
import EditPropertyForm from "../forms/EditPropertyForm.vue";
import EditRequiredForm from "../forms/EditRequiredForm.vue";

type CurForm = "add" | "edit" | "required" | null;

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

const curForm = ref<CurForm>(null);
const showForm = (form: CurForm) => {
  curForm.value = form;
};

const colorScheme = computed(() => getHoverColorScheme(props.level));

const propertyAttributes = computed(() =>
  getSchemaAttributesDisplay(props.node[1])
);

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
  value: entity.SchemaProperty;
}) => {
  delete props.node[1].properties[curKey];
  props.node[1].properties[editKey] = value;
};

const updateValue = (update: {
  editKey: string;
  curKey: string;
  value: entity.SchemaProperty;
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
  value: entity.SchemaProperty;
}) => {
  if (editKey !== curKey) {
    delete props.node[1].properties[curKey];
  }
  props.node[1].properties[editKey] = value;
};

const deleteProperty = (keyToDelete: string) => {
  if (props.level === 1) {
    emit("deleteBaseProperty", keyToDelete);
  } else {
    emit("deleteParentProperty", keyToDelete);
  }
};

const deleteParentProperty = (keyToDelete: string) => {
  const isObjOrArr =
    props.node[1].type === "object" || props.node[1].type === "array";
  let message = "";
  switch (isObjOrArr) {
    case true:
      message = `Deleting "${keyToDelete}" will also delete any descendents of "${keyToDelete}." Do you wish to proceed?`;
      break;
    case false:
      message = `Deleting "${keyToDelete}" cannot be undone. Do you wish to proceed?`;
      break;
  }
  if (confirm(message)) {
    delete props.node[1].properties[keyToDelete];
  }
};

const addNewProperty = ({
  key,
  value,
}: {
  key: string;
  value: entity.SchemaProperty;
}) => {
  if (props.node[1].properties !== undefined) {
    props.node[1].properties[key] = value;
  } else {
    props.node[1].properties = { [key]: value };
  }
};

const updateRequired = (required: string[]) => {
  props.node[1].required = required;
  curForm.value = null;
};
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
          :class="{ closing_bracket: node[1].type === 'object' }"
          :style="`color: ${isHovering ? colorScheme.font : 'black'}`"
        >
          <div
            :class="`d-flex ${
              node[1].type === 'object' ? 'opening_bracket pb-1' : ''
            }`"
          >
            <p
              v-if="
                curForm !== 'edit' &&
                node[1].type !== 'object' &&
                node[1].type !== 'array'
              "
            >
              <b>{{ node[0] }}:</b> {{ node[1].type }}
            </p>

            <p v-if="curForm !== 'edit' && node[1].type === 'array'">
              <b>{{ node[0] }}:</b> {{ node[1].type }}
              {{ node[1].items ? `[ ${node[1].items.type} ]` : "" }}
            </p>

            <p v-if="curForm !== 'edit' && node[1].type === 'object'">
              <b>{{ node[0] }}:</b>
            </p>

            <SchemaNodeButtons
              v-if="isHovering && !curForm"
              :is-object-property="node[1].type === 'object'"
              @show-form="showForm"
              @delete-property="deleteProperty(node[0])"
            />
            <EditPropertyForm
              v-if="curForm === 'edit'"
              :node="node"
              @update-key="updateKey"
              @update-value="updateValue"
              @close-form="curForm = null"
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

          <SchemaNode
            v-if="node[1].properties && curForm !== 'required'"
            v-for="([k, v], i) in Object.entries(node[1].properties)"
            :key="`${level + 1}-json-${k}`"
            :node="[k, v]"
            :level="level + 1"
            @update-parent-key="updateParentKey"
            @update-parent-value="updateParentValue"
            @delete-parent-property="deleteParentProperty"
          />

          <AddPropertyForm
            v-if="node[1].type === 'object' && curForm === 'add'"
            @close-form="curForm = null"
            @add-new-property="addNewProperty"
          />

          <EditRequiredForm
            v-if="node[1].type === 'object' && curForm === 'required'"
            :schema="node[1]"
            @close-form="curForm = null"
            @update-required="updateRequired"
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
