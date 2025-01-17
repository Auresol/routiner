package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;->"`
	RoutineID uint      `json:"routine_id"`
	Begin     time.Time `json:"begin"`
	Due       time.Time `json:"due"`
	Detail    string    `json:"detail"`
	Status    bool      `json:"status"`

	// Embedded struct to store routine data (preload)
	Routine *Routine `json:"routine" gorm:"foreignKey:RoutineID;preload:false"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
