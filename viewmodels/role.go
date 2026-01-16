package viewmodels

import "github.com/invertedbit/gms/models"

type RoleFormViewModel struct {
	Role       *models.Role
	IsEdit     bool
	SubmitURL  string
	CancelURL  string
	FormErrors map[string]string
}

func NewRoleFormViewModel(role *models.Role, isEdit bool) *RoleFormViewModel {
	submitURL := "/admin/roles"
	if isEdit && role != nil {
		submitURL = "/admin/roles/" + role.ID.String()
	}

	return &RoleFormViewModel{
		Role:       role,
		IsEdit:     isEdit,
		SubmitURL:  submitURL,
		CancelURL:  "/admin/roles",
		FormErrors: make(map[string]string),
	}
}

func (vm *RoleFormViewModel) GetRoleName() string {
	if vm.Role == nil {
		return ""
	}
	return vm.Role.Name
}

func (vm *RoleFormViewModel) GetRoleSlug() string {
	if vm.Role == nil {
		return ""
	}
	return vm.Role.Slug
}

func (vm *RoleFormViewModel) GetRoleDescription() string {
	if vm.Role == nil {
		return ""
	}
	return vm.Role.Description
}

func (vm *RoleFormViewModel) GetFormError(field string) string {
	if err, exists := vm.FormErrors[field]; exists {
		return err
	}
	return ""
}
