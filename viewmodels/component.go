package viewmodels

import "github.com/invertedbit/gms/models"

type ComponentViewModel struct {
	IsEdit            bool
	SubmitURL         string
	CancelURL         string
	FormErrors        map[string]string
	ComponentInstance *models.ComponentInstance
	Name              string
}

func NewComponentViewModel(instance *models.ComponentInstance, name string, isEdit bool) *ComponentViewModel {
	submitURL := "/admin/instances/"
	if isEdit {
		submitURL = "/admin/instances/" + instance.Slug
	}
	cancelURL := "/admin/instances"

	return &ComponentViewModel{
		IsEdit:            isEdit,
		SubmitURL:         submitURL,
		CancelURL:         cancelURL,
		FormErrors:        make(map[string]string),
		ComponentInstance: instance,
		Name:              name,
	}
}

func (vm *ComponentViewModel) GetFormError(field string) string {
	if err, exists := vm.FormErrors[field]; exists {
		return err
	}
	return ""
}
