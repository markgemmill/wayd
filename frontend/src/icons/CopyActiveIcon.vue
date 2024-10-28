<template>
<div class="icon">
    <svg xmlns="http://www.w3.org/2000/svg" 
        width="16" 
        height="16" 
        fill="currentColor" 
        class="bi bi-arrow-clockwise timer" 
        viewBox="0 0 16 16"
        >
        <path fill-rule="evenodd" d="M8 3a5 5 0 1 0 4.546 2.914.5.5 0 0 1 .908-.417A6 6 0 1 1 8 2z"/>
        <path d="M8 4.466V.534a.25.25 0 0 1 .41-.192l2.36 1.966c.12.1.12.284 0 .384L8.41 4.658A.25.25 0 0 1 8 4.466"/>
    </svg>
    <div class="counter">{{ countDown }}</div>
</div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

const countDown = ref(0)

const props = defineProps<{
    countDown: number
}>()

onMounted(() => {
    countDown.value = props.countDown 
    const intervalId = setInterval(() => {
        countDown.value -= 1 
        if (countDown.value == 0) {
            clearInterval(intervalId)
        }
    }, 1000)
})

</script>

<style scoped>
.icon {
    display: flex;
    flex-direction: row;
    flex-grow: 0;
    align-items: center;
    border-radius: 4px;
    border: 1px solid transparent;
    padding: 3px 4px;
    background-color: transparent;
    vertical-align: middle;
}
.timer {
    rotate: 0deg;
    flex-grow: 0;
    animation: rotate 1s infinite linear;
    color: white;
    height: 32px;
    width: 32px;
}
.counter {
    color: white;
    flex-grow: 0;
    font-size: 0.7em;
    margin-left: 0.20em;
    min-width: 1.25em;
    text-align: end;
}

@keyframes rotate {
    0% {
        rotate: 0deg;
    }
    50% {
        rotate: 180deg;
    }
    100% {
        rotate: 359deg;
    }
}
</style>