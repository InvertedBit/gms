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
