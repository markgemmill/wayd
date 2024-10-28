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
        <path d="M1 3.5A1.5 1.5 0 0 1 2.5 2h2.764c.958 0 1.76.56 2.311 1.184C7.985 3.648 8.48 4 9 4h4.5A1.5 1.5 0 0 1 15 5.5v7a1.5 1.5 0 0 1-1.5 1.5h-11A1.5 1.5 0 0 1 1 12.5zM2.5 3a.5.5 0 0 0-.5.5V6h12v-.5a.5.5 0 0 0-.5-.5H9c-.964 0-1.71-.629-2.174-1.154C6.374 3.334 5.82 3 5.264 3zM14 7H2v5.5a.5.5 0 0 0 .5.5h11a.5.5 0 0 0 .5-.5z"/>
    </svg>    
    <span class="tooltip">projects</span>
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
/* icon:hover {
    color: black;
    cursor: pointer;
} */
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