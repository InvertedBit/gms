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
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
)

func HandleUserList(c *fiber.Ctx) error {
	adminLayoutModel := GetAdminLayoutModel(c, "Users")

	// Add breadcrumbs
	adminLayoutModel.AddBreadcrumb("Admin", "/admin")
	adminLayoutModel.AddBreadcrumb("Users", "")

	// Add action button
	adminLayoutModel.AddActionButton("Add user", "/admin/users/new", "ri-add-line", true)

	// Fetch users from database
	var users []models.User
	database.DBConn.Preload("Role").Order("email ASC").Find(&users)

	// Build table data
	userTableData := &components.TableData{
		Columns: []components.TableColumn{
			{Name: "email", Label: "Email"},
			{Name: "role", Label: "Role"},
		},
		Rows:          []components.TableRow{},
		Editable:      true,
		Deletable:     true,
		EditRoute:     "/admin/users",
		DeleteRoute:   "/admin/users",
		IDField:       "id",
		RefreshTarget: "#data-table-container",
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
	return handlerutils.ReturnHandler(c, adminviews.UserFormModal(vm))
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
	return handlerutils.ReturnHandler(c, adminviews.UserFormModal(vm))
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
		return c.Status(400).SendString("Error creating user")
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
		return c.Status(400).SendString("Error updating user")
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
	// Fetch users from database
	var users []models.User
	database.DBConn.Preload("Role").Order("email ASC").Find(&users)

	// Build table data
	userTableData := &components.TableData{
		Columns: []components.TableColumn{
			{Name: "email", Label: "Email"},
			{Name: "role", Label: "Role"},
		},
		Rows:          []components.TableRow{},
		Editable:      true,
		Deletable:     true,
		EditRoute:     "/admin/users",
		DeleteRoute:   "/admin/users",
		IDField:       "id",
		RefreshTarget: "#data-table-container",
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

	return handlerutils.ReturnHandler(c, components.DataTable(userTableData))
}

