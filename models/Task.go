package models

import (
	"time"
)

type Task struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Titulo      string    `json:"titulo" gorm:"not null;unique"`
	Descripcion string    `json:"descripcion"`
	Done        bool      `json:"done" gorm:"default:false"`
	UserId      uint      `json:"user_id"`
	User        *User     `gorm:"foreignKey:UserId" json:"user,omitempty"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
