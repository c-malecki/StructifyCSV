<script lang="ts" setup>
import ModelDisplay from "./components/ModelDisplay.vue";
import SelectSchema from "./components/SelectSchema.vue";
import CreateSchema from "./components/CreateSchema.vue";
import EditSchema from "./components/EditSchema.vue";
</script>

<template>
  <div class="row">
    <div class="container col">
      <ModelDisplay />
    </div>
    <div class="container col">
      <CreateSchema />
      <SelectSchema />
      <EditSchema />
    </div>
  </div>
</template>

<style scoped></style>
