package model

import (
	"time"

	"gorm.io/gorm"
)

type Log struct {
	ID     uint      `json:"id" gorm:"primaryKey;autoIncrement;->"`
	Date   time.Time `json:"date" gorm:"index"`
	TaskID uint      `json:"task_id"`
	Detail string    `json:"detail"`

	CreatedAt time.Time      `json:"create_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
