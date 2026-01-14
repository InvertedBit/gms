package adminviews

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func DashboardPage() gomponents.Node {
	return html.Div(
		html.P(
			gomponents.Text("Welcome to the admin dashboard! Here you can manage the application settings and view important metrics."),
		),
	)
}
