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
  <div class="program-bar">
    <b class="ml-4">StructifyCSV</b>
    <button class="ml-auto program-button" @click="MinimizeWindow">
      <v-icon size="x-small"> mdi-window-minimize </v-icon>
    </button>
    <button class="program-button" @click="toggleMaximize">
      <v-icon size="x-small">
        {{ isMaximized ? "mdi-window-restore" : "mdi-window-maximize" }}
      </v-icon>
    </button>
    <button class="program-button" @click="ExitProgram">
      <v-icon size="x-small"> mdi-window-close </v-icon>
    </button>
  </div>
</template>

<style scoped>
.program-bar {
  position: relative;
  display: flex;
  flex-direction: row;
  align-items: center;
  height: 32px;
  background-color: #e0e0e0;
}
.program-button {
  width: 36px;
}

.program-button:hover {
  background-color: #eeeeee;
}
</style>
