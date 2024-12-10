package model

import (
	"time"
)

type Calender struct {
	Date          time.Time      `json:"date" gorm:"primaryKey"`
	FinishedRatio float32        `json:"finished_ratio" `
	Status        CalenderStatus `json:"status"`
}

type CalenderStatus int

/* >100% 100% 80% 50% 0% */
const (
	EXTRA CalenderStatus = iota
	COMPLETE
	NEARLY
	SOME
	NONE
	EMPTY
)

func CalenderStatusUpdate(calender *Calender, finish int, total int) {
	if total == 0 {
		calender.FinishedRatio = 0
		calender.Status = EMPTY
		return
	}

	percent := float32(finish) / float32(total)
	calender.FinishedRatio = percent

	if percent > 1 {
		calender.Status = EXTRA
	} else if percent == 1 {
		calender.Status = COMPLETE
	} else if percent >= 0.7 {
		calender.Status = NEARLY
	} else if percent > 0 {
		calender.Status = SOME
	} else {
		calender.Status = NONE
	}

}
