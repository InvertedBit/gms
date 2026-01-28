package adminviews

import (
	admincomponents "github.com/invertedbit/gms/html/components/admin"
	"github.com/invertedbit/gms/viewmodels"
	"maragu.dev/gomponents"
	htmx "maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/html"
)

func InstanceListPage(instanceTableData *admincomponents.TableData) gomponents.Node {
	return html.Div(
		html.ID("instances-list"),
		admincomponents.ModalContainer(false),
		admincomponents.DataTable(instanceTableData),
	)
}

func InstanceFormModal(vm *viewmodels.InstanceFormViewModel) gomponents.Node {
	if vm == nil {
		return gomponents.Text("Error: ViewModel is nil")
	}
	title := "Create Instance"
	if vm.IsEdit {
		title = "Edit Instance"
	}

	return admincomponents.Modal("instance-form-modal", title,
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
					html.Value(vm.GetInstanceName()),
				),
				gomponents.If(vm.GetFormError("name") != "",
					html.Label(
						html.Class("label"),
						html.Span(
							html.Class("label-text text-error"),
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
					html.Value(vm.GetInstanceSlug()),
				),
				gomponents.If(vm.GetFormError("slug") != "",
					html.Label(
						html.Class("label"),
						html.Span(
							html.Class("label-text text-error"),
							gomponents.Text(vm.GetFormError("slug")),
						),
					),
				),
			),
		),
	)
}
