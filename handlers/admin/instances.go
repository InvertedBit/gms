package adminhandlers

import (
	"github.com/gofiber/fiber/v2"
	handlerutils "github.com/invertedbit/gms/handlers/utils"
	"github.com/invertedbit/gms/html"
	adminviews "github.com/invertedbit/gms/html/views/admin"
	"github.com/invertedbit/gms/viewmodels"
)

func addInstancesBreadcrumbs(adminLayoutModel *viewmodels.AdminLayoutViewModel) {
	adminLayoutModel.AddBreadcrumb("Admin", "/admin")
	adminLayoutModel.AddBreadcrumb("Pages", "/admin/pages")
}

func HandleInstanceList(c *fiber.Ctx) error {
	adminLayoutModel := GetAdminLayoutModel(c, "Instances")
	addInstancesBreadcrumbs(adminLayoutModel)

	adminLayoutModel.AddActionButton("Add instance", "/admin/instances/new", "ri-add-line", true)

	instanceListPage := html.AdminPage{
		Title:                "Instances - GMS",
		PageContent:          adminviews.PageListPage(buildPageTableData()),
		AdminLayoutViewModel: adminLayoutModel,
	}

	return handlerutils.ReturnHandler(c, instanceListPage)
}
