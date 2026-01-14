package adminhandlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/invertedbit/gms/auth"
	"github.com/invertedbit/gms/viewmodels"
)

// GetDefaultAdminNavigation returns the default navigation structure for the admin panel
func GetDefaultAdminNavigation() []*viewmodels.AdminNavigationItem {
	return []*viewmodels.AdminNavigationItem{
		{
			Label: "Dashboard",
			Link:  "/admin/dashboard",
			Icon:  "ri-dashboard-line",
			Order: 10,
		},
		{
			Label: "Content",
			Icon:  "ri-file-list-line",
			Order: 40,
			Children: []*viewmodels.AdminNavigationItem{
				{
					Label: "Pages",
					Link:  "/admin/pages",
					Icon:  "ri-restaurant-line",
				},
				{
					Label: "Posts",
					Link:  "/admin/posts",
					Icon:  "ri-plant-line",
				},
				{
					Label: "Categories",
					Link:  "/admin/categories",
					Icon:  "ri-folder-line",
				},
			},
		},
		{
			Label: "Users",
			Icon:  "ri-user-line",
			Order: 90,
			Children: []*viewmodels.AdminNavigationItem{
				{
					Label: "All Users",
					Link:  "/admin/users",
					Icon:  "ri-user-line",
				},
				{
					Label: "Roles",
					Link:  "/admin/roles",
					Icon:  "ri-shield-user-line",
				},
			},
		},
		{
			Label: "Settings",
			Link:  "/admin/settings",
			Icon:  "ri-settings-line",
			Order: 100,
		},
	}
}

// GetAdminLayoutModel creates an admin layout view model with default navigation
func GetAdminLayoutModel(c *fiber.Ctx, title string) *viewmodels.AdminLayoutViewModel {
	layoutViewModel := viewmodels.NewAdminLayoutViewModel(title, title, c)

	// Add default navigation
	for _, item := range GetDefaultAdminNavigation() {
		layoutViewModel.AddNavigationItem(item)
	}

	// Get current user if logged in
	session, err := auth.SessionStore.Get(c)
	if err == nil {
		userId := session.Get("user_uuid")
		if userId != nil && userId != "" {
			currentUser, err := auth.GetUserFromUUID(fmt.Sprintf("%v", userId))
			if err == nil {
				layoutViewModel.CurrentUser = &currentUser
			}
		}
	}

	return layoutViewModel
}
