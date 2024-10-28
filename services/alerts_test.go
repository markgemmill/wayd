package services

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAlertSyncToHalfHour(t *testing.T) {
	type data struct {
		delay, hr, min, sec, xhr, xmin, xsec int
	}
	testData := []data{
		data{30, 2, 35, 22, 3, 0, 0},  //  0  2:35 -> 3:00 (25)
		data{30, 2, 40, 22, 3, 0, 0},  //  1  2:35 -> 3:00 (20)
		data{30, 2, 45, 22, 3, 30, 0}, //  2  2:45 -> 3:30 (45)
		data{30, 2, 55, 22, 3, 30, 0}, //  3  2:55 -> 3:30 (35)
		data{30, 3, 00, 22, 3, 30, 0}, //  4  2:55 -> 3:30 (30)
		data{30, 3, 05, 22, 3, 30, 0}, //  5  3:05 -> 3:30 (25)
		data{30, 3, 10, 22, 3, 30, 0}, //  6  3:10 -> 3:30 (20)
		data{30, 3, 15, 22, 4, 00, 0}, //  7  3:15 -> 4:00 (45)
		data{30, 3, 20, 22, 4, 00, 0}, //  8  3:20 -> 4:00 (40)
		data{30, 3, 25, 22, 4, 00, 0}, //  9  3:25 -> 4:00 (35)
		data{30, 3, 30, 22, 4, 00, 0}, // 10  3:30 -> 4:00 (30)
	}
	for index, d := range testData {
		tm := time.Date(2024, time.October, 30, d.hr, d.min, d.sec, 0, time.UTC)
		ctm := DelayCalculation(tm, d.delay, SYNC_TO_THE_HALF_HOUR)
		assert.Equal(t, ctm.Hour(), d.xhr, fmt.Sprintf("[test %d] hour %d, expected %d", index, ctm.Hour(), d.xhr))
		assert.Equal(t, ctm.Minute(), d.xmin, fmt.Sprintf("[test %d] minute %d, expectd %d", index, ctm.Minute(), d.xmin))
		assert.Equal(t, ctm.Second(), d.xsec, fmt.Sprintf("[test %d] seconds %d, expected %d", index, ctm.Second(), d.xsec))
	}
}

func TestAlertSyncToQuarterHour(t *testing.T) {
	type data struct {
		delay, hr, min, sec, xhr, xmin, xsec int
	}
	testData := []data{
		data{15, 2, 35, 22, 2, 45, 0}, //  0  2:35 -> 2:45 (10)
		data{15, 2, 40, 22, 3, 00, 0}, //  1  2:35 -> 3:00 (25)
		data{15, 2, 45, 22, 3, 00, 0}, //  2  2:45 -> 3:00 (15)
		data{15, 2, 50, 22, 3, 00, 0}, //  3  2:50 -> 3:00 (10)
		data{15, 2, 55, 22, 3, 15, 0}, //  4  2:55 -> 3:15 (20)
		data{15, 3, 00, 22, 3, 15, 0}, //  5  3:00 -> 3:15 (30)
		data{15, 3, 05, 22, 3, 15, 0}, //  6  3:05 -> 3:15 (25)
		data{15, 3, 10, 22, 3, 30, 0}, //  7  3:10 -> 3:30 (20)
		data{15, 3, 15, 22, 3, 30, 0}, //  8  3:15 -> 3:30 (45)
		data{15, 3, 20, 22, 3, 30, 0}, //  9  3:20 -> 3:30 (40)
		data{15, 3, 25, 22, 3, 45, 0}, // 10  3:25 -> 3:45 (35)
		data{15, 3, 30, 22, 3, 45, 0}, // 11  3:30 -> 3:45 (30)
	}
	for index, d := range testData {
		tm := time.Date(2024, time.October, 30, d.hr, d.min, d.sec, 0, time.UTC)
		ctm := DelayCalculation(tm, d.delay, SYNC_TO_THE_QUARTER_HOUR)
		assert.Equal(t, ctm.Hour(), d.xhr, fmt.Sprintf("[test %d] hour %d, expected %d", index, ctm.Hour(), d.xhr))
		assert.Equal(t, ctm.Minute(), d.xmin, fmt.Sprintf("[test %d] minute %d, expectd %d", index, ctm.Minute(), d.xmin))
		assert.Equal(t, ctm.Second(), d.xsec, fmt.Sprintf("[test %d] seconds %d, expected %d", index, ctm.Second(), d.xsec))
	}
}
