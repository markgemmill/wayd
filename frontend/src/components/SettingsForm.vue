<template>
   <div class="row w-100">
        <form class="form w-100">
            <div class="mb-3"><h2>Settings</h2></div>
            <div class="mb-2 w-100">
                <label for="day-start" class="form-label">Day Start</label>
                <VueDatePicker time-picker :is-24="false" id="day-start" v-model="startDate"></VueDatePicker>
            </div>
            <div class="mb-2">
                <label for="day-end" class="form-label">Day End</label>
                <VueDatePicker time-picker :is-24="false" id="day-end" v-model="endDate"></VueDatePicker>
            </div>
            <div class="mb-2">
                <label for="prompt-cycle" class="form-label">Prompt Cycle</label>
                <input type="number" id="prompt-cycle" class="form-control form-control-sm" v-model="promptCycle" min="10"/>
            </div>
            <div class="mb-2">
                <label for="sync-cycle" class="form-label">Sync Cycle To</label>
                <select id="sync-cycle" class="form-control form-control-sm" v-model="syncCycle">
                    <option value="NON">No Sync</option>
                    <option value="TOH">Top of the Hour</option>
                    <option value="BOH">Bottom of the Hour</option>
                    <option value="QTH">Quarter Hour</option>
                    <option value="HFH">Half Hour</option>
                </select>
            </div>
            <div class="d-flex flex-row-reverse">
                <button type="button" class="btn btn-sm btn-outline-secondary" :disabled="!changes" @click="saveSettings">Save</button>
            </div>
        </form>
   </div> 
</template>

<script setup lang="ts">
import { ref, watch, onMounted } from 'vue'
import { useStore } from '../store';
import { router } from '../router';
import VueDatePicker from "@vuepic/vue-datepicker"
import "@vuepic/vue-datepicker/dist/main.css"
import dayjs from 'dayjs';
import { Time } from '../domain';

const store = useStore()

const startDate = ref<Time>()
const endDate = ref<Time>()
const promptCycle = ref<number>()
const syncCycle = ref<string>()
const changes = ref<boolean>(false)

const saveSettings = () => {
    store.settings.DayEndsAt = endDate.value?.toString()
    store.settings.DayStartsAt = startDate.value?.toString()
    store.settings.PromptCycle = promptCycle.value 
    store.settings.SyncCycleTo = syncCycle.value
    store.saveSettings().then(() => {
        router.push("/")        
    })
}

onMounted(() => {
    startDate.value = Time.fromTimeString(store.settings.DayStartsAt) 
    endDate.value = Time.fromTimeString(store.settings.DayEndsAt) 
    promptCycle.value = store.settings.PromptCycle
    syncCycle.value = store.settings.SyncCycleTo

    watch([startDate, endDate, promptCycle, syncCycle], () => {
        changes.value = true 
    })
})


</script>

<style scoped>
</style>