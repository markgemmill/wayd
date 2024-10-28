<template>
<div v-if="showModal" class="modal-backdrop">
    <div ref="modal" class="modal-form">
        <ExitIcon class="exit-icon" :width="20" :height="20" @clicked="onExit"></ExitIcon>
        <slot></slot>
    </div>
</div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { onClickOutside } from '@vueuse/core';
import ExitIcon from '../icons/ExitIcon.vue';

const modal = ref()
const showModal = ref(false)

const onExit = () => {
    showModal.value = false 
}
const onDisplay = () => {
    showModal.value = true 
}

onClickOutside(modal, onExit)

defineExpose({
    "hide": onExit,
    "show": onDisplay,
})

const emit = defineEmits<{
    "on-exit": []
}>()

</script>

<style scoped>
.exit-icon {
    position: absolute;
    top: 5px;
    right: 5px;
}
.modal-backdrop {
    position: absolute;
    top: 0;
    left: 0;
    background-color: transparent;
    z-index: 99;
    width: 100%;
    height: 100%;
    padding-top: 25px;
}
.modal-form {
    position: relative;
    margin: 10px auto;
    display: flex;
    flex-direction: column;
    align-items: start;
    z-index: 100;
    background-color: white;
    border: 1px solid green;
    border-radius: 5px;
    width: 80%;
    padding: 20px;
}
.exit-icon {
    position: absolute;
    top: 10px;
    right: 10px;
}
</style>