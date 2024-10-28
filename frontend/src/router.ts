import { createRouter, createMemoryHistory } from "vue-router";
import NewTask from "./components/NewTask.vue";
import EntryStatus from "./components/EntryStatus.vue";
import WelcomePane from "./components/WelcomePane.vue";
import ProjectList from "./components/ProjectList.vue";
import CategoryList from "./components/CategoryList.vue";
import ChangePrompt from "./components/ChangePrompt.vue";
import SettingsForm from "./components/SettingsForm.vue";
import ReportingPane from "./components/ReportingPane.vue";
import * as log from "../bindings/github.com/markgemmill/wayd/services/logger"

const routes = [
    { path: "/", name: "home", component: WelcomePane },
    { path: "/new", name: "new", component: NewTask},
    { path: "/current", name: "current", component: EntryStatus },
    { path: "/projects", name: "projects", component: ProjectList },
    { path: "/categories", name: "categories", component: CategoryList },
    { path: "/settings", name: "settings", component: SettingsForm },
    { path: "/reporting", name: "reporting", component: ReportingPane},
    { path: "/prompt", name: "prompt", component: ChangePrompt },
] 

export const router = createRouter({
    history: createMemoryHistory(),
    routes
})

router.beforeEach((to, from) => {
    log.Debug(`Navigation -> ${from.fullPath} -> ${to.fullPath}`)
    return true 
})