package viewmodels

import "github.com/invertedbit/gms/models"

type UserFormViewModel struct {
	User       *models.User
	Roles      []models.Role
	IsEdit     bool
	SubmitURL  string
	CancelURL  string
	FormErrors map[string]string
}

func NewUserFormViewModel(user *models.User, roles []models.Role, isEdit bool) *UserFormViewModel {
	submitURL := "/admin/users"
	if isEdit && user != nil {
		submitURL = "/admin/users/" + user.ID.String()
	}

	return &UserFormViewModel{
		User:       user,
		Roles:      roles,
		IsEdit:     isEdit,
		SubmitURL:  submitURL,
		CancelURL:  "/admin/users",
		FormErrors: make(map[string]string),
	}
}

func (vm *UserFormViewModel) GetFormError(field string) string {
	if err, exists := vm.FormErrors[field]; exists {
		return err
	}
	return ""
}

func (vm *UserFormViewModel) GetUserEmail() string {
	if vm.User == nil {
		return ""
	}
	return vm.User.Email
}

func (vm *UserFormViewModel) GetUserRoleSlug() string {
	if vm.User == nil {
		return ""
	}
	return vm.User.RoleSlug
}

func (vm *UserFormViewModel) GetUserRoleName() string {
	roleSlug := vm.GetUserRoleSlug()
	if roleSlug == "" {
		return ""
	}
	for _, role := range vm.Roles {
		if role.Slug == roleSlug {
			return role.Name
		}
	}
	return ""
}
