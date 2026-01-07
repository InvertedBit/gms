package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/invertedbit/gms/html"
	htmlviews "github.com/invertedbit/gms/html/views"
)

func HandleBackendDashboard(c *fiber.Ctx) error {
	dashboardPage := html.Page{
		Title:       "Login",
		PageContent: htmlviews.LoginPage(),
	}

	return ReturnHandler(c, dashboardPage)
}
