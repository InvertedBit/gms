package adminviews

import (
	admincomponents "github.com/invertedbit/gms/html/components/admin"
	"github.com/invertedbit/gms/models"
	"github.com/invertedbit/gms/viewmodels"
	"maragu.dev/gomponents"
	htmx "maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/html"
)

func PageListPage(pageTableData *admincomponents.TableData) gomponents.Node {
	return html.Div(
		html.ID("pages-list"),
		admincomponents.ModalContainer(false),
		admincomponents.DataTable(pageTableData),
	)
}

func PageFormModal(vm *viewmodels.PageFormViewModel) gomponents.Node {
	if vm == nil {
		return gomponents.Text("Error: ViewModel is nil")
	}
	title := "Create Page"
	if vm.IsEdit {
		title = "Edit Page"
	}

	return admincomponents.Modal("page-form-modal", title,
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

			// Title field
			html.Div(
				html.Class("form-control"),
				html.Label(
					html.Class("label"),
					html.Span(
						html.Class("label-text"),
						gomponents.Text("Title"),
					),
				),
				html.Input(
					html.Type("text"),
					html.Name("title"),
					html.Class("input input-bordered"),
					html.Required(),
					html.Value(vm.GetPageTitle()),
				),
				gomponents.If(vm.GetFormError("title") != "",
					html.Label(
						html.Class("label"),
						html.Span(
							html.Class("label-text-alt text-error"),
							gomponents.Text(vm.GetFormError("title")),
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
						gomponents.If(vm.IsEdit,
							gomponents.Text("Slug"),
						),
						gomponents.If(!vm.IsEdit,
							gomponents.Text("Slug"),
						),
					),
				),
				html.Input(
					html.Type("text"),
					html.Name("slug"),
					html.Class("input input-bordered"),
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

			// Layout field
			html.Div(
				html.Class("form-control"),
				html.Label(
					html.Class("label"),
					html.Span(
						html.Class("label-text"),
						gomponents.Text("Layout"),
					),
				),
				html.Select(
					html.Name("layout"),
					html.Class("select select-bordered"),
					html.Option(
						html.Value("/none"),
						gomponents.Text("None"),
					),
					gomponents.Map(vm.Layouts, func(layout models.Layout) gomponents.Node {
						return html.Option(
							html.Value(layout.Slug),
							gomponents.If(vm.GetLayoutSlug() == layout.Slug,
								html.Selected(),
							),
							gomponents.Text(layout.Title),
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
					htmx.Get("/admin/pages"),
					htmx.Target("#pages-list"),
					gomponents.Text("Cancel"),
				),
			),
		),
	)
}
