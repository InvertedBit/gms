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
		components.ModalContainer(false),
		components.DataTable(roleTableData),
	)
}

func RoleFormModal(vm *viewmodels.RoleFormViewModel) gomponents.Node {
	if vm == nil {
		return gomponents.Text("Error: ViewModel is nil")
	}
	title := "Create Role"
	if vm.IsEdit {
		title = "Edit Role"
	}

	return components.Modal("role-form-modal", title,
		html.Form(
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
					html.Value(vm.GetRoleName()),
				),
				gomponents.If(vm.GetFormError("name") != "",
					html.Label(
						html.Class("label"),
						html.Span(
							html.Class("label-text-alt text-error"),
							gomponents.Text(vm.GetFormError("name")),
						),
					),
				),
			),

			// Slug field
			html.Div(
				html.Class("form-control"),
				html.Label(
					html.Class("label"),
					html.Span(
						html.Class("label-text"),
						gomponents.Text("Slug"),
					),
				),
				html.Input(
					html.Type("text"),
					html.Name("slug"),
					html.Class("input input-bordered"),
					html.Required(),
					html.Value(vm.GetRoleSlug()),
				),
				gomponents.If(vm.GetFormError("slug") != "",
					html.Label(
						html.Class("label"),
						html.Span(
							html.Class("label-text-alt text-error"),
							gomponents.Text(vm.GetFormError("slug")),
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
					gomponents.Text(vm.GetRoleDescription()),
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
