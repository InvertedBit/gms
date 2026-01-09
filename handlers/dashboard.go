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
	adminLayoutModel := GetAdminLayoutModel(c, "Admin Dashboard")
	
	// Add breadcrumbs
	adminLayoutModel.AddBreadcrumb("Home", "/")
	adminLayoutModel.AddBreadcrumb("Admin", "/admin")
	adminLayoutModel.AddBreadcrumb("Dashboard", "")
	
	// Add example action button
	adminLayoutModel.AddActionButton("New Item", "/admin/items/new", "ri-add-line", true)
	
	dashboardPage := html.AdminPage{
		Title:                "Admin Dashboard - GMS",
		PageContent:          adminviews.DashboardPage(),
		AdminLayoutViewModel: adminLayoutModel,
	}

	return ReturnHandler(c, dashboardPage)
}

func HandleBackendDashboardRedirect(c *fiber.Ctx) error {
	hxfiber.Response(c, hx.Status(http.StatusMovedPermanently), hx.Redirect("/admin/dashboard"))
	return nil
}
