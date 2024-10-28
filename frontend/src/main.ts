import { createApp } from 'vue'
import App from './App.vue'
import { createPinia } from 'pinia'
import { router } from './router'
import dayjs from "dayjs"
import duration from "dayjs/plugin/duration"
import "bootstrap/dist/css/bootstrap.min.css"


dayjs.extend(duration)

const pinia = createPinia()
const app = createApp(App)

app.use(pinia)
app.use(router)
app.mount('#app')
