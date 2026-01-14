package adminviews

import (
	"github.com/invertedbit/gms/html/components"
	"github.com/invertedbit/gms/viewmodels"
	"maragu.dev/gomponents"
	htmx "maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/html"
)

func RoleListPage(roleTableData *components.TableData) gomponents.Node {
	return html.Div(
		components.ModalContainer(),
		components.DataTable(roleTableData),
	)
}

func RoleFormModal(vm *viewmodels.RoleFormViewModel) gomponents.Node {
	title := "Create Role"
	if vm.IsEdit {
		title = "Edit Role"
	}

	return components.Modal("role-form-modal", title,
		html.FormEl(
			html.Action(vm.SubmitURL),
			gomponents.If(vm.IsEdit,
				htmx.Put(vm.SubmitURL),
			),
			gomponents.If(!vm.IsEdit,
				htmx.Post(vm.SubmitURL),
			),
			htmx.Target("#data-table-container"),
			htmx.Swap("outerHTML"),
			html.Class("space-y-4"),

			// Name field
			html.Div(
				html.Class("form-control"),
				html.Label(
					html.Class("label"),
					html.Span(
						html.Class("label-text"),
						gomponents.Text("Name"),
					),
				),
				html.Input(
					html.Type("text"),
					html.Name("name"),
					html.Class("input input-bordered"),
					html.Required(),
					gomponents.If(vm.Role != nil && vm.Role.Name != "",
						html.Value(vm.Role.Name),
					),
				),
				gomponents.If(vm.FormErrors["name"] != "",
					html.Label(
						html.Class("label"),
						html.Span(
							html.Class("label-text-alt text-error"),
							gomponents.Text(vm.FormErrors["name"]),
						),
					),
				),
			),

			// Description field
			html.Div(
				html.Class("form-control"),
				html.Label(
					html.Class("label"),
					html.Span(
						html.Class("label-text"),
						gomponents.Text("Description"),
					),
				),
				html.Textarea(
					html.Name("description"),
					html.Class("textarea textarea-bordered"),
					gomponents.If(vm.Role != nil && vm.Role.Description != "",
						gomponents.Text(vm.Role.Description),
					),
				),
			),

			// Action buttons
			html.Div(
				html.Class("modal-action"),
				html.Button(
					html.Type("submit"),
					html.Class("btn btn-primary"),
					gomponents.Text("Save"),
				),
				html.Button(
					html.Type("button"),
					html.Class("btn"),
					htmx.Get(""),
					htmx.Target("#modal-container"),
					htmx.Swap("innerHTML"),
					gomponents.Text("Cancel"),
				),
			),
		),
	)
}
