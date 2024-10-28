<template>
    <svg xmlns="http://www.w3.org/2000/svg" 
         :width="width" 
         :height="height" 
         fill="currentColor" 
         class="bi bi-x icon" 
         viewBox="0 0 16 16"
         @click="iconClicked">
        <path d="M4.646 4.646a.5.5 0 0 1 .708 0L8 7.293l2.646-2.647a.5.5 0 0 1 .708.708L8.707 8l2.647 2.646a.5.5 0 0 1-.708.708L8 8.707l-2.646 2.647a.5.5 0 0 1-.708-.708L7.293 8 4.646 5.354a.5.5 0 0 1 0-.708"/>
    </svg>
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
.icon {
    position: relative;
    top: 3px;
    color: #888;
    margin: auto 0;
}
.icon:hover {
    color: black;
    cursor: pointer;
}
</style>