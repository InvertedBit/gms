package auth

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/invertedbit/gms/models"
)

var SessionStore session.Store

var UserCache map[string]models.User

func CacheUser(user models.User) {
	if UserCache == nil {
		UserCache = make(map[string]models.User)
	}
	UserCache[user.ID.String()] = user
}

func GetCachedUser(userID string) (*models.User, bool) {
	if UserCache == nil {
		UserCache = make(map[string]models.User)
		return nil, false
	}
	user, exists := UserCache[userID]
	return &user, exists
}

func RemoveCachedUser(userID string) {
	if UserCache != nil {
		delete(UserCache, userID)
	}
}
