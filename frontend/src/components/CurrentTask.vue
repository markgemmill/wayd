<template>
    <div class="column current-entry">
        <div class="form">
            <div class="mb-3">
                <TimeClock :format="'clock'" :start-time="startTime"></TimeClock>
            </div>
            <div class="mb-3">
                <label for="entry-project" class="form-label fw-bold">Project</label>
                <div id="entry-project" class="project fs-3">{{  name }}</div>
            </div>
            <div>
                <label for="entry-comments" class="form-label fw-bold">Comments</label>
                <textarea id="entry-comments" class="form-control" rows="5" v-model="comment"></textarea>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, watch, computed, onMounted } from 'vue'
import dayjs from "dayjs"
import type { Dayjs } from "dayjs"
import TimeClock from "./TimeClock.vue"
import * as log from "../../bindings/wayd/services/loggerservice"

const { start, name, comments } = defineProps<{
    start: string 
    name: string
    comments: string
}>()

const comment = ref("")

const emit = defineEmits<{
    "comment-changed": [string]
}>()

const startTime = computed<Dayjs>(() => {
    return dayjs(start) 
})

onMounted(() => {
    comment.value = comments 
})

watch(comment, (newValue: string) => {
    log.Debug(newValue)
    emit("comment-changed", newValue)
})
</script>

<style scoped>
/* .current-entry {
    width: 100%;
    border: 1px solid #CCC;
    border-radius: 3px; 
    padding: 10px;
    margin-bottom: 10px;
}

.clock {
    margin-bottom: 10px;
}

.label {
    font-size: 0.9em;
    font-weight: bold;
    color: #666;
    text-align: left;
    margin-left: 3px;
}
.project {
    font-size: 1.2em;
    color: #333;
    text-align: left;
    margin-left: 3px;
}
.entry-note {
    width: 100%;
    outline: none !important;
    border: 1px solid #DDD;
    border-radius: 3px;;
} */
</style>