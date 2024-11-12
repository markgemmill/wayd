<template>
    <div class="column">
        <div class="row mt-5">
            <button type="button" class="start-button" @click="startWithEntry">
                New
            </button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { useStore } from '../store';
import { router } from "../router"
import * as log from "../../bindings/wayd/services/loggerservice"
import ProjectSummary from './ProjectSummary.vue';

const store = useStore()

const startWithEntry = () => {
    log.Debug(store.currentEntry.ID.toString())
    log.Debug("checking up on the currentEntry state...")
    if (store.currentEntryIsZero) {
        log.Debug("open new entry form...")
        router.push("/new") 
    } else {
        log.Debug("open current entry form...")
        router.push("/current")
    }
}

const viewProjects = () => {
    router.push("/projects") 
}

</script>

<style scoped>
.message {
    text-align: left;
    border: 1px solid black;
    padding: 10px;
    flex-grow: 1;
}
.titled {
    font-size: 1.5em;
    font-weight: bold;
}
.start-button {
    height: 240px;
    width: 240px;
    padding: 10px;
    margin: auto;
    border-radius: 50%;
    border: 1px solid #CCC;
    background-color: white;
    color: #CCC;
    font-size: 3em;
}
.start-button:hover {
    border-color: black;
    color: #444;
}
</style>