<template>
  <div class="column">
    <form class="form">
        <div class="mb-2"><h2>New Task</h2></div>
        <div class="mb-3">
            <label class="form-label">Select Project</label>
            <select class="form-control form-control-sm" v-model="selectedProject">
                <option v-for="project in store.projects" :key="project.ID" :value="project">{{ project.Name }}</option>
            </select> 
        </div>
        <div class="d-flex flex-row-reverse">
            <button type="button" class="btn btn-sm btn-primary" @click="startNewEntry">Start</button>
        </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { useStore } from '../store';
import { router } from '../router';
import { Project } from '../../bindings/wayd/services/database';

const store = useStore()
const selectedProject = ref<Project>()

const startNewEntry = () => {
    store.newEntry(selectedProject.value).then(() => {
        router.push("/current") 
    })
}

</script>

<style scoped>
</style>