<template>
<div class="col">
        <ModalFrame ref="modal">
            <div class="col">
                <form class="form">
                    <div class="row"><h2>New Project</h2></div>
                    <div class="row mb-2">
                        <label for="category" class="form-label">Category</label>
                        <select id="category" class="form-control form-control-sm" v-model="modalCategory">
                            <option v-for="category in store.categories" :key="category.ID" :value="category">
                                {{ category.Name }}
                            </option>
                        </select>
                    </div>
                    <div class="row mb-2">
                        <label for="project-name" class="form-label">Name</label>
                        <input type="text" id="project-name" class="form-control form-control-sm ps-1" v-model="modalValue">
                    </div>
                    <div class="row d-flex flex-row-reverse">
                        <button type="button" class="btn btn-sm btn-primary" :disabled="!allowCreateNew" @click="saveProject">SAVE</button>
                    </div>
                </form>
            </div>
        </ModalFrame>
        <div class="project-list column">
            <div class="content-header">
                <PlusCircleIcon :height="24" :width="24" @clicked="newProject"></PlusCircleIcon>
                <h2 class="d-flex flex-grow-1 align-left">Projects</h2>
            </div>
            <div class="row content-body">
                <table class="table">
                    <!-- <thead>
                        <tr>
                            <th scope="col">#</th>
                            <th scope="col">Project Name</th>
                        </tr>
                    </thead> -->
                    <tbody>
                        <tr v-for="project in store.projects" :key="project.ID">
                            <td>
                                <div class="col pe-1">
                                    <div class="row project-category">Id</div>
                                    <div class="row project-name">{{  formatID(project.ID) }}</div>
                                </div>
                            </td>
                            <td>
                                <div class="col">
                                    <div class="row project-category">{{  project.Category.Name }}</div>
                                    <div class="row project-name">{{  project.Name }}</div>
                                </div>
                            </td>
                            <td class="button-column">
                                <DeleteIcon class="delete-button" :height="16" :width="16" @clicked="deleteProject(project.ID)"></DeleteIcon>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div> 
    </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { useStore } from '../store';
import { router } from '../router';
import PlusCircleIcon from '../icons/PlusCircleIcon.vue';
import DeleteIcon from '@/icons/DeleteIcon.vue';
import ModalFrame from './ModalFrame.vue';
import { Category } from "../../bindings/wayd/services/database"

const store = useStore()
const modalValue = ref("")
const modalCategory = ref<Category>()
const modal = ref()

const newProject = () => {
    modalValue.value = ""
    modalCategory.value = undefined
    modal.value.show()
}
const saveProject = () => {
    store.newProject(modalValue.value, modalCategory.value).then(() => {
        modal.value.hide() 
    })
}

const deleteProject = (projectId: number) => {
    store.deleteProject(projectId) 
}

const allowCreateNew = computed(() => {
    return modalValue.value.length > 0 && modalCategory.value
})

const formatID = (id: number) => {
    return id.toString().padStart(3, "0") 
}

</script>

<style scoped>
.project-list {
    position: relative;
    padding: 5px;
}
.project {
    position: relative;
    margin-bottom: 0.5em;
}
.project-id {
    display: flex;
    flex-grow: 0;
    margin-right: 10px;
    text-align: right;
    color: #AAA;
}
.project-category {
    font-size: 0.6em;
    padding-left: 5px;
    color: #444;
}
.project-name {
    display: flex;
    flex-grow: 1;
    padding-left: 5px;
    font-weight: bold;
}
.project-menu {
    border: 1px solid red;
    margin-bottom: 10px;
    justify-content: end;
}
.label {
    font-size: 0.8em;
    color: #555;
}
.input {
    margin-bottom: 10px;
}
.button-column {
    position: relative;
    /* border: 1px solid red; */
}
.delete-button {
    margin: 0;
    padding: 0;
    /* border: 1px solid red; */
    position: absolute;
    right: 8px;
    bottom: 15px;
}
</style>