<template>
<div class="wayd-iconic">
  <svg xmlns="http://www.w3.org/2000/svg" 
    :width="width" 
    :height="height" 
    fill="currentColor" 
    class="bi bi-file-earmark-bar-graph wayd-icon" 
    viewBox="0 0 16 16"
    @click="iconClicked"
    >
        <path d="M10 13.5a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5h-1a.5.5 0 0 0-.5.5zm-2.5.5a.5.5 0 0 1-.5-.5v-4a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-.5.5zm-3 0a.5.5 0 0 1-.5-.5v-2a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5z"/>
        <path d="M14 14V4.5L9.5 0H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2M9.5 3A1.5 1.5 0 0 0 11 4.5h2V14a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1h5.5z"/>
    </svg>
  <span class="wayd-tooltip">reporting</span>
</div>
</template>

<script setup lang="ts">
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
</style>
