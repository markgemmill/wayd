<template>
  <div class="column">
    <div class="form">
        <div class="mb-3"><h2>Current Task</h2></div>
        <div class="mb-3">
            <CurrentTask
                :start="store.currentEntry.Start"
                :name="store.currentEntry.Project.Name"
                :comments="store.currentEntry.Note"
                @comment-changed="onCommentChange"
            ></CurrentTask>
        </div>
        <div class="mb-2">
            <button class="btn btn-lg btn-outline-primary" @click="newEntry">Start something else.</button>
        </div>
        <div class="">
            <button class="btn btn-lg btn-outline-primary" @click="stopEntry">It's the end of the day.</button>
        </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import CurrentTask from './CurrentTask.vue';
import { useStore } from '../store';
import { router } from '../router';
import * as log from "../../bindings/github.com/markgemmill/wayd/services/logger"

const store = useStore()

const commentChanges = new Array<string>() 

const onCommentChange = (comment: string) => {
    commentChanges.push(comment)
}

const continueEntry = () => {
    const value = commentChanges.pop()
    log.Debug(`updated comments: ${value}`)
    if (value) {
        store.updateCurrentEntry(value).finally(() => {
            commentChanges.length = 0
        })
    }
}

const updateCurrentEntry = (routePath: string) => {
    let promise = Promise.resolve() 

    const value = commentChanges.pop()

    log.Debug(`updated comments: ${value}`)

    if (value) {
        promise = store.updateCurrentEntry(value)
    } 

    promise.finally(() => {
        commentChanges.length = 0
        store.stopCurrentEntry().then(() => {
            router.push(routePath)
        })
    })
} 

const newEntry = () => {
    updateCurrentEntry("/new")
}

const stopEntry = () => {
    updateCurrentEntry("/")
}

</script>

<style scoped>
/* button {
  border: 1px solid #CCC;
  border-radius: 3px;
  width: 70%;
  padding: auto;
  background-color: white;
  text-align: center;
  align-content: center;
  margin-bottom: 1em;
} */
</style>