package auth

import (
	"github.com/invertedbit/gms/database"
	"github.com/invertedbit/gms/models"
	"golang.org/x/crypto/bcrypt"
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
	var user models.User
	if err := database.DBConn.Where("email = ?", email).First(&user).Error; err != nil {
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

	var user models.User
	if err := database.DBConn.Where("id = ?", uuid).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return UserWithoutPassword(&user), nil
}

func UserWithoutPassword(user *models.User) models.User {
	returnUser := models.User{
		Model:    user.Model,
		Email:    user.Email,
		RoleSlug: user.RoleSlug,
	}
	return returnUser
}
