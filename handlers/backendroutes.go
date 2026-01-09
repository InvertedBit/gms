package handlers

import "github.com/gofiber/fiber/v2"

func RegisterBackendRoutes(router fiber.Router) error {

	router.Get("/", HandleBackendDashboardRedirect)
	router.Get("/dashboard", HandleBackendDashboard)

	return nil
}
