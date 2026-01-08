package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/invertedbit/gms/auth"
	"github.com/invertedbit/gms/html"
	htmlviews "github.com/invertedbit/gms/html/views"
	hx "github.com/stackus/hxgo"
	"github.com/stackus/hxgo/hxfiber"
)

func HandleLoginView(c *fiber.Ctx) error {
	title := "Login"

	session, err := auth.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}
	visitedLoginPage := session.Get("visited_login_page")

	if visitedLoginPage == nil {
		session.Set("visited_login_page", true)
		session.Save()
	} else {
		title = "Login again"
	}

	loginPage := html.Page{
		Title:           title,
		PageContent:     htmlviews.LoginPage(),
		LayoutViewModel: GetLayoutModel(c, "Login"),
	}

	return ReturnHandler(c, loginPage)
}

func HandleLogin(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		eventData := make(map[string]interface{})
		eventData["message"] = "Invalid credentials"
		hxfiber.Response(c, hx.Status(http.StatusBadRequest), hx.Trigger(hx.Event("gms:login-failed", eventData)))
	}

	// Perform login logic here (e.g., check credentials, create session, etc.)

	return nil
}
