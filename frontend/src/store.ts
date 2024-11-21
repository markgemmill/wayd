import { ref, computed } from "vue"
import { get } from "lodash"
import type { Ref } from "vue"
import dayjs from "dayjs"
import type { Dayjs } from "dayjs"
// import weekday from "dayjs/plugin/weekday"
import { defineStore } from "pinia"
import { DatabaseService, Project, Category, Entry, ProjectDuration } from "../bindings/wayd/services/database" 
import { Settings } from "../bindings/wayd/services/models" 
import * as SettingsService from "../bindings/wayd/services/settingsservice"
import * as log from "../bindings/wayd/services/loggerservice"
import { DateRange } from "./domain"

// dayjs.extend(weekday)

export const useStore = defineStore("", () => {
    const projects = ref<Project[]>(new Array()) 
    const categories = ref<Category[]>(new Array())
    const currentEntry = ref<Entry>(Entry.createFrom({}))
    const currentSummaryStartDate = ref<Dayjs>()
    const currentSummaryEndDate = ref<Dayjs>()
    const currentSummaryResults = ref<ProjectDuration[]>(new Array())
    const settings = ref<Settings>(Settings.createFrom({}))

    const loadSettings = (): Promise<void> => {
       return SettingsService.GetSettings().then((value: Settings) => {
            settings.value = value
            return Promise.resolve()
       }) 
    }

    const saveSettings = (): Promise<void> => {
        return SettingsService.SetSettings(settings.value).then(() => {
            return Promise.resolve() 
        }) 
    }

    const loadProjects = (): Promise<void> => {
        log.Debug("Loading projects")
        return DatabaseService.GetAllActiveProjects().then((value: Project[]) => {
            value.forEach((project: Project) => {
                log.Debug(`Project: ${project.ID} :: ${project.Name}`)
                projects.value.push(project) 
            })
            return Promise.resolve()
        }) 
    }

    const loadCategories = (): Promise<void> => {
        log.Debug("Loading categories...") 
        return DatabaseService.GetAllActiveCategories().then((data: Category[]) => {
            data.forEach((category: Category) => {
                categories.value.push(category) 
            }) 
        })
    }

    const loadProjectSummary = (obj: object): Promise<void> => {

        const option = get(obj, "selection", "TODAY")
        const groupBy = get(obj, "groupBy", "PROJECT")

        let dateRange = new DateRange({})

        if (option === "DATE RANGE") {
            dateRange =  new DateRange({
                startDate: dayjs(get(obj, "startDate")),
                endDate: dayjs(get(obj, "endDate")),
            })
        } else {
            dateRange = DateRange.rangeAs(option)
        }

        log.Debug(dateRange.startDate.toString())
        log.Debug(dateRange.endDate.toString())

        currentSummaryStartDate.value = dateRange.startDate 
        currentSummaryEndDate.value = dateRange.endDate
        
        return DatabaseService.ProjectDurationTimes(
            groupBy,
            currentSummaryStartDate.value?.toDate(),
            currentSummaryEndDate.value?.toDate()
        ).then((results: ProjectDuration[]) => {
            currentSummaryResults.value.length = 0
            results.forEach((pd: ProjectDuration) => {
                log.Debug(`!!!!!!! ${pd.Name}`)
                currentSummaryResults.value.push(pd)
            })
            return Promise.resolve()
        })
    }

    const loadCurrentEntry = (): Promise<void> => {
        // do this at start up to restart any entry that was
        // left running.
        return DatabaseService.GetActiveEntry().then((entry: Entry) => {
            if (entry.ID !== undefined && entry.ID !== null && entry.ID > 0) {
                currentEntry.value = entry
            }
            return Promise.resolve()
        })
    }

    const updateCurrentEntry = (comment: string): Promise<void> => {
        currentEntry.value.Note = comment
        return DatabaseService.SaveEntry(currentEntry.value).then(() => {
            return Promise.resolve()
        })
    }

    const stopCurrentEntry = (): Promise<void> => {
       return DatabaseService.StopEntry(currentEntry.value).then(() => {
            currentEntry.value = Entry.createFrom({}) 
            return Promise.resolve()
       })
    }

    const newEntry = (project: Project): Promise<void> => {
        log.Debug(`Store.newEntry -> ${project.ID}`)
        return DatabaseService.NewEntry(project).then((entry: Entry) => {
            currentEntry.value = entry
            return Promise.resolve()
        })
    }

    const newProject = (name: string, category: Category): Promise<void> => {
        return DatabaseService.NewProject(name, category).then((project: Project) => {
            projects.value.push(project)
        })
    }

    const deleteProject = (projectId: number): Promise<void> => {
        return DatabaseService.DeleteProject(projectId).then(() => {
            let idToDelete = -1
            for (let i = 0; i < projects.value.length; i++) {
                if (projects.value[i].ID === projectId) {
                    idToDelete = i
                    projects.value.splice(i, 1)
                    break
                }
            }
            return 
        }).catch((err) => {
            log.Error(`${err}`)
        }) 
    }

    const newCategory = (name: string): Promise<void> => {
        return DatabaseService.NewCategory(name).then((category: Category) => {
            categories.value.push(category) 
        })
    }

    const currentEntryIsZero = computed<boolean>(() => {
        return currentEntry.value.ID === undefined || currentEntry.value.ID === null || currentEntry.value.ID === 0 
    })

    return {
        settings,
        projects,
        categories,
        currentEntry,
        currentEntryIsZero,
        currentSummaryResults,
        currentSummaryStartDate,
        currentSummaryEndDate,
        loadSettings,
        loadProjects,
        loadCategories,
        loadCurrentEntry,
        loadProjectSummary,
        saveSettings,
        updateCurrentEntry,
        stopCurrentEntry,
        newEntry,
        newProject,
        newCategory,
        deleteProject,
    }
    
})