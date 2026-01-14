package adminhandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/invertedbit/gms/database"
	handlerutils "github.com/invertedbit/gms/handlers/utils"
	"github.com/invertedbit/gms/html"
	"github.com/invertedbit/gms/html/components"
	adminviews "github.com/invertedbit/gms/html/views/admin"
	"github.com/invertedbit/gms/models"
	"github.com/invertedbit/gms/viewmodels"
)

func HandleRoleList(c *fiber.Ctx) error {
	adminLayoutModel := GetAdminLayoutModel(c, "Roles")

	// Add breadcrumbs
	adminLayoutModel.AddBreadcrumb("Admin", "/admin")
	adminLayoutModel.AddBreadcrumb("Roles", "")

	// Add action button
	adminLayoutModel.AddActionButton("Add role", "/admin/roles/new", "ri-add-line", true)

	roleTableData := buildRoleTableData()

	roleListPage := html.AdminPage{
		Title:                "Roles - GMS",
		PageContent:          adminviews.RoleListPage(roleTableData),
		AdminLayoutViewModel: adminLayoutModel,
	}

	return handlerutils.ReturnHandler(c, roleListPage)
}

func HandleRoleNew(c *fiber.Ctx) error {
	vm := viewmodels.NewRoleFormViewModel(nil, false)
	return handlerutils.RenderNode(c, adminviews.RoleFormModal(vm))
}

func HandleRoleEdit(c *fiber.Ctx) error {
	roleID := c.Params("id")

	var role models.Role
	if err := database.DBConn.Where("id = ?", roleID).First(&role).Error; err != nil {
		return c.Status(404).SendString("Role not found")
	}

	vm := viewmodels.NewRoleFormViewModel(&role, true)
	return handlerutils.RenderNode(c, adminviews.RoleFormModal(vm))
}

func HandleRoleCreate(c *fiber.Ctx) error {
	name := c.FormValue("name")
	description := c.FormValue("description")

	role := models.Role{
		Name:        name,
		Description: description,
	}

	if err := database.DBConn.Create(&role).Error; err != nil {
		// Check for unique constraint violation
		if isDuplicateKeyError(err) {
			c.Status(400)
			vm := viewmodels.NewRoleFormViewModel(&role, false)
			vm.FormErrors["name"] = "A role with this name already exists"
			return handlerutils.RenderNode(c, adminviews.RoleFormModal(vm))
		}
		return c.Status(500).SendString("Error creating role")
	}

	// Return updated table
	return renderRoleTable(c)
}

func HandleRoleUpdate(c *fiber.Ctx) error {
	roleID := c.Params("id")

	var role models.Role
	if err := database.DBConn.Where("id = ?", roleID).First(&role).Error; err != nil {
		return c.Status(404).SendString("Role not found")
	}

	role.Name = c.FormValue("name")
	role.Description = c.FormValue("description")

	if err := database.DBConn.Save(&role).Error; err != nil {
		// Check for unique constraint violation
		if isDuplicateKeyError(err) {
			c.Status(400)
			vm := viewmodels.NewRoleFormViewModel(&role, true)
			vm.FormErrors["name"] = "A role with this name already exists"
			return handlerutils.RenderNode(c, adminviews.RoleFormModal(vm))
		}
		return c.Status(500).SendString("Error updating role")
	}

	// Return updated table
	return renderRoleTable(c)
}

func HandleRoleDelete(c *fiber.Ctx) error {
	roleID := c.Params("id")

	if err := database.DBConn.Delete(&models.Role{}, "id = ?", roleID).Error; err != nil {
		return c.Status(400).SendString("Error deleting role")
	}

	// Return updated table
	return renderRoleTable(c)
}

func renderRoleTable(c *fiber.Ctx) error {
	roleTableData := buildRoleTableData()
	return handlerutils.RenderNode(c, components.DataTable(roleTableData))
}

// buildRoleTableData fetches roles from database and builds table data
func buildRoleTableData() *components.TableData {
	// Fetch roles from database
	var roles []models.Role
	database.DBConn.Order("name ASC").Find(&roles)

	// Build table data
	roleTableData := &components.TableData{
		Columns: []components.TableColumn{
			{Name: "name", Label: "Name"},
			{Name: "description", Label: "Description"},
		},
		Rows:             []components.TableRow{},
		Editable:         true,
		Deletable:        true,
		EditRoute:        "/admin/roles",
		DeleteRoute:      "/admin/roles",
		IDField:          "id",
		RefreshTarget:    "#data-table-container",
		DeleteConfirmMsg: "Are you sure you want to delete this role? Users assigned to this role will have their role removed.",
	}

	for _, role := range roles {
		roleTableData.Rows = append(roleTableData.Rows, components.TableRow{
			Values: map[string]string{
				"id":          role.ID.String(),
				"name":        role.Name,
				"description": role.Description,
			},
		})
	}

	return roleTableData
}
