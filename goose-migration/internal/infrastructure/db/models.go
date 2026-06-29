package db

import (
	"time"

	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type Note struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title     string    `gorm:"size:50;not null"`
	Content   string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
