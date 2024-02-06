<script lang="ts" setup>
const props = defineProps({
  showAddButton: {
    type: Boolean,
    default: true,
  },
});
const emit = defineEmits(["showAddForm", "showEditForm", "deleteProperty"]);
</script>

<template>
  <div class="d-flex">
    <v-btn
      v-if="props.showAddButton"
      size="x-small"
      class="ml-4"
      @click="emit('showAddForm')"
    >
      add
    </v-btn>
    <v-btn size="x-small" class="ml-4" @click="emit('showEditForm')">
      edit
    </v-btn>
    <v-btn size="x-small" class="ml-4" @click="emit('deleteProperty')">
      delete
    </v-btn>
  </div>
</template>
