<script lang="ts" setup>
import { ref, reactive } from "vue";
import {
  MinimizeWindow,
  UnmaximizeWindow,
  MaximizeWindow,
  ExitProgram,
} from "../../wailsjs/go/main/App";

type MenuControl = {
  show: boolean;
  menu: "schema" | "csv" | null;
};

const emit = defineEmits([
  "newSchema",
  "importSchema",
  "exportSchema",
  "importCsv",
]);

const isMaximized = ref(false);
const toggleMaximize = () => {
  if (isMaximized.value) {
    UnmaximizeWindow();
  } else {
    MaximizeWindow();
  }
  isMaximized.value = !isMaximized.value;
};

const menuControl = reactive<MenuControl>({
  show: false,
  menu: null,
});

const toggleMenu = (menu: MenuControl["menu"]) => {
  if (menuControl.show && menuControl.menu === menu) {
    menuControl.show = false;
    menuControl.menu = null;
  } else {
    menuControl.menu = menu;
    menuControl.show = true;
  }
};

const clickOutside = () => {
  menuControl.show = false;
  menuControl.menu = null;
};

const include = () => {
  return Array.from(document.getElementsByClassName("included"));
};

defineExpose({ menuControl });
</script>

<template>
  <div class="program-bar">
    <button
      :class="`menu-button included ${
        menuControl.menu === 'schema' ? 'active' : ''
      }`"
      @click="toggleMenu('schema')"
    >
      File
    </button>
    <div
      :class="`menu included ${
        menuControl.show && menuControl.menu === 'schema' ? '' : 'hide'
      }`"
      v-click-outside="{
        handler: clickOutside,
        include,
      }"
    >
      <v-list bg-color="grey-lighten-3" density="compact">
        <v-list-item
          title="New Schema"
          prepend-icon="mdi-file-plus-outline"
          @click="emit('newSchema')"
        />
        <v-list-item
          title="Open Schema"
          prepend-icon="mdi-folder-open-outline"
          @click="emit('importSchema')"
        />
        <!-- <v-list-item
          title="Save"
          prepend-icon="mdi-content-save-edit-outline"
          @click="emit('exportSchema')"
        /> -->
        <v-list-item
          title="Save Schema as"
          prepend-icon="mdi-content-save-edit-outline"
          @click="emit('exportSchema')"
        />
        <v-list-item
          title="Import CSV"
          prepend-icon="mdi-table-arrow-left"
          @click="emit('importCsv')"
        />
      </v-list>

      <v-spacer />
    </div>
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
  height: 32px;
  background-color: #e0e0e0;
}

.menu-button {
  width: 60px;
}

.program-button {
  width: 40px;
}

.program-button:hover {
  background-color: #eeeeee;
}

.menu-button:hover {
  background-color: #eeeeee;
}
.active {
  background-color: #eeeeee;
}

.hide {
  visibility: hidden;
}
.menu {
  position: absolute;
  background-color: #eeeeee;
  top: 32px;
  z-index: 2;
}
</style>
