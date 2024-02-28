package task

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string         `json:"name"`
	Content   string         `json:"content"`
	Status    string         `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
