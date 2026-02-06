package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/invertedbit/gms/htmx"
)

func RequireAuthenticatedUser(c fiber.Ctx) error {
	isAuthenticated := false
	fmt.Println("Auth middleware called")
	session := session.FromContext(c)
	userId := session.Get("user_uuid")
	fmt.Println(fmt.Sprintf("Got user_uuid: %v", userId))

	if userId != "" && userId != nil {
		isAuthenticated = true
	}

	if isAuthenticated {
		return c.Next()
	} else {
		hxHeader := new(htmx.HXHeader)
		c.Bind().Header(hxHeader)
		if hxHeader.IsHTMXRequest() {
			htmx.HXLocation.Set(c, "/auth/login")
		}
		return c.Redirect().To("/auth/login")
	}
}
