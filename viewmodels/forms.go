package viewmodels

type FormViewModel struct {
	IsEdit     bool
	SubmitURL  string
	CancelURL  string
	FormErrors map[string]string
}

func (vm *FormViewModel) GetFormError(field string) string {
	if err, exists := vm.FormErrors[field]; exists {
		return err
	}
	return ""
}
