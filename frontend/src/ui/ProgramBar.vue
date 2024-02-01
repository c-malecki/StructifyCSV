<script lang="ts" setup>
import { ref } from "vue";

const emit = defineEmits(["new", "import", "export"]);

const showMenu = ref(false);
const clickOutside = () => {
  showMenu.value = false;
};

defineExpose({ showMenu });
</script>

<template>
  <div class="program-bar">
    <button :class="{ active: showMenu }" @click="showMenu = !showMenu">File</button>
    <div v-if="showMenu" class="menu" v-click-outside="clickOutside">
      <ul>
        <li>
          <button @click="emit('new')">new schema</button>
        </li>
        <li>
          <button @click="emit('import')">import schema</button>
        </li>
        <li>
          <button @click="emit('export')">export schema</button>
        </li>
      </ul>
    </div>
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

button {
  padding: 0 1rem;
}

button:hover {
  background-color: #eeeeee;
}
.active {
  background-color: #eeeeee;
}

.menu {
  position: absolute;
  width: 400px;
  height: 200px;
  background-color: #eeeeee;
  top: 32px;
  z-index: 2;
}
</style>
