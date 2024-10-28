<template>
<div class="wayd-iconic">
    <svg xmlns="http://www.w3.org/2000/svg" 
        :width="width" 
        :height="height" 
        fill="currentColor" 
        class="bi bi-house wayd-icon" 
        viewBox="0 0 16 16"
        @click="iconClicked"
        >
        <path d="M8.707 1.5a1 1 0 0 0-1.414 0L.646 8.146a.5.5 0 0 0 .708.708L2 8.207V13.5A1.5 1.5 0 0 0 3.5 15h9a1.5 1.5 0 0 0 1.5-1.5V8.207l.646.647a.5.5 0 0 0 .708-.708L13 5.793V2.5a.5.5 0 0 0-.5-.5h-1a.5.5 0 0 0-.5.5v1.293zM13 7.207V13.5a.5.5 0 0 1-.5.5h-9a.5.5 0 0 1-.5-.5V7.207l5-5z"/>
    </svg>    
    <span class="wayd-tooltip">home</span>
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