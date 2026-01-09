package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/invertedbit/gms/html"
	adminviews "github.com/invertedbit/gms/html/views/admin"
	hx "github.com/stackus/hxgo"
	"github.com/stackus/hxgo/hxfiber"
)

func HandleBackendDashboard(c *fiber.Ctx) error {
	dashboardPage := html.Page{
		Title:           "Admin Dashboard",
		PageContent:     adminviews.DashboardPage(),
		LayoutViewModel: GetLayoutModel(c, "Admin Dashboard"),
	}

	return ReturnHandler(c, dashboardPage)
}

func HandleBackendDashboardRedirect(c *fiber.Ctx) error {
	hxfiber.Response(c, hx.Status(http.StatusMovedPermanently), hx.Redirect("/admin/dashboard"))
	return nil
}
