package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Nombre    string    `json:"nombre" gorm:"not null"`
	Apellidos string    `json:"apellidos" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null;unique"`
	Tasks     []Task    `json:"tasks,omitempty" gorm:"foreignKey:UserId"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
