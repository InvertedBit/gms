package models

type User struct {
	Model
	Email             string `gorm:"type:string"`
	EncryptedPassword string `gorm:"type:string"`
	Role              string `gorm:"type:string"`
}

func (u User) TableName() string {
	return "auth.users"
}

type Editor struct {
	Model
	UserName string `gorm:"uniqueIndex;type:string;size:256;not null"`
	User     User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
