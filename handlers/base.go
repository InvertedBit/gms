package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func ReturnHandler(c *fiber.Ctx, h http.Handler) error {
	handler := adaptor.HTTPHandler(h)
	return handler(c)
}
