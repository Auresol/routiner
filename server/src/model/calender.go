package model

import (
	"time"

	"gorm.io/gorm"
)

type Calender struct {
	Date          time.Time `json:"date" gorm:"primaryKey"`
	TaskIsCreated bool
	TotalTasks    int
	FinishedTasks int
	Status        WorkCompleteStatus `json:"status"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type WorkCompleteStatus int

/* >100% 100% 80% 50% 0% */
const (
	EXTRA WorkCompleteStatus = iota
	COMPLETE
	NEARLY
	SOME
	NONE
	EMPTY
)
