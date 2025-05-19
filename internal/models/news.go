package models

import (
	"time"

	"github.com/google/uuid"
)

type NewsMeta struct {
	ID             uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title          string    `gorm:"not null"`
	TitleHash      []byte    `gorm:"not null"`
	URL            string    `gorm:"not null"`
	Source         string    `gorm:"not null"`
	PublishedAt    *time.Time
	Lang           string    `gorm:"size:2"`
	Embedding      []float32 `gorm:"type:float8[]"` // временно
	Status         string    `gorm:"default:NEW;not null"`
	OriginalNewsID *uuid.UUID
	CreatedAt      time.Time `gorm:"default:now()"`
	UpdatedAt      time.Time `gorm:"default:now()"`
}

type NewsText struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	NewsID    uuid.UUID `gorm:"type:uuid;not null"`
	RawText   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:now()"`
}
