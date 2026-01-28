package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type Model struct {
	ID        uuid.UUID      `gorm:"primarykey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time      `gorm:"default:now()"`
	UpdatedAt time.Time      `gorm:"default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type EntryType int

const (
	EntryTypeActive EntryType = iota
	EntryTypeEdit
	EntryTypeHistory
)

type CoWModel struct {
	Model
	EntryType  EntryType `gorm:"type:int;not null;default:0;index"`
	PreviousID *string   `gorm:"type:uuid;index;nullable"`
}

func GenerateSlug(text string) string {
	return slug.Make(text)
}
