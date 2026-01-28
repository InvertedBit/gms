package viewmodels

import "github.com/invertedbit/gms/models"

type PageFormViewModel struct {
	FormViewModel
	Page    *models.Page
	Layouts []models.Layout
	Title   string
	Slug    string
}

func NewPageFormViewModel(page *models.Page, layouts []models.Layout, isEdit bool) *PageFormViewModel {
	submitURL := "/admin/pages"
	if isEdit && page != nil {
		submitURL = "/admin/pages/" + page.Slug
	}

	return &PageFormViewModel{
		Page: page,
		FormViewModel: FormViewModel{
			IsEdit:     isEdit,
			SubmitURL:  submitURL,
			CancelURL:  "/admin/pages",
			FormErrors: make(map[string]string),
		},
		Layouts: layouts,
	}
}

func (vm *PageFormViewModel) GetPageTitle() string {
	if vm.Page == nil {
		return ""
	}
	return vm.Page.Title
}

func (vm *PageFormViewModel) GetPageSlug() string {
	if vm.Page == nil {
		return ""
	}
	return vm.Page.Slug
}

func (vm *PageFormViewModel) GetLayoutSlug() string {
	if vm.Page == nil || vm.Page.Layout == nil {
		return ""
	}
	return vm.Page.Layout.Slug
}

func (vm *PageFormViewModel) GetLayoutTitle() string {
	if vm.Page == nil || vm.Page.Layout == nil {
		return ""
	}
	return vm.Page.Layout.Title
}
