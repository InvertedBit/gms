package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/invertedbit/gms/auth"
	hx "github.com/stackus/hxgo"
	"github.com/stackus/hxgo/hxfiber"
)

func RequireAuthenticatedUser(c *fiber.Ctx) error {
	isAuthenticated := false
	fmt.Println("Auth middleware called")
	session, err := auth.SessionStore.Get(c)
	if err != nil {
		fmt.Println("could not retrieve session")
	}
	userId := session.Get("user_uuid")
	fmt.Println(fmt.Sprintf("Got user_uuid: %v", userId))

	if userId != "" && userId != nil {
		isAuthenticated = true
	}

	if isAuthenticated {
		return c.Next()
	} else {
		if hxfiber.IsHtmx(c) {
			hxfiber.Response(c, hx.Location("/auth/login"))
		}
		return c.Redirect("/auth/login")
	}
}
