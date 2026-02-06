package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/invertedbit/gms/auth"
	handlerutils "github.com/invertedbit/gms/handlers/utils"
	"github.com/invertedbit/gms/html"
	htmlviews "github.com/invertedbit/gms/html/views"
	"github.com/invertedbit/gms/htmx"
)

func HandleLoginView(c fiber.Ctx) error {
	title := "Login"

	session := session.FromContext(c)

	visitedLoginPage := session.Get("visited_login_page")

	if visitedLoginPage == nil {
		session.Set("visited_login_page", true)
	} else {
		title = "Login again"
	}

	loginPage := html.Page{
		Title:           title,
		PageContent:     htmlviews.LoginPage(),
		LayoutViewModel: GetLayoutModel(c, title),
	}

	return handlerutils.ReturnHandler(c, loginPage)
}

func HandleLogin(c fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		eventData := make(map[string]interface{})
		eventData["message"] = "Please provide both username and password"
		c.Status(http.StatusBadRequest)
		event := make(map[string]interface{})
		event["gms:login-failed"] = eventData
		htmx.HXTrigger.SetJson(c, event)
		return nil
	}

	user, err := auth.AuthenticateUser(username, password)
	if err != nil {
		eventData := make(map[string]interface{})
		eventData["message"] = "Invalid credentials"
		c.Status(http.StatusUnauthorized)
		event := make(map[string]interface{})
		event["gms:login-failed"] = eventData
		htmx.HXTrigger.SetJson(c, event)
		return nil
	}

	session := session.FromContext(c)
	fmt.Printf("Logging in user: %s\n", user.ID)
	session.Set("user_uuid", user.ID.String())
	auth.CacheUser(user)

	eventData := make(map[string]interface{})
	eventData["message"] = "Login successful"
	event := make(map[string]interface{})
	event["gms:login-success"] = eventData
	htmx.HXTrigger.SetJson(c, event)
	htmx.HXRedirect.Set(c, "/admin/dashboard")
	c.Status(http.StatusOK)

	return nil
}

func HandleLogout(c fiber.Ctx) error {
	session := session.FromContext(c)

	userID := session.Get("user_uuid")
	if userID != nil {
		auth.RemoveCachedUser(fmt.Sprintf("%v", userID))
	}

	session.Destroy()
	htmx.HXRedirect.Set(c, "/")
	c.Status(http.StatusOK)

	return nil
}
