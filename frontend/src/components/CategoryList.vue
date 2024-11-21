<template>
    <div class="column">
        <ModalFrame ref="modal">
            <div class="row w-100">
                <form class="w-100">
                    <div><h2>New Category</h2></div>
                    <div>
                        <div class="form-label">Name</div>
                        <input type="text" class="form-control form-control-sm mb-2" v-model="modalValue">
                    </div>
                    <div class="d-flex flex-row-reverse">
                        <button type="button" class="btn btn-sm btn-primary":disabled="modalValue.length == 0" @click="saveProject">SAVE</button>
                    </div>
                </form>
            </div>
        </ModalFrame>
        <div class="project-list column">
            <div class="content-header">
                <PlusCircleIcon :height="24" :width="24" @clicked="newCategory"></PlusCircleIcon>
                <h2 class="d-flex flex-grow-1 align-left">Categories</h2>
            </div>
            <div class="row content-body">
                <table class="table">
                    <tbody>
                        <tr v-for="category in store.categories" :key="category.ID">
                            <td>
                                <div class="col pe-2">
                                    <div class="row project-category">Id</div>
                                    <div class="row project-name">{{  formatID(category.ID) }}</div>
                                </div>
                            </td>
                            <td>
                                <div class="col">
                                    <div class="row project-category">Name</div>
                                    <div class="row project-name">{{  category.Name }}</div>
                                </div>
                            </td>
                            <td class="button-column">
                                <DeleteIcon class="delete-button" :height="16" :width="16" @clicked="deleteCategory(category.ID)"></DeleteIcon>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
        <!-- </div>  -->
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useStore } from '../store';
import PlusCircleIcon from '../icons/PlusCircleIcon.vue';
import DeleteIcon from '@/icons/DeleteIcon.vue';
import ModalFrame from './ModalFrame.vue';

const store = useStore()
const modalValue = ref("")
const modal = ref()

const newCategory = () => {
    modalValue.value = ""
    modal.value.show()
}
const saveProject = () => {
    store.newCategory(modalValue.value).then(() => {
        modal.value.hide() 
    })
}

const deleteCategory = (categoryId: number) => {
    
}

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