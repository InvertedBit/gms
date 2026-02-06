package viewmodels

import (
	"sort"

	"github.com/gofiber/fiber/v3"
	"github.com/invertedbit/gms/htmx"
)

// AdminNavigationItem represents an item in the admin navigation menu
type AdminNavigationItem struct {
	Label    string
	Link     string
	Icon     string // RemixIcon class name
	IsOpen   bool   // Whether the navigation item is expanded (for items with children)
	IsActive bool   // Whether the navigation item is currently active
	Order    int
	Children []*AdminNavigationItem
}

// Breadcrumb represents a single breadcrumb item
type Breadcrumb struct {
	Label string
	Link  string
}

// ActionButton represents an action button in the admin header
type ActionButton struct {
	Label   string
	Link    string
	Icon    string
	Primary bool // Whether this is a primary action button
}

// AdminLayoutViewModel contains data for the admin panel layout
type AdminLayoutViewModel struct {
	Page           string
	Title          string
	Breadcrumbs    []Breadcrumb
	Navigation     []*AdminNavigationItem
	ActionButtons  []ActionButton
	LayoutType     LayoutType
	ToastViewModel *ToastViewModel
	CurrentUser    interface{} // Can be *models.User, using interface to avoid circular dependency
}

// NewAdminLayoutViewModel creates a new admin layout view model
func NewAdminLayoutViewModel(page string, title string, c fiber.Ctx) *AdminLayoutViewModel {
	hxHeader := new(htmx.HXHeader)
	c.Bind().Header(hxHeader)
	layoutType := LayoutFull
	hxTarget := hxHeader.GetTarget()
	if hxHeader.IsBoosted() || hxHeader.IsHTMXRequest() {
		if hxTarget == "body" || hxTarget == "" {
			layoutType = LayoutBodyOnly
		} else {
			layoutType = LayoutPartialOnly
		}
	}

	return &AdminLayoutViewModel{
		Page:           page,
		Title:          title,
		Breadcrumbs:    []Breadcrumb{},
		Navigation:     []*AdminNavigationItem{},
		ActionButtons:  []ActionButton{},
		LayoutType:     layoutType,
		ToastViewModel: NewToastViewModel(),
	}
}

// AddBreadcrumb adds a breadcrumb to the layout
func (a *AdminLayoutViewModel) AddBreadcrumb(label string, link string) {
	a.Breadcrumbs = append(a.Breadcrumbs, Breadcrumb{
		Label: label,
		Link:  link,
	})
}

// AddNavigationItem adds a navigation item to the layout
func (a *AdminLayoutViewModel) AddNavigationItem(item *AdminNavigationItem) {
	a.Navigation = append(a.Navigation, item)
}

// AddActionButton adds an action button to the layout
func (a *AdminLayoutViewModel) AddActionButton(label string, link string, icon string, primary bool) {
	a.ActionButtons = append(a.ActionButtons, ActionButton{
		Label:   label,
		Link:    link,
		Icon:    icon,
		Primary: primary,
	})
}

func (a *AdminLayoutViewModel) GetNavigation() []*AdminNavigationItem {
	sort.SliceStable(a.Navigation, func(i, j int) bool {
		return a.Navigation[i].Order < a.Navigation[j].Order
	})
	return a.Navigation
}
