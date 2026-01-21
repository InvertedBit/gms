package adminviews

import (
	"github.com/invertedbit/gms/html/components"
	"github.com/invertedbit/gms/models"
	"github.com/invertedbit/gms/viewmodels"
	"maragu.dev/gomponents"
	htmx "maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/html"
)

func UserListPage(userTableData *components.TableData) gomponents.Node {
	return html.Div(
		html.ID("users-list"),
		components.ModalContainer(false),
		components.DataTable(userTableData),
	)
}

func UserFormModal(vm *viewmodels.UserFormViewModel) gomponents.Node {
	if vm == nil {
		return gomponents.Text("Error: ViewModel is nil")
	}
	title := "Create User"
	if vm.IsEdit {
		title = "Edit User"
	}

	return components.Modal("user-form-modal", title,
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

			// Email field
			html.Div(
				html.Class("form-control"),
				html.Label(
					html.Class("label"),
					html.Span(
						html.Class("label-text"),
						gomponents.Text("Email"),
					),
				),
				html.Input(
					html.Type("email"),
					html.Name("email"),
					html.Class("input input-bordered"),
					html.Required(),
					html.Value(vm.GetUserEmail()),
				),
				gomponents.If(vm.GetFormError("email") != "",
					html.Label(
						html.Class("label"),
						html.Span(
							html.Class("label-text-alt text-error"),
							gomponents.Text(vm.GetFormError("email")),
						),
					),
				),
			),

			// Password field
			html.Div(
				html.Class("form-control"),
				html.Label(
					html.Class("label"),
					html.Span(
						html.Class("label-text"),
						gomponents.If(vm.IsEdit,
							gomponents.Text("Password (leave blank to keep current)"),
						),
						gomponents.If(!vm.IsEdit,
							gomponents.Text("Password"),
						),
					),
				),
				html.Input(
					html.Type("password"),
					html.Name("password"),
					html.Class("input input-bordered"),
					gomponents.If(!vm.IsEdit,
						html.Required(),
					),
				),
				gomponents.If(vm.GetFormError("password") != "",
					html.Label(
						html.Class("label"),
						html.Span(
							html.Class("label-text-alt text-error"),
							gomponents.Text(vm.GetFormError("password")),
						),
					),
				),
			),

			// Role field
			html.Div(
				html.Class("form-control"),
				html.Label(
					html.Class("label"),
					html.Span(
						html.Class("label-text"),
						gomponents.Text("Role"),
					),
				),
				html.Select(
					html.Name("role_slug"),
					html.Class("select select-bordered"),
					html.Option(
						html.Value(""),
						gomponents.Text("Select a role"),
					),
					gomponents.Map(vm.Roles, func(role models.Role) gomponents.Node {
						return html.Option(
							html.Value(role.Slug),
							gomponents.If(vm.GetUserRoleSlug() == role.Slug,
								html.Selected(),
							),
							gomponents.Text(role.Name),
						)
					}),
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
					htmx.Get("/admin/users"),
					htmx.Target("#users-list"),
					gomponents.Text("Cancel"),
				),
			),
		),
	)
}
