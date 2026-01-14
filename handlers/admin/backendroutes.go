package adminhandlers

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterBackendRoutes(router fiber.Router) error {

	// router.Use(middleware.RequireAuthenticatedUser)
	router.Get("/", HandleBackendDashboardRedirect)
	router.Get("/dashboard", HandleBackendDashboard)

	router.Get("/pages", HandlePageList)

	router.Get("/users", HandleUserList)

	return nil
}
