package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/invertedbit/gms/auth"
	"github.com/invertedbit/gms/html"
	htmlviews "github.com/invertedbit/gms/html/views"
	"github.com/invertedbit/gms/viewmodels"
)

func GetNavbarModel(c *fiber.Ctx) *viewmodels.NavbarViewModel {
	navbarModel := viewmodels.NewNavbarViewModel()

	navbarModel.AddItem(&viewmodels.NavbarMenuItem{
		Label: "Dashboard",
		Link:  "/",
	})

	myRecipes := viewmodels.NavbarMenuItem{
		Label: "My Recipes",
		Link:  "/recipes",
	}

	categories := viewmodels.NavbarMenuItem{
		Label: "Categories",
		Link:  "/recipe-categories",
	}

	navbarModel.AddItem(&viewmodels.NavbarMenuItem{
		Label:    "Recipes",
		Link:     "#",
		Children: []*viewmodels.NavbarMenuItem{&myRecipes, &categories},
	})

	navbarModel.AddItem(&viewmodels.NavbarMenuItem{
		Label: "Ingredients",
		Link:  "/ingredients",
	})

	return navbarModel
}

func GetLayoutModel(c *fiber.Ctx, title string) *viewmodels.LayoutViewModel {
	layoutViewModel := viewmodels.NewLayoutViewModel(title, GetNavbarModel(c), false, 2025, c)
	session, err := auth.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}
	userId := session.Get("user_uuid")
	if userId != nil && userId != "" {
		currentUser, err := auth.GetUserFromUUID(fmt.Sprintf("%v", userId))
		if err != nil {
			panic(err)
		}
		layoutViewModel.CurrentUser = &currentUser
	}

	return layoutViewModel
}

func New() *fiber.App {
	app := fiber.New()

	auth.SessionStore = session.New()

	// app.Use(auth.SessionStore)

	app.Static("/", "./assets")

	app.Get("/", HandleViewHome)

	app.Get("/auth/login", HandleLoginView)
	app.Post("/auth/login", HandleLogin)

	RegisterBackendRoutes(app.Group("/backend"))

	app.Get("*", HandleNotFound)

	return app
}

func HandleNotFound(c *fiber.Ctx) error {
	// notFound := views.NotFoundPage(GetLayoutModel(c, "Not Found", false))
	notFoundPage := html.Page{
		Title:           "404 Not Found - GoCook",
		PageContent:     htmlviews.NotFoundPage(viewmodels.NewNotFoundViewModel(c)),
		LayoutViewModel: GetLayoutModel(c, "Not Found"),
	}

	handler := adaptor.HTTPHandler(notFoundPage)
	return handler(c)
}

func HandleViewHome(c *fiber.Ctx) error {
	// home := views.HomePage(GetLayoutModel(c, "Dashboard", false), "Hello World!")

	homePage := html.Page{
		Title:           "Home - GoCook",
		PageContent:     htmlviews.HomePage(),
		LayoutViewModel: GetLayoutModel(c, "Dashboard"),
	}

	handler := adaptor.HTTPHandler(homePage)

	return handler(c)
}
