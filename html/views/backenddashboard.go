package htmlviews

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func BackendDashboardPage() gomponents.Node {
	return html.Div(
		html.H1(
			gomponents.Text("Login to GMS!"),
		),
	)
}
