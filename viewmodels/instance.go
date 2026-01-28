package viewmodels

import "github.com/invertedbit/gms/models"

type InstanceFormViewModel struct {
	FormViewModel
	Instance *models.ComponentInstance
	Title    string
	Slug     string
}

func NewInstanceFormViewModel(instance *models.ComponentInstance, isEdit bool) *InstanceFormViewModel {
	submitURL := "/admin/instances"
	if isEdit {
		submitURL = "/admin/instances/" + instance.Slug
	}

	return &InstanceFormViewModel{
		Instance: instance,
		FormViewModel: FormViewModel{
			IsEdit:     isEdit,
			SubmitURL:  submitURL,
			CancelURL:  "/admin/instances",
			FormErrors: make(map[string]string),
		},
	}
}

func (vm *InstanceFormViewModel) GetInstanceName() string {
	if vm.Instance == nil {
		return ""
	}
	return vm.Instance.Name
}

func (vm *InstanceFormViewModel) GetInstanceSlug() string {
	if vm.Instance == nil {
		return ""
	}
	return vm.Instance.Slug
}
