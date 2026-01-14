package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/sqlite3/v2"
	"github.com/invertedbit/gms/auth"
	adminhandlers "github.com/invertedbit/gms/handlers/admin"
	"github.com/invertedbit/gms/html"
	htmlviews "github.com/invertedbit/gms/html/views"
	"github.com/invertedbit/gms/middleware"
	"github.com/invertedbit/gms/viewmodels"
)

func GetNavbarModel(c *fiber.Ctx) *viewmodels.NavbarViewModel {
	navbarModel := viewmodels.NewNavbarViewModel()

	navbarModel.AddItem(&viewmodels.NavbarMenuItem{
		Label: "Dashboard",
		Link:  "/",
	})

	myRecipes := viewmodels.NavbarMenuItem{
		Label: "Admin",
		Link:  "/admin",
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

	store := sqlite3.New(sqlite3.Config{
		Database:        "./gms-data.sqlite3",
		Table:           "sessions",
		Reset:           false,
		GCInterval:      10 * time.Second,
		MaxOpenConns:    100,
		MaxIdleConns:    100,
		ConnMaxLifetime: 1 * time.Second,
	})

	auth.SessionStore = *session.New(session.Config{
		Storage: store,
	})

	// app.Use(auth.SessionStore)

	app.Static("/", "./assets")

	app.Get("/", HandleViewHome)

	app.Get("/auth/login", HandleLoginView)
	app.Post("/auth/login", HandleLogin)
	app.Get("/auth/logout", HandleLogout)

	adminhandlers.RegisterBackendRoutes(app.Group("/admin", middleware.RequireAuthenticatedUser))

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
