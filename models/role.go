package models

type Role struct {
	Model
	Name        string `gorm:"uniqueIndex;type:string;size:50;not null"`
	Description string `gorm:"type:string;size:255"`
}

func (r Role) TableName() string {
	return "auth.roles"
}
