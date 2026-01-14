package adminhandlers

import (
	"github.com/gofiber/fiber/v2"
	handlerutils "github.com/invertedbit/gms/handlers/utils"
	"github.com/invertedbit/gms/html"
	adminviews "github.com/invertedbit/gms/html/views/admin"
)

func HandlePageList(c *fiber.Ctx) error {
	adminLayoutModel := GetAdminLayoutModel(c, "Pages")

	// Add breadcrumbs
	adminLayoutModel.AddBreadcrumb("Admin", "/admin")
	adminLayoutModel.AddBreadcrumb("Pages", "")

	// Add example action button
	adminLayoutModel.AddActionButton("Add page", "/admin/pages/new", "ri-add-line", true)

	userListPage := html.AdminPage{
		Title:                "Users - GMS",
		PageContent:          adminviews.PageListPage(),
		AdminLayoutViewModel: adminLayoutModel,
	}

	return handlerutils.ReturnHandler(c, userListPage)
}
