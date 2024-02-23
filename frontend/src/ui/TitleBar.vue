<script lang="ts" setup>
import { ref } from "vue";
import {
  MinimizeWindow,
  UnmaximizeWindow,
  MaximizeWindow,
  ExitProgram,
} from "../../wailsjs/go/main/App";

const isMaximized = ref(false);
const toggleMaximize = () => {
  if (isMaximized.value) {
    UnmaximizeWindow();
  } else {
    MaximizeWindow();
  }
  isMaximized.value = !isMaximized.value;
};
</script>

<template>
  <v-toolbar height="36" color="blue-grey-lighten-3">
    <v-toolbar-title> StructifyCSV </v-toolbar-title>
    <v-spacer />
    <v-btn icon="mdi-window-minimize" size="small" @click="MinimizeWindow" />
    <v-btn
      :icon="isMaximized ? 'mdi-window-restore' : 'mdi-window-maximize'"
      size="small"
      @click="toggleMaximize"
    />
    <v-btn icon="mdi-window-close" size="small" @click="ExitProgram" />
  </v-toolbar>
</template>
