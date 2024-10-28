<template>
    <div class="column">
        <div class="row">
        </div>
        <div class="row mb-3">
            <button type="button" class="btn btn-lg btn-outline-primary" @click="startWithEntry">Start New Timer</button>
        </div>
        <div class="row">
            <ProjectSummary></ProjectSummary>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { useStore } from '../store';
import { router } from "../router"
import * as log from "../../bindings/github.com/markgemmill/wayd/services/logger"
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
</style>