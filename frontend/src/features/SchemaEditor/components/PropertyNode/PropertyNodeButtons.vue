<script lang="ts" setup>
const props = defineProps({
  isObjectProperty: {
    type: Boolean,
    required: true,
  },
});
const emit = defineEmits(["showForm", "deleteProperty"]);
</script>

<template>
  <div class="d-flex">
    <v-btn
      size="x-small"
      class="ml-4"
      prepend-icon="mdi-delete-outline"
      @click="emit('deleteProperty')"
    >
      delete
    </v-btn>
    <v-btn
      size="x-small"
      class="ml-4"
      prepend-icon="mdi-pencil-box-outline"
      @click="emit('showForm', 'edit')"
    >
      edit
    </v-btn>
    <v-btn
      v-if="props.isObjectProperty"
      size="x-small"
      class="ml-4"
      prepend-icon="mdi-plus"
      @click="emit('showForm', 'add')"
    >
      add
    </v-btn>
    <v-btn
      v-if="props.isObjectProperty"
      size="x-small"
      class="ml-4"
      prepend-icon="mdi-pencil-box-outline"
      @click="emit('showForm', 'required')"
    >
      required
    </v-btn>
  </div>
</template>
