package handlerutils

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"maragu.dev/gomponents"
)

func ReturnHandler(c fiber.Ctx, h http.Handler) error {
	handler := adaptor.HTTPHandler(h)
	return handler(c)
}

// RenderNode renders a gomponents.Node directly to the response
func RenderNode(c fiber.Ctx, node gomponents.Node) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return node.Render(c.Response().BodyWriter())
}
