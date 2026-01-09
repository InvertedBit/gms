package handlers

import (
	"fmt"
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
		LayoutViewModel: GetLayoutModel(c, title),
	}

	return ReturnHandler(c, loginPage)
}

func HandleLogin(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		eventData := make(map[string]interface{})
		eventData["message"] = "Please provide both username and password"
		hxfiber.Response(c, hx.Status(http.StatusBadRequest), hx.Trigger(hx.Event("gms:login-failed", eventData)))
		return nil
	}

	user, err := auth.AuthenticateUser(username, password)
	if err != nil {
		eventData := make(map[string]interface{})
		eventData["message"] = "Invalid credentials"
		hxfiber.Response(c, hx.Status(http.StatusUnauthorized), hx.Trigger(hx.Event("gms:login-failed", eventData)))
		return nil
	}

	session, err := auth.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Logging in user: %s\n", user.ID)
	session.Set("user_uuid", user.ID.String())
	auth.CacheUser(user)
	session.Save()

	eventData := make(map[string]interface{})
	eventData["message"] = "Login successful"
	hxfiber.Response(c, hx.Status(http.StatusOK), hx.Trigger(hx.Event("gms:login-success", eventData)), hx.Redirect("/admin/dashboard"))

	return nil
}

func HandleLogout(c *fiber.Ctx) error {
	session, err := auth.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	userID := session.Get("user_uuid")
	if userID != nil {
		auth.RemoveCachedUser(fmt.Sprintf("%v", userID))
	}

	session.Destroy()
	hxfiber.Response(c, hx.Status(http.StatusOK), hx.Redirect("/"))

	return nil
}
