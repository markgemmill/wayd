<template>
    <div v-if="viewReady" class="column">
        <ModalFrame ref="modal">
            <div class="row w-100">
                <form class="w-100">
                    <div><h2>Summary Parameters</h2></div>
                    <div class="mb-2">
                        <label for="group-by" class="form-label">Group By</label>
                        <select name="group-by" id="group-by" class="form-control form-control-sm" v-model="groupBy">
                            <option value="CATEGORY">CATEGORY</option>
                            <option value="PROJECT">PROJECT</option>
                        </select>
                    </div>
                    <div class="mb-2">
                        <label class="form-label" for="selection-type">Type</label>
                        <select id="selection-type" class="form-control form-control-sm" v-model="dateSelection">
                            <option value="TODAY" selected>TODAY</option>
                            <option value="YESTERDAY">YESTERDAY</option>
                            <option value="CURRENT WEEK">CURRENT WEEK</option>
                            <option value="LAST WEEK">LAST WEEK</option>
                            <option value="DATE RANGE">DATE RANGE</option>
                        </select>
                    </div>
                    <div v-if="dateSelection == 'DATE RANGE'" class="mb-1">
                        <label for="start-date" class="form-label">Start Date</label>
                        <VueDatePicker 
                            :time-picker="false" 
                            :auto-apply="true" 
                            :format="'MM/dd/yyyy'"
                            id="start-date" 
                            v-model="startDate"></VueDatePicker>
                    </div>
                    <div v-if="dateSelection == 'DATE RANGE'" class="mb-2">
                        <label for="end-date" class="form-label">End Date</label>
                        <VueDatePicker 
                            :time-picker="false" 
                            :auto-apply="true" 
                            :format="'MM/dd/yyyy'"
                            id="end-date" 
                            v-model="endDate"></VueDatePicker>
                    </div>
                    <div class="d-flex flex-row-reverse">
                        <button type="button" class="btn btn-sm btn-outline-secondary" @click="updateSummary">Select</button>
                    </div>
                </form>
            </div>
        </ModalFrame>
        <div class="d-flex flex-row-reverse mb-2">
            <button type="button" class="btn btn-sm btn-outline-secondary mb-auto" @click="modal.show()">...</button>
            <div class="col">
                <div class="me-2 d-flex flex-grow-1 fw-bold fs-6">{{ dateSelection }}</div>
                <div class="fw-light fst-italic fs-6">{{ selectedDates }}</div> 
            </div>
        </div>
        <table class="table">
            <thead>
                <tr>
                    <th scope="col">#</th>
                    <th scope="col">{{ capitalize(groupBy) }}</th>
                    <th scope="col">Cnt</th>
                    <th scope="col">Duration</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="project, index in store.currentSummaryResults" :key="index">
                    <td>{{ formatID(index) }}</td>
                    <td>{{ project.Name }}</td>
                    <td>{{ project.EntryCount }}</td>
                    <td>{{ project.Duration }}</td>
                </tr>
            </tbody>
        </table>
    </div> 
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useStore } from '../store';
import ModalFrame from './ModalFrame.vue';
import VueDatePicker from "@vuepic/vue-datepicker"
import { Time } from '../domain';
import { capitalize } from "lodash"

const modal = ref()
const store = useStore()
const viewReady = ref(false)
const dateSelection = ref("TODAY")
const startDate = ref<Date>()
const endDate = ref<Date>()
const groupBy = ref("PROJECT")

const formatID = (id: number) => {
    return id.toString().padStart(3, "0") 
}

const updateSummary = () => {
    viewReady.value = false
    const args = {
        groupBy: groupBy.value,
        selection: dateSelection.value,
        startDate: startDate.value,
        endDate: endDate.value,
    }
    store.loadProjectSummary(args).then(() => {
        viewReady.value = true
    })
}

const selectedDates = computed<string>((): string => {
    const sd = store.currentSummaryStartDate.format("YYYY-MM-DD")
    const ed = store.currentSummaryEndDate.format("YYYY-MM-DD")
    if (sd === ed) {
        return sd
    }
    return `${sd} - ${ed}`
})


onMounted(() => {
    store.loadProjectSummary({selection: "TODAY"}).then(() => {
        viewReady.value = true
    })
})

</script>

<style scoped>
.project-list {
    padding: 5px;
    width: 100%;
    font-family: monospace;
    font-weight: normal;
}
.project-row {
    position: relative;
    margin-bottom: 0.1em;
    width: 100%;
    /* border: 1px solid red; */
}
.project-id {
    display: flex;
    flex-grow: 0;
    margin-right: 10px;
    text-align: right;
    color: #AAA;
}
.project-name {
    display: flex;
    flex-grow: 1;
    padding-left: 5px;
    font-weight: bold;
    text-align: left;
}
.project-time {
    display: flex;
    flex-grow: 0;
    padding-left: 5px;
    font-weight: bold;
    text-align: right;
}
</style>