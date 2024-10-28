<template>
   <div class="timer" :class="class_" @click="onClick">{{ displayTime }}</div> 
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted, onUnmounted } from 'vue'
import dayjs from "dayjs"
import type { Dayjs } from "dayjs"

let timerId = null
const startValue = "00:00:00"
const displayTime = ref<string>(startValue)

const { startTime, format="clock" }= defineProps<{
    startTime: Dayjs | null 
    format: string
}>()

const emit = defineEmits<{
    "clicked": []
}>()

const class_ = computed(() => {
    return {
        "clock": format === "clock",
        "timer-icon": format === "icon"
    } 
}) 

const SECONDS_IN_HOUR = 60 * 60
const SECONDS_IN_MIN = 60 

const elapsedTime = () => {
    if (!startTime) {
        displayTime.value = startValue
    }
    const currentTime: Dayjs = dayjs() 
    let seconds: number = currentTime.diff(startTime, "seconds") 

    // calculate hours
    const hours = Math.floor(seconds / SECONDS_IN_HOUR)
    const hour_seconds = hours * SECONDS_IN_HOUR
    if (hours > 0) {
        seconds = seconds - hour_seconds
    }
    // calculate minutes
    const minutes = Math.floor(seconds / SECONDS_IN_MIN)
    const minute_seconds = minutes * SECONDS_IN_MIN
    if (minutes > 0) {
        seconds = seconds - minute_seconds 
    }

    displayTime.value = `${hours.toString().padStart(2, "0")}:${minutes.toString().padStart(2, "0")}:${seconds.toString().padStart(2, "0")}`
}

const onClick = () => {
    if (format === "icon") {
        emit("clicked")    
    }
}

watch(() => startTime, () => {
    if (startTime) {
        timerId = setInterval(elapsedTime, 1000) 
    } else if (timerId) {
        clearInterval(timerId) 
        timerId = null 
        displayTime.value = startValue 
    }
})

onMounted(() => {
    if (startTime) {
        timerId = setInterval(elapsedTime, 1000) 
    } else {
        timerId = null 
    }
})

onUnmounted(() => {
    if (timerId) {
        clearInterval(timerId) 
    }
})


</script>

<style scoped>
.timer {
    font-family: monospace;
    text-align: center;
    vertical-align: middle;
    color: #333;
    border: 1px solid #CCC;
    padding: 0;
    margin: 0;
}

.clock {
    width: 100%;
    font-size: 4em;
    line-height: normal;
    padding: 20px 10px;
    border: 2px solid #444;
    border-radius: 3px;
    background-color: #777;
    color: white;
}

.timer-icon {
    position: relative;
    top: 3px;
    font-size: 0.7em;
    line-height: normal;
    color: black;
    border-radius: 2px;
    margin:  auto 0px; 
    padding: 3px;
}
.timer-icon:hover {
    background-color: black;
    color: white;
    cursor: pointer;
}
</style>