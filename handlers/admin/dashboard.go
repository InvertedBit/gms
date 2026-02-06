package adminhandlers

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	handlerutils "github.com/invertedbit/gms/handlers/utils"
	"github.com/invertedbit/gms/html"
	adminviews "github.com/invertedbit/gms/html/views/admin"
	"github.com/invertedbit/gms/htmx"
)

func HandleBackendDashboard(c fiber.Ctx) error {
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

	return handlerutils.ReturnHandler(c, dashboardPage)
}

func HandleBackendDashboardRedirect(c fiber.Ctx) error {
	hxHeader := new(htmx.HXHeader)
	c.Bind().Header(hxHeader)
	c.Status(http.StatusMovedPermanently)
	htmx.HXRedirect.Set(c, "/admin/dashboard")
	if hxHeader.IsHTMXRequest() {
		return nil
	}
	return c.Redirect().To("/admin/dashboard")
}
