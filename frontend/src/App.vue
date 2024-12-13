<template>
  <div id="application" class="column">
    <div v-if="mainWindow" id="header" class="d-flex flex-row mb-3">
      <TimeClock :format="'icon'" :start-time="currentTimer" @clicked="navigateEntry"></TimeClock>
      <div class="d-flex flex-grow-1"></div>
      <HomeIcon class="me-2" :width="20" :height="20" @click="navigateTo('')"></HomeIcon>
      <DockIcon class="me-2" :width="18" :height="18" :dock-position="store.settings.DockPosition" @click="dockWindow"></DockIcon>
      <ReportFileIcon class="me-2" :width="20" :height="20" @click="navigateTo('reporting')"></ReportFileIcon>
      <InboxIcon class="me-2" :width="20" :height="20" @click="navigateTo('categories')"></InboxIcon>
      <InboxesIcon class="me-2" :width="20" :height="20" @click="navigateTo('projects')"></InboxesIcon>
      <GearIcon class="" :width="20" :height="20" @click="navigateTo('settings')"></GearIcon>
    </div>
    <div v-if="initialLoad == 15" id="body" class="row">
      <RouterView></RouterView>
    </div>
    <div v-else id="body" class="row">
      LOADING...
    </div>
    <div id="footer" class="row"></div>
  </div>
</template>

<script setup lang="ts">
import { ref, onBeforeMount } from "vue";
import { useStore } from "./store"
import { RouterView } from "vue-router";
import HomeIcon from "./icons/HomeIcon.vue";
import TimeClock from "./components/TimeClock.vue";
import InboxIcon from "./icons/InboxIcon.vue";
import InboxesIcon from "./icons/InboxesIcon.vue";
import GearIcon from "./icons/GearIcon.vue";
import DockIcon from "./icons/DockIcon.vue";
import ReportFileIcon from "./icons/ReportFileIcon.vue";
import { router } from "./router";
import dayjs from "dayjs"
import type { Dayjs } from "dayjs"
import * as log from "../bindings/wayd/services/loggerservice"
import {Events} from "@wailsio/runtime"

const store = useStore()
const initialLoad = ref(0)
const currentTimer = ref<Dayjs|null>()
const mainWindow = ref(true)

const navigateTo = (location: string) => {
    router.push(`/${location}`)
}

const navigateEntry = () => {
    if (store.currentEntryIsZero) {
        router.push("/new") 
    } else {
        router.push("/current")
    }
}
 
store.$subscribe((mutation, state) => {
    log.Debug("Hey the store state changed!")
    if (!store.currentEntryIsZero) {
      currentTimer.value = dayjs(store.currentEntry.Start) 
    } else {
      currentTimer.value = null
    }
})

const dockWindow = () => {
    const event = new Events.WailsEvent("dock-window", { Position: store.settings.DockPosition})
    Events.Emit(event)
}

onBeforeMount(() => {

    log.Debug(`>>>> START URL >>>> ${window.location.href}`)

    store.loadProjects().then(() => {
        initialLoad.value += 1 
    }) 

    store.loadCurrentEntry().then(() => {
        initialLoad.value += 2 
    })

    store.loadCategories().then(() => {
        initialLoad.value += 4
    })

    store.loadSettings().then(() => {
        initialLoad.value += 8        
    })

    Events.On("display-prompt", () => {
        log.Debug("frontend display prompt")        
        router.push("/prompt")
    })

})
</script>

<style scoped>
</style>
