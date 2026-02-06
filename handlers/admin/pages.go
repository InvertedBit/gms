package adminhandlers

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/invertedbit/gms/database"
	handlerutils "github.com/invertedbit/gms/handlers/utils"
	"github.com/invertedbit/gms/html"
	admincomponents "github.com/invertedbit/gms/html/components/admin"
	adminviews "github.com/invertedbit/gms/html/views/admin"
	"github.com/invertedbit/gms/models"
	"github.com/invertedbit/gms/viewmodels"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func addPagesBreadcrumbs(adminLayoutModel *viewmodels.AdminLayoutViewModel) {
	adminLayoutModel.AddBreadcrumb("Admin", "/admin")
	adminLayoutModel.AddBreadcrumb("Pages", "/admin/pages")
}

func HandlePageList(c fiber.Ctx) error {
	adminLayoutModel := GetAdminLayoutModel(c, "Pages")

	addPagesBreadcrumbs(adminLayoutModel)

	// Add example action button
	adminLayoutModel.AddActionButton("Add page", "/admin/pages/new", "ri-add-line", true)

	pageListPage := html.AdminPage{
		Title:                "Pages - GMS",
		PageContent:          adminviews.PageListPage(buildPageTableData()),
		AdminLayoutViewModel: adminLayoutModel,
	}

	return handlerutils.ReturnHandler(c, pageListPage)
}

func HandlePageNew(c fiber.Ctx) error {
	layouts, err := gorm.G[models.Layout](database.DBConn).Where("entry_type = ?", models.EntryTypeActive).Order("name ASC").Find(context.Background())
	if err != nil {
		return c.Status(500).SendString("Error fetching layouts")
	}

	vm := viewmodels.NewPageFormViewModel(nil, layouts, false)

	return handlerutils.RenderNode(c, adminviews.PageFormModal(vm))
}

func HandlePageCreate(c fiber.Ctx) error {
	title := c.FormValue("title")
	slug := c.FormValue("slug")
	layoutSlug := c.FormValue("layout")

	if slug == "" {
		slug = models.GenerateSlug(title)
	}
	if layoutSlug == "/none" {
		layoutSlug = ""
	}

	layoutData := &models.Layout{
		Slug: layoutSlug,
	}

	if layoutSlug != "" {
		layout, err := gorm.G[models.Layout](database.DBConn).Where("slug = ?", layoutSlug).First(context.Background())
		if err != nil {
			layoutData = &layout
		}
	}

	page := &models.Page{
		Title: title,
		Slug:  slug,
	}
	if layoutSlug != "" && layoutData != nil && layoutData.Slug != "" {
		page.LayoutSlug = &layoutData.Slug
	}

	err := gorm.G[models.Page](database.DBConn).Create(context.Background(), page)
	if err != nil {
		return c.Status(500).SendString("Error creating page: " + err.Error())
	}

	return renderPageTable(c)
}

func renderPageTable(c fiber.Ctx) error {
	pageTableData := buildPageTableData()
	return handlerutils.RenderNode(c, admincomponents.DataTable(pageTableData))
}

func buildPageTableData() *admincomponents.TableData {
	pages, err := gorm.G[models.Page](database.DBConn).Joins(clause.LeftJoin.Association("ComponentInstance"), nil).Order("title ASC").Find(context.Background())
	if err != nil {
		return &admincomponents.TableData{}
	}

	// Build table data
	pageTableData := &admincomponents.TableData{
		Columns: []admincomponents.TableColumn{
			{Name: "slug", Label: "Slug"},
			{Name: "title", Label: "Title"},
			{Name: "component", Label: "Component"},
		},
		Rows:             []admincomponents.TableRow{},
		Editable:         true,
		Deletable:        true,
		EditRoute:        "/admin/pages/edit",
		DeleteRoute:      "/admin/pages",
		IDField:          "id",
		RefreshTarget:    "#data-table-container",
		DeleteConfirmMsg: "Are you sure you want to delete this table? This action cannot be undone.",
	}

	for _, page := range pages {
		// componentName := ""
		// if page.ComponentInstanceSlug != nil {
		// 	component, err := gorm.G[models.ComponentInstance](database.DBConn).Where("slug = ?", *page.ComponentInstanceSlug).First(context.Background())
		// 	if err == nil {
		// 		componentName = component.Name
		// 	}
		// }
		componentName := "None"
		if page.Layout != nil {
			componentName = page.Layout.Title
		}
		pageTableData.Rows = append(pageTableData.Rows, admincomponents.TableRow{
			Values: map[string]string{
				"id":        page.ID.String(),
				"slug":      page.Slug,
				"title":     page.Title,
				"component": componentName,
			},
		})
	}

	return pageTableData
}
