<template>
<div class="iconic">
    <svg xmlns="http://www.w3.org/2000/svg" 
        :width="width" 
        :height="height" 
        fill="currentColor" 
        class="bi bi-folder2 icon" 
        viewBox="0 0 16 16"
        @click="iconClicked"
        >
        <slot></slot>
    </svg>    
    <span class="tooltip">{{ tooltip }}</span>
</div>
</template>

<script setup lang="ts">

const { tooltip, width=16, height=16 }= defineProps<{
    tooltip: string
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
.iconic {
    position: relative;
    display: block;
    /* border: 1px solid red; */
}

.tooltip {
    visibility: hidden;
    position: absolute; 
    display: inline-block;
    bottom: -120%;
    right: -100%;
    z-index: 10;
    font-size: 0.65em;
    border: 1px solid #CCC;
    border-radius: 3px;
    padding: 2px;
    background-color: white;
    width: 300%;
    height: auto;
    padding: 0px 3px;
    text-align: center;
}

.icon {
    position: relative;
    top: 3px;
    color: #888;
    margin: auto 0;
}
icon:hover {
    color: black;
    cursor: pointer;
}
.iconic:hover .icon {
    color: black;
    cursor: pointer;
}
.iconic:hover .tooltip {
    visibility: visible;
    transform: 1s all ease;
    transition-delay: 1s;

}
</style>