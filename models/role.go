package models

type Role struct {
	Model
	Slug        string `gorm:"uniqueIndex;type:string;size:50;not null"`
	Name        string `gorm:"type:string;size:255;not null"`
	Description string `gorm:"type:string;size:255"`
}

func (r Role) TableName() string {
	return "public.roles"
}
