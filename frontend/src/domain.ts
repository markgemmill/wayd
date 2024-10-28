import dayjs from 'dayjs'
import type { Dayjs } from 'dayjs'
import { get, isNil } from "lodash"


export class TaskEntry {
    startTime: Dayjs 
    endTime: Dayjs
    project: string
    constructor (startTime: Dayjs, project: string) {
        this.startTime = startTime 
        this.project = project
    }

    stop() {
        this.endTime = dayjs()
    }
}

export class Time {
    hours: number 
    minutes: number

    constructor(hours: number, minutes: number) {
        this.hours = hours
        this.minutes = minutes
    }

    toString (): string {
        const hrs = this.hours.toString().padStart(2, "0")
        const min = this.minutes.toString().padStart(2, "0")
        return `${hrs}:${min}:00`
    }

    static fromTimeString(timeString: string): Time {
        const d = dayjs(`1900-01-01 ${timeString}`, "HH:MM:SS")
        return new Time(d.hour(), d.minute())
    }
}

export const absoluteStartDate = (dte: Dayjs): Dayjs => {
    dte = dte.set("hours", 0)
    dte = dte.set("minutes", 0)
    dte = dte.set("seconds", 0)
    return dte
}

export const absoluteEndDate = (dte: Dayjs): Dayjs => {
    dte = dte.set("hours", 23)
    dte = dte.set("minutes", 59)
    dte = dte.set("seconds", 59)
    return dte
}

export class DateRange {
    startDate: Dayjs
    endDate: Dayjs
    constructor (params: object) {
        let start = get(params, "startDate")
        let end = get(params, "endDate")
        if (isNil(start)) {
            start = dayjs() 
            end = dayjs() 
        }
        this.startDate = absoluteStartDate(start)
        this.endDate = absoluteEndDate(end)
    }

    static today(): DateRange {
        return new DateRange({})
    } 

    static yesterday(): DateRange {
        const yd = dayjs().subtract(1, "day")
        return new DateRange({
            startDate: yd,
            endDate: yd,
        })
    }

    static currentWeek(): DateRange {
        const startDate = dayjs().weekday(0)
        const endDate = dayjs().weekday(6)
        return new DateRange({
            startDate: startDate,
            endDate: endDate,
        })
    }

    static lastWeek(): DateRange {
        const weekAgo = dayjs().subtract(7, "days")
        const startDate = weekAgo.weekday(0)
        const endDate = weekAgo.weekday(6)
        return new DateRange({
            startDate: startDate,
            endDate: endDate,
        })
    }

    static rangeAs(rangeName: string): DateRange {
        switch (rangeName) {
            case "YESTERDAY":
                return DateRange.yesterday()
            case "CURRENT WEEK":
                return DateRange.currentWeek()
            case "LAST WEEK":
                return DateRange.lastWeek()
        }
        return DateRange.today()
    }
}
