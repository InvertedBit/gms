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
