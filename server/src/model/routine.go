package model

import (
	"time"

	"gorm.io/gorm"
)

type Routine struct {
	ID            uint   `json:"id" gorm:"primaryKey;autoIncrement;->"`
	Title         string `json:"title" gorm:"unique"`
	Explain       string `json:"explain"`
	IconCodePoint int    `json:"icon_code_point"`

	ActiveDate   time.Time   `json:"active_date"`
	DueIn        int         `json:"due_in"`
	ForceReset   bool        `json:"force_reset"`
	RoutineMode  RoutineMode `json:"routine_mode"`
	DayInWeekly  int8        `json:"day_in_weekly"`
	Frequency    int         `json:"frequency" gorm:"default:1"`
	ResetOnMonth bool        `json:"reset_on_month"`

	Task []Task `gorm:"foreignKey:RoutineID"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
type RoutineMode int

const (
	WEEKLY RoutineMode = iota + 1
	PERIOD
	TODO
)

/* Decodingthe number to a bit string */
func BitmaskDecoding(bit int8) [7]bool {
	var days [7]bool

	i := 0
	for bit > 0 {
		if bit&1 == 1 {
			days[i] = true
		}
		i++
		bit = bit >> 1
	}

	return days
}
