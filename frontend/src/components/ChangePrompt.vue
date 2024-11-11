<template>
    <div class="column">
        <div class="column prompt">
            <div class="row mb-1">
                <h1>What Are You Doing?</h1>
            </div>
            <div class="row mb-3">
                <label class="form-label">Current Task</label>
                <div v-if="currentEntry">{{ currentEntry }}</div>
                <div v-else >No current entry running.</div>
            </div>
            <div class="row mb-2">
                <label for="delay-prompt" class="form-label">Don't Remind Me For...</label>
                <select id="delay-prompt" class="form-select" v-model="delayPrompt">
                    <option value="1" selected>1 minute</option>
                    <option value="2" selected>2 minute</option>
                    <option value="5">5 minutes</option>
                    <option value="10">10 minutes</option>
                    <option value="15">15 minutes</option>
                    <option value="30">30 minutes</option>
                    <option value="45">45 minutes</option>
                    <option value="60">1 hour</option>
                    <option value="120">2 hour</option>
                    <option value="180">3 hour</option>
                    <option value="240">4 hour</option>
                </select>
            </div>
            <div class="row ms-1 me-1">
                <button type="button" class="btn btn-sm btn-primary" @click="handleClose">Close</button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import * as runtime from '@wailsio/runtime';
import { useStore } from '../store';
import * as log from "../../bindings/wayd/services/loggerservice"
import { router } from "../router";

const store = useStore()
const currentEntry = ref("")
const delayPrompt = ref<number>(1)

const handleClose = () => {
    log.Debug(`[GUI ALERT] closing prompt... ${delayPrompt.value}`)
    const delay = Number.parseInt(delayPrompt.value.toString()).toFixed(0)

    log.Debug(`[GUI ALERT] closing prompt... delay of ${delay} minute(s)...`)
    const event = new runtime.Events.WailsEvent("close-prompt", { Delay: delay})

    log.Debug(`[GUI ALERT] emiting event... delay of ${delay} minute(s)...`)
    runtime.Events.Emit(event)

    router.push("/current")
}

onMounted(() => {
    if (store.currentEntryIsZero) {
        currentEntry.value = ""
    } else {
        currentEntry.value = store.currentEntry.Project.Name 
    }
})

</script>

<style scoped>
</style>