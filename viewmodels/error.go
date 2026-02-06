package viewmodels

import (
	"github.com/gofiber/fiber/v3"
	"github.com/invertedbit/gms/htmx"
)

type NotFoundViewModel struct {
	CurrentURL string
}

func NewNotFoundViewModel(c fiber.Ctx) *NotFoundViewModel {
	hxHeader := new(htmx.HXHeader)
	c.Bind().Header(hxHeader)
	currentUrl := hxHeader.GetCurrentURL()
	return &NotFoundViewModel{
		CurrentURL: currentUrl,
	}
}

func (vm *NotFoundViewModel) HasCurrentURL() bool {
	return vm.CurrentURL != ""
}
