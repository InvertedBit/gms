package adminviews

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func DashboardPage() gomponents.Node {
	return html.Div(
		html.H1(
			html.Class("text-3xl font-bold mb-4"),
			gomponents.Text("Admin Dashboard"),
		),
		html.P(
			gomponents.Text("Welcome to the admin dashboard! Here you can manage the application settings and view important metrics."),
		),
	)
}
