package adminhandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/invertedbit/gms/database"
	handlerutils "github.com/invertedbit/gms/handlers/utils"
	"github.com/invertedbit/gms/html"
	"github.com/invertedbit/gms/html/components"
	adminviews "github.com/invertedbit/gms/html/views/admin"
	"github.com/invertedbit/gms/models"
	"github.com/invertedbit/gms/viewmodels"
	"golang.org/x/crypto/bcrypt"
)

func HandleUserList(c *fiber.Ctx) error {
	adminLayoutModel := GetAdminLayoutModel(c, "Users")

	// Add breadcrumbs
	adminLayoutModel.AddBreadcrumb("Admin", "/admin")
	adminLayoutModel.AddBreadcrumb("Users", "")

	// Add action button
	adminLayoutModel.AddActionButton("Add user", "/admin/users/new", "ri-add-line", true)

	userTableData := buildUserTableData()

	userListPage := html.AdminPage{
		Title:                "Users - GMS",
		PageContent:          adminviews.UserListPage(userTableData),
		AdminLayoutViewModel: adminLayoutModel,
	}

	return handlerutils.ReturnHandler(c, userListPage)
}

func HandleUserNew(c *fiber.Ctx) error {
	var roles []models.Role
	database.DBConn.Order("name ASC").Find(&roles)

	vm := viewmodels.NewUserFormViewModel(nil, roles, false)
	return handlerutils.RenderNode(c, adminviews.UserFormModal(vm))
}

func HandleUserEdit(c *fiber.Ctx) error {
	userID := c.Params("id")

	var user models.User
	if err := database.DBConn.Preload("Role").Where("id = ?", userID).First(&user).Error; err != nil {
		return c.Status(404).SendString("User not found")
	}

	var roles []models.Role
	database.DBConn.Order("name ASC").Find(&roles)

	vm := viewmodels.NewUserFormViewModel(&user, roles, true)
	return handlerutils.RenderNode(c, adminviews.UserFormModal(vm))
}

func HandleUserCreate(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	roleIDStr := c.FormValue("role_id")

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).SendString("Error hashing password")
	}

	user := models.User{
		Email:             email,
		EncryptedPassword: string(hashedPassword),
	}

	if roleIDStr != "" {
		roleID, err := uuid.Parse(roleIDStr)
		if err != nil {
			return c.Status(400).SendString("Invalid role ID")
		}
		user.RoleID = &roleID
	}

	if err := database.DBConn.Create(&user).Error; err != nil {
		// Check for unique constraint violation
		if isDuplicateKeyError(err) {
			c.Status(400)
			var roles []models.Role
			database.DBConn.Order("name ASC").Find(&roles)
			vm := viewmodels.NewUserFormViewModel(&user, roles, false)
			vm.FormErrors["email"] = "A user with this email already exists"
			return handlerutils.RenderNode(c, adminviews.UserFormModal(vm))
		}
		return c.Status(500).SendString("Error creating user")
	}

	// Return updated table
	return renderUserTable(c)
}

func HandleUserUpdate(c *fiber.Ctx) error {
	userID := c.Params("id")

	var user models.User
	if err := database.DBConn.Where("id = ?", userID).First(&user).Error; err != nil {
		return c.Status(404).SendString("User not found")
	}

	user.Email = c.FormValue("email")
	
	// Only update password if provided
	password := c.FormValue("password")
	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(500).SendString("Error hashing password")
		}
		user.EncryptedPassword = string(hashedPassword)
	}

	roleIDStr := c.FormValue("role_id")
	if roleIDStr != "" {
		roleID, err := uuid.Parse(roleIDStr)
		if err != nil {
			return c.Status(400).SendString("Invalid role ID")
		}
		user.RoleID = &roleID
	} else {
		user.RoleID = nil
	}

	if err := database.DBConn.Save(&user).Error; err != nil {
		// Check for unique constraint violation
		if isDuplicateKeyError(err) {
			c.Status(400)
			var roles []models.Role
			database.DBConn.Order("name ASC").Find(&roles)
			vm := viewmodels.NewUserFormViewModel(&user, roles, true)
			vm.FormErrors["email"] = "A user with this email already exists"
			return handlerutils.RenderNode(c, adminviews.UserFormModal(vm))
		}
		return c.Status(500).SendString("Error updating user")
	}

	// Return updated table
	return renderUserTable(c)
}

func HandleUserDelete(c *fiber.Ctx) error {
	userID := c.Params("id")

	if err := database.DBConn.Delete(&models.User{}, "id = ?", userID).Error; err != nil {
		return c.Status(400).SendString("Error deleting user")
	}

	// Return updated table
	return renderUserTable(c)
}

func renderUserTable(c *fiber.Ctx) error {
	userTableData := buildUserTableData()
	return handlerutils.RenderNode(c, components.DataTable(userTableData))
}

// buildUserTableData fetches users from database and builds table data
func buildUserTableData() *components.TableData {
	// Fetch users from database
	var users []models.User
	database.DBConn.Preload("Role").Order("email ASC").Find(&users)

	// Build table data
	userTableData := &components.TableData{
		Columns: []components.TableColumn{
			{Name: "email", Label: "Email"},
			{Name: "role", Label: "Role"},
		},
		Rows:             []components.TableRow{},
		Editable:         true,
		Deletable:        true,
		EditRoute:        "/admin/users",
		DeleteRoute:      "/admin/users",
		IDField:          "id",
		RefreshTarget:    "#data-table-container",
		DeleteConfirmMsg: "Are you sure you want to delete this user? This action cannot be undone.",
	}

	for _, user := range users {
		roleName := ""
		if user.Role != nil {
			roleName = user.Role.Name
		}
		userTableData.Rows = append(userTableData.Rows, components.TableRow{
			Values: map[string]string{
				"id":    user.ID.String(),
				"email": user.Email,
				"role":  roleName,
			},
		})
	}

	return userTableData
}

