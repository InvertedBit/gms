package models

import "github.com/google/uuid"

type User struct {
	Model
	Email             string     `gorm:"type:string"`
	EncryptedPassword string     `gorm:"type:string"`
	RoleID            *uuid.UUID `gorm:"type:uuid"`
	Role              *Role      `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (u User) TableName() string {
	return "auth.users"
}

type Editor struct {
	Model
	UserName string `gorm:"uniqueIndex;type:string;size:256;not null"`
	User     User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
