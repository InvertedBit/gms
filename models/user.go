package models

type User struct {
	Model
	Email             string `gorm:"type:string"`
	EncryptedPassword string `gorm:"type:string"`
	RoleSlug          string `gorm:"type:string;index"`
	Role              Role   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:RoleSlug;references:Slug"`
}

func (u User) TableName() string {
	return "public.users"
}

// type Editor struct {
// 	Model
// 	UserName string `gorm:"uniqueIndex;type:string;size:256;not null"`
// 	User     User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
// }
