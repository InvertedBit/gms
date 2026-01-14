package adminhandlers

import (
	"github.com/gofiber/fiber/v2"
	handlerutils "github.com/invertedbit/gms/handlers/utils"
	"github.com/invertedbit/gms/html"
	"github.com/invertedbit/gms/html/components"
	adminviews "github.com/invertedbit/gms/html/views/admin"
)

func HandleUserList(c *fiber.Ctx) error {
	adminLayoutModel := GetAdminLayoutModel(c, "Users")

	// Add breadcrumbs
	adminLayoutModel.AddBreadcrumb("Admin", "/admin")
	adminLayoutModel.AddBreadcrumb("Users", "")

	// Add example action button
	adminLayoutModel.AddActionButton("Add user", "/admin/users/new", "ri-add-line", true)

	userTableData := &components.TableData{
		Columns: []components.TableColumn{
			{Name: "id", Label: "ID"},
			{Name: "username", Label: "Username"},
			{Name: "email", Label: "Email"},
			{Name: "role", Label: "Role"},
		},
		Rows: []components.TableRow{
			{Values: map[string]string{"id": "1", "username": "admin", "email": "admin@admin.com", "role": "admin"}},
			{Values: map[string]string{"id": "1", "username": "admin", "email": "admin@admin.com", "role": "admin"}},
			{Values: map[string]string{"id": "1", "username": "admin", "email": "admin@admin.com", "role": "admin"}},
		},
	}

	userListPage := html.AdminPage{
		Title:                "Users - GMS",
		PageContent:          adminviews.UserListPage(userTableData),
		AdminLayoutViewModel: adminLayoutModel,
	}

	return handlerutils.ReturnHandler(c, userListPage)
}
