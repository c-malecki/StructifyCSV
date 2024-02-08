<script lang="ts" setup>
import { reactive } from "vue";

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
    </div>
    <!-- 
    <button
      :class="`menu-button included ${
        menuControl.menu === 'csv' ? 'active' : ''
      }`"
      @click="toggleMenu('csv')"
    >
      CSV
    </button>
    <div
      :class="`menu csv-menu included ${
        menuControl.show && menuControl.menu === 'csv' ? '' : 'hide'
      }`"
      v-click-outside="{
        handler: clickOutside,
        include,
      }"
    >
      <v-list bg-color="grey-lighten-3" density="compact">
        <v-list-item
          title="Import"
          prepend-icon="mdi-table-arrow-left"
          @click="emit('importCsv')"
        />
      </v-list>
    </div> -->
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
  padding: 0 1rem;
  width: 90px;
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

.csv-menu {
  left: 90px;
}

ul {
  list-style: none;
  width: 200px;
  padding: 0.5rem 0;
}

ul li {
  padding: 0.2rem 0.5rem;
}

ul li:hover {
  background-color: #f5f5f5;
}
</style>
