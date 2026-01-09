package auth

import (
	"context"

	"github.com/invertedbit/gms/database"
	"github.com/invertedbit/gms/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func AuthenticateUser(email, password string) (models.User, error) {
	user, err := gorm.G[models.User](database.DBConn).Where("email = ?", email).First(context.Background())
	if err != nil {
		return models.User{}, err
	}

	if err := VerifyPassword(user.EncryptedPassword, password); err != nil {
		return models.User{}, err
	}

	return UserWithoutPassword(&user), nil
}

func GetUserFromUUID(uuid string) (models.User, error) {
	cachedUser, exists := GetCachedUser(uuid)
	if exists {
		return *cachedUser, nil
	}

	user, err := gorm.G[models.User](database.DBConn).Where("id = ?", uuid).First(context.Background())
	if err != nil {
		return models.User{}, err
	}
	return UserWithoutPassword(&user), nil
}

func UserWithoutPassword(user *models.User) models.User {
	returnUser := models.User{
		Model: user.Model,
		Email: user.Email,
		Role:  user.Role,
	}
	return returnUser
}
