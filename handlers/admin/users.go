package adminhandlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/invertedbit/gms/database"
	handlerutils "github.com/invertedbit/gms/handlers/utils"
	"github.com/invertedbit/gms/html"
	admincomponents "github.com/invertedbit/gms/html/components/admin"
	adminviews "github.com/invertedbit/gms/html/views/admin"
	"github.com/invertedbit/gms/models"
	"github.com/invertedbit/gms/viewmodels"
	"github.com/stackus/hxgo/hxfiber"
	"golang.org/x/crypto/bcrypt"
)

func HandleUserList(c *fiber.Ctx) error {
	adminLayoutModel := GetAdminLayoutModel(c, "Users")

	// Add breadcrumbs
	adminLayoutModel.AddBreadcrumb("Admin", "/admin")
	adminLayoutModel.AddBreadcrumb("Users", "")

	// Add action button
	adminLayoutModel.AddActionButton("Add user", "/admin/users/new", "ri-add-line", true)

	if hxfiber.IsHtmx(c) {
		if hxfiber.GetTarget(c) == "#users-list" {
			adminLayoutModel.LayoutType = viewmodels.LayoutPartialOnly
		}
	}

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
	if err := database.DBConn.Order("name ASC").Find(&roles).Error; err != nil {
		fmt.Println(err)
		return c.Status(500).SendString("Error fetching roles")
	}

	vm := viewmodels.NewUserFormViewModel(nil, roles, false)
	return handlerutils.RenderNode(c, adminviews.UserFormModal(vm))
}

func HandleUserEdit(c *fiber.Ctx) error {
	fmt.Println("HandleUserEdit called")
	userID := c.Params("id")
	fmt.Printf("Got user id: %s\n", userID)
	var user models.User
	if err := database.DBConn.Where("id = ?", userID).First(&user).Error; err != nil {
		fmt.Println(err)
		return c.Status(404).SendString("User not found")
	}

	var roles []models.Role
	if err := database.DBConn.Order("name ASC").Find(&roles).Error; err != nil {
		fmt.Println(err)
		return c.Status(500).SendString("Error fetching roles")
	}

	vm := viewmodels.NewUserFormViewModel(&user, roles, true)
	return handlerutils.RenderNode(c, adminviews.UserFormModal(vm))
}

func HandleUserCreate(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	roleSlug := c.FormValue("role_slug")

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).SendString("Error hashing password")
	}

	user := models.User{
		Model: models.Model{
			ID: uuid.New(),
		},
		Email:             email,
		EncryptedPassword: string(hashedPassword),
	}

	if roleSlug != "" {
		user.RoleSlug = roleSlug
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

	roleSlug := c.FormValue("role_slug")
	if roleSlug != "" {
		user.RoleSlug = roleSlug
	} else {
		user.RoleSlug = ""
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
	return handlerutils.RenderNode(c, admincomponents.DataTable(userTableData))
}

// buildUserTableData fetches users from database and builds table data
func buildUserTableData() *admincomponents.TableData {
	// Fetch users from database
	var users []models.User
	database.DBConn.Preload("Role").Order("email ASC").Find(&users)

	// Build table data
	userTableData := &admincomponents.TableData{
		Columns: []admincomponents.TableColumn{
			{Name: "email", Label: "Email"},
			{Name: "role", Label: "Role"},
		},
		Rows:             []admincomponents.TableRow{},
		Editable:         true,
		Deletable:        true,
		EditRoute:        "/admin/users/edit",
		DeleteRoute:      "/admin/users",
		IDField:          "id",
		RefreshTarget:    "#data-table-container",
		DeleteConfirmMsg: "Are you sure you want to delete this user? This action cannot be undone.",
	}

	for _, user := range users {
		roleName := ""
		if user.RoleSlug != "" {
			var role models.Role
			if err := database.DBConn.Where("slug = ?", user.RoleSlug).First(&role).Error; err == nil {
				roleName = role.Name
			}
		}
		userTableData.Rows = append(userTableData.Rows, admincomponents.TableRow{
			Values: map[string]string{
				"id":    user.ID.String(),
				"email": user.Email,
				"role":  roleName,
			},
		})
	}

	return userTableData
}
