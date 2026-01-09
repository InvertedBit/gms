package html

import (
	"fmt"
	"io"

	"github.com/invertedbit/gms/viewmodels"
	"maragu.dev/gomponents"
	htmx "maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/html"
)

// AdminLayout represents the admin panel layout
type AdminLayout struct {
	AdminLayoutViewModel *viewmodels.AdminLayoutViewModel
}

// IsOOB returns whether the layout should be rendered out-of-band
func (a *AdminLayout) IsOOB() bool {
	layoutType := a.AdminLayoutViewModel.LayoutType
	return layoutType == viewmodels.LayoutPartialOnly || layoutType == viewmodels.LayoutPartialWithFragments
}

// GetAdminHeader renders the minimal admin header with breadcrumbs, title, and action buttons
func (a *AdminLayout) GetAdminHeader() gomponents.Node {
	return html.Header(
		html.ID("admin-header"),
		gomponents.If(a.IsOOB(), htmx.SwapOOB("true")),
		html.Class("bg-base-100 border-b border-base-300 px-6 py-4"),
		html.Div(
			html.Class("flex items-center justify-between"),
			// Left side: Breadcrumbs and Title
			html.Div(
				html.Class("flex-1"),
				// Breadcrumbs
				gomponents.If(len(a.AdminLayoutViewModel.Breadcrumbs) > 0,
					html.Div(
						html.Class("text-sm breadcrumbs mb-1"),
						html.Ul(
							gomponents.Map(a.AdminLayoutViewModel.Breadcrumbs, func(breadcrumb viewmodels.Breadcrumb) gomponents.Node {
								return html.Li(
									gomponents.If(breadcrumb.Link != "",
										html.A(
											html.Href(breadcrumb.Link),
											gomponents.Text(breadcrumb.Label),
										),
									),
									gomponents.If(breadcrumb.Link == "",
										gomponents.Text(breadcrumb.Label),
									),
								)
							}),
						),
					),
				),
				// Title
				html.H1(
					html.Class("text-2xl font-bold"),
					gomponents.Text(a.AdminLayoutViewModel.Title),
				),
			),
			// Right side: Action buttons
			gomponents.If(len(a.AdminLayoutViewModel.ActionButtons) > 0,
				html.Div(
					html.Class("flex gap-2"),
					gomponents.Map(a.AdminLayoutViewModel.ActionButtons, func(button viewmodels.ActionButton) gomponents.Node {
						btnClass := "btn btn-sm"
						if button.Primary {
							btnClass = "btn btn-sm btn-primary"
						}
						return html.A(
							html.Class(btnClass),
							html.Href(button.Link),
							gomponents.If(button.Icon != "",
								html.I(html.Class(button.Icon+" mr-1")),
							),
							gomponents.Text(button.Label),
						)
					}),
				),
			),
		),
	)
}

// GetAdminNavigation renders the multi-level collapsible navigation menu
func (a *AdminLayout) GetAdminNavigation() gomponents.Node {
	return html.Aside(
		html.ID("admin-navigation"),
		gomponents.If(a.IsOOB(), htmx.SwapOOB("true")),
		html.Class("w-64 bg-base-200 border-r border-base-300 flex-shrink-0"),
		html.Div(
			html.Class("p-4"),
			html.Div(
				html.Class("font-bold text-xl mb-6 text-accent"),
				gomponents.Text("GMS Admin"),
			),
			html.Ul(
				html.Class("menu"),
				gomponents.Map(a.AdminLayoutViewModel.Navigation, func(item *viewmodels.AdminNavigationItem) gomponents.Node {
					return a.renderNavigationItem(item)
				}),
			),
		),
	)
}

// renderNavigationItem renders a single navigation item (recursive for children)
func (a *AdminLayout) renderNavigationItem(item *viewmodels.AdminNavigationItem) gomponents.Node {
	hasChildren := len(item.Children) > 0

	if hasChildren {
		return html.Li(
			html.Details(
				html.Summary(
					gomponents.If(item.Icon != "",
						html.I(html.Class(item.Icon+" mr-2")),
					),
					gomponents.Text(item.Label),
				),
				html.Ul(
					gomponents.Map(item.Children, func(child *viewmodels.AdminNavigationItem) gomponents.Node {
						return a.renderNavigationItem(child)
					}),
				),
			),
		)
	}

	return html.Li(
		html.A(
			html.Href(item.Link),
			gomponents.If(item.Icon != "",
				html.I(html.Class(item.Icon+" mr-2")),
			),
			gomponents.Text(item.Label),
		),
	)
}

// GetAdminFooter renders the small GMS-branded footer in accent color
func (a *AdminLayout) GetAdminFooter() gomponents.Node {
	return html.Footer(
		html.ID("admin-footer"),
		gomponents.If(a.IsOOB(), htmx.SwapOOB("true")),
		html.Class("bg-accent text-accent-content p-3 text-center text-sm"),
		html.Div(
			html.Span(
				html.Class("font-semibold"),
				gomponents.Text("GMS"),
			),
			gomponents.Text(" · Guild Management System"),
		),
	)
}

// GetAdminBody renders the complete admin layout body
func (a *AdminLayout) GetAdminBody(pageContent gomponents.Node) []gomponents.Node {
	return []gomponents.Node{
		html.Div(
			a.GetToastContainer(),
			html.Class("min-h-screen flex flex-col"),
			html.Div(
				html.Class("flex flex-1"),
				a.GetAdminNavigation(),
				html.Div(
					html.Class("flex-1 flex flex-col"),
					a.GetAdminHeader(),
					html.Main(
						html.Class("flex-1 p-6 bg-base-100"),
						pageContent,
					),
					a.GetAdminFooter(),
				),
			),
		),
	}
}

// GetToastContainer renders the toast notification container
func (a *AdminLayout) GetToastContainer() gomponents.Node {
	return html.Div(
		html.ID("toast-container"),
		gomponents.If(a.IsOOB(), htmx.SwapOOB("true")),
		html.Class("toast toast-top toast-end"),
		gomponents.Map(a.AdminLayoutViewModel.ToastViewModel.Messages, func(toast viewmodels.ToastMessage) gomponents.Node {
			return html.Div(
				html.Class(fmt.Sprintf("alert %s", toast.GetClassName())),
				html.Span(
					gomponents.Text(toast.Message),
				),
			)
		}),
	)
}

// Render renders the admin layout based on the layout type
func (a *AdminLayout) Render(w io.Writer, pageContent gomponents.Node) error {
	switch a.AdminLayoutViewModel.LayoutType {
	case viewmodels.LayoutFull:
		return a.RenderFullLayout(w, pageContent)
	case viewmodels.LayoutBodyOnly:
		return a.RenderBodyOnly(w, pageContent)
	case viewmodels.LayoutPartialOnly:
		return a.RenderPartialOnly(w, pageContent)
	default:
		return a.RenderFullLayout(w, pageContent)
	}
}

// RenderFullLayout renders the complete HTML document
func (a *AdminLayout) RenderFullLayout(w io.Writer, pageContent gomponents.Node) error {
	return html.Doctype(
		html.HTML(
			html.Lang("en"),
			html.Head(
				html.Meta(html.Charset("UTF-8")),
				html.Meta(
					html.Name("viewport"),
					html.Content("width=device-width, initial-scale=1"),
				),
				html.TitleEl(
					gomponents.Text(a.AdminLayoutViewModel.Page),
				),
				html.Link(
					html.Rel("stylesheet"),
					html.Href("/css/theme.css"),
				),
				html.Link(
					html.Rel("stylesheet"),
					html.Href("https://cdn.jsdelivr.net/npm/remixicon@4.5.0/fonts/remixicon.css"),
				),
				html.Script(
					html.Src("https://cdn.jsdelivr.net/npm/htmx.org@2.0.6/dist/htmx.min.js"),
				),
				html.Script(
					html.Src("https://unpkg.com/hyperscript.org@0.9.14"),
				),
				html.Script(
					html.Src("/js/wasm_exec.js"),
				),
				html.Script(
					html.Src("/js/wasm_init.js"),
				),
			),
			html.Body(
				a.GetAdminBody(pageContent)...,
			),
		),
	).Render(w)
}

// RenderBodyOnly renders just the body content
func (a *AdminLayout) RenderBodyOnly(w io.Writer, pageContent gomponents.Node) error {
	err := html.TitleEl(
		htmx.SwapOOB("true"),
		gomponents.Text(a.AdminLayoutViewModel.Page),
	).Render(w)
	if err != nil {
		return err
	}
	return html.Body(a.GetAdminBody(pageContent)...).Render(w)
}

// RenderPartialOnly renders just the partial content with OOB updates
func (a *AdminLayout) RenderPartialOnly(w io.Writer, pageContent gomponents.Node) error {
	err := html.TitleEl(
		gomponents.Text(a.AdminLayoutViewModel.Page),
	).Render(w)
	if err != nil {
		return err
	}
	err = a.GetAdminHeader().Render(w)
	if err != nil {
		return err
	}
	err = a.GetToastContainer().Render(w)
	if err != nil {
		return err
	}
	return pageContent.Render(w)
}
