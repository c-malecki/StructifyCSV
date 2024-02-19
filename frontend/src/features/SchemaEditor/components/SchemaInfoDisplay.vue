<script lang="ts" setup>
import { entity } from "../../../../wailsjs/go/models";
import { type PropType } from "vue";

const props = defineProps({
  schema: {
    type: Object as PropType<entity.JsonSchema>,
    required: true,
  },
});

const emit = defineEmits(["showEdit"]);
</script>

<template>
  <div class="relative pa-4">
    <h3>{{ props.schema.title }}</h3>
    <p>{{ props.schema.description }}</p>
    <v-btn
      position="absolute"
      size="small"
      location="top right"
      @click="emit('showEdit')"
    >
      edit
    </v-btn>
  </div>
</template>

<style scoped>
.hide {
  display: none;
}
.relative {
  position: relative;
}
</style>
