package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID            uint   `json:"id" gorm:"primaryKey;autoIncrement;->"`
	Title         string `json:"title" gorm:"unique"`
	Explain       string `json:"explain"`
	IconCodePoint int    `json:"icon_code_point"`

	RoutineMode  RoutineMode `json:"routine_mode"`
	DayInWeekly  int8        `json:"day_in_weeky"`
	Frequency    int         `json:"frequency" gorm:"default:1"`
	ResetOnMonth bool        `json:"reset_on_month"`

	Log []Log `gorm:"foreignKey:TaskID"`

	CreatedAt time.Time      `json:"create_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type RoutineMode int

const (
	WEEKLY RoutineMode = iota + 1
	PERIOD
)

/* check if target is decoding in the bitmask */
func BitmaskDecoding(bit int8, target *int) bool {
	i := 0
	for bit > 0 {
		if (bit&1 == 1) && (*target == i) {
			return true
		}
		i++
		bit = bit >> 1
	}

	return false
}
