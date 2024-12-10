package util

import (
	"time"
)

func GetTodayBegin() time.Time {
	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)
}

func GetDateBegin(timeIn *time.Time) time.Time {
	return time.Date(timeIn.Year(), timeIn.Month(), timeIn.Day(), 0, 0, 0, 0, time.Local)
}
