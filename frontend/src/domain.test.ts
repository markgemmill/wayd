import { expect, test } from "vitest"
import { DateRange, absoluteEndDate, absoluteStartDate } from "./domain.js"
import dayjs from "dayjs"


test("absoluteStartDate...", () => {
    const dte = absoluteStartDate(dayjs()) 

    expect(dte.hour()).toBe(0)
    expect(dte.minute()).toBe(0)
    expect(dte.second()).toBe(0)
})

test("absoluteEndDate...", () => {
    const dte = absoluteEndDate(dayjs()) 

    expect(dte.hour()).toBe(23)
    expect(dte.minute()).toBe(59)
    expect(dte.second()).toBe(59)
})

test("DateRange", () => {
    const rng = DateRange.today() 

    expect(rng.startDate.hour()).toBe(0)
    expect(rng.endDate.hour()).toBe(23)
})

