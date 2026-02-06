package viewmodels

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/invertedbit/gms/htmx"
	"github.com/invertedbit/gms/models"
)

type LayoutType int

const (
	LayoutFull LayoutType = iota
	LayoutBodyOnly
	LayoutPartialOnly
	LayoutPartialWithFragments
	LayoutError
)

type LayoutViewModel struct {
	Page           string
	Navbar         *NavbarViewModel
	IsError        bool
	CopyrightYear  int
	LayoutType     LayoutType
	ToastViewModel *ToastViewModel
	CurrentUser    *models.User
}

func NewLayoutViewModel(page string, navbar *NavbarViewModel, isError bool, copyrightYear int, c fiber.Ctx) *LayoutViewModel {
	hxHeader := new(htmx.HXHeader)
	c.Bind().Header(hxHeader)
	layoutType := LayoutFull
	hxTarget := hxHeader.GetTarget()
	if hxHeader.IsBoosted() || hxHeader.IsHTMXRequest() {
		if hxTarget == "body" || hxTarget == "" {
			layoutType = LayoutBodyOnly
			fmt.Println("Boosted request detected")
		} else {
			layoutType = LayoutPartialOnly
			fmt.Println("HTMX partial request detected")
		}
	}
	return &LayoutViewModel{
		Page:           page,
		Navbar:         navbar,
		IsError:        isError,
		CopyrightYear:  copyrightYear,
		LayoutType:     layoutType,
		ToastViewModel: NewToastViewModel(),
	}
}
