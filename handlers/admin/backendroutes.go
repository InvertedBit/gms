package adminhandlers

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterBackendRoutes(router fiber.Router) error {

	// router.Use(middleware.RequireAuthenticatedUser)
	router.Get("/", HandleBackendDashboardRedirect)
	router.Get("/dashboard", HandleBackendDashboard)

	router.Get("/pages", HandlePageList)

	// User routes
	router.Get("/users", HandleUserList)
	router.Get("/users/new", HandleUserNew)
	router.Get("/users/:id", HandleUserEdit)
	router.Post("/users", HandleUserCreate)
	router.Put("/users/:id", HandleUserUpdate)
	router.Delete("/users/:id", HandleUserDelete)

	// Role routes
	router.Get("/roles", HandleRoleList)
	router.Get("/roles/new", HandleRoleNew)
	router.Get("/roles/:id", HandleRoleEdit)
	router.Post("/roles", HandleRoleCreate)
	router.Put("/roles/:id", HandleRoleUpdate)
	router.Delete("/roles/:id", HandleRoleDelete)

	return nil
}
