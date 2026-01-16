package components

import (
	"maragu.dev/gomponents"
	htmx "maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/html"
)

// Modal renders a modal overlay for forms and content
func Modal(id string, title string, content gomponents.Node) gomponents.Node {
	return html.Div(
		html.ID(id),
		html.Class("modal modal-open"),
		html.Div(
			html.Class("modal-box max-w-2xl"),
			html.H3(
				html.Class("font-bold text-lg mb-4"),
				gomponents.Text(title),
			),
			html.Div(
				html.Class("modal-content"),
				content,
			),
		),
		html.Form(
			html.Method("dialog"),
			html.Class("modal-backdrop"),
			html.Button(
				html.Type("button"),
			),
		),
	)
}

// ModalContainer renders the container for modals
func ModalContainer(oob bool) gomponents.Node {
	return html.Div(
		html.ID("modal-container"),
		gomponents.If(oob,
			htmx.SwapOOB("true"),
		),
	)
}
