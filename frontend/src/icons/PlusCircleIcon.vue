<template>
<svg xmlns="http://www.w3.org/2000/svg" 
     :width="width" 
     :height="height" 
     fill="currentColor" 
     class="bi bi-plus-circle icon" 
     viewBox="0 0 16 16"
     @click="iconClicked"
     >
  <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14m0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16"/>
  <path d="M8 4a.5.5 0 0 1 .5.5v3h3a.5.5 0 0 1 0 1h-3v3a.5.5 0 0 1-1 0v-3h-3a.5.5 0 0 1 0-1h3v-3A.5.5 0 0 1 8 4"/>
</svg>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'

const { width=16, height=16 }= defineProps<{
    width: number
    height: number
}>()

const emit = defineEmits<{
    "clicked": []
}>()

const iconClicked = () => {
    emit("clicked") 
}
</script>

<style scoped>
.icon {
    color: #888;
}
.icon:hover {
    color: black;
    cursor: pointer;
}
</style>