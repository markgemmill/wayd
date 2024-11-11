import { createRouter, createMemoryHistory, createWebHistory, createWebHashHistory, type RouteLocationNormalizedGeneric, type LocationQuery } from "vue-router";
import NewTask from "./components/NewTask.vue";
import EntryStatus from "./components/EntryStatus.vue";
import WelcomePane from "./components/WelcomePane.vue";
import ProjectList from "./components/ProjectList.vue";
import CategoryList from "./components/CategoryList.vue";
import ChangePrompt from "./components/ChangePrompt.vue";
import SettingsForm from "./components/SettingsForm.vue";
import ReportingPane from "./components/ReportingPane.vue";
import * as log from "../bindings/wayd/services/loggerservice"

const homeNavigation = (to: RouteLocationNormalizedGeneric, from: RouteLocationNormalizedGeneric) => {
    log.Debug(`home navigation from: ${from.fullPath}`)
    log.Debug(`home navigation to: ${to.fullPath}`)
    return true 
}

const routes = [
    { path: "/", name: "home", component: WelcomePane, beforeEnter: homeNavigation },
    { path: "/new", name: "new", component: NewTask},
    { path: "/current", name: "current", component: EntryStatus },
    { path: "/projects", name: "projects", component: ProjectList },
    { path: "/categories", name: "categories", component: CategoryList },
    { path: "/settings", name: "settings", component: SettingsForm },
    { path: "/reporting", name: "reporting", component: ReportingPane},
    { path: "/prompt", name: "prompt", component: ChangePrompt },
] 

export const router = createRouter({
    history: createWebHashHistory(),
    routes
})

router.beforeEach((to: RouteLocationNormalizedGeneric, from: RouteLocationNormalizedGeneric) => {
    log.Debug(`Navigation -> ${from.fullPath} -> ${to.fullPath}`)
    // const qry: LocationQuery = to.query
    // if (qry["prompt"]) {
    //     return { name: "prompt"}
    // }
    return true 
})