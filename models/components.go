package models

import "github.com/google/uuid"

type Page struct {
	CoWModel
	Slug       string  `gorm:"type:varchar(255);not null;uniqueIndex"`
	Title      string  `gorm:"type:varchar(255);not null;index"`
	LayoutSlug *string `gorm:"type:varchar(1023);nullable;index"`
	Layout     *Layout `gorm:"foreignKey:LayoutSlug;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type ComponentInstance struct {
	CoWModel
	Slug       string              `gorm:"type:varchar(1023);not null;uniqueIndex"`
	Name       string              `gorm:"type:varchar(255);not null;index"`
	ParentSlug *string             `gorm:"type:varchar(255);index;nullable"`
	Children   []ComponentInstance `gorm:"foreignKey:ParentSlug;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Layouts    []Layout            `gorm:"foreignKey:ComponentInstanceSlug;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type ComponentPropertyType int

const (
	Default ComponentPropertyType = iota
	LayoutOverride
	PageOverride
)

type ComponentProperty struct {
	CoWModel
	Slug                  string                `gorm:"type:varchar(1023);not null;uniqueIndex"`
	ComponentInstanceSlug string                `gorm:"type:varchar(1023);not null;index"`
	Key                   string                `gorm:"type:varchar(255);not null;index"`
	Value                 string                `gorm:"type:text;not null"`
	Type                  ComponentPropertyType `gorm:"type:int;not null;default:0;index"`
}

type ComponentMedia struct {
	CoWModel
	Slug                  string    `gorm:"type:varchar(1023);not null;uniqueIndex"`
	ComponentInstanceSlug string    `gorm:"type:varchar(1023);not null;index"`
	MediaID               uuid.UUID `gorm:"type:uuid;not null;index"`
	Media                 Media     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Layout struct {
	CoWModel
	Slug                  string             `gorm:"type:varchar(1023);not null;uniqueIndex"`
	Title                 string             `gorm:"type:varchar(1023);not null;index"`
	ComponentInstanceSlug *string            `gorm:"type:varchar(1023);nullable;index"`
	ComponentInstance     *ComponentInstance `gorm:"foreignKey:ComponentInstanceSlug;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Pages                 []Page             `gorm:"foreignKey:LayoutSlug;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
