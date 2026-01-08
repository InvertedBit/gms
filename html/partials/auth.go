package htmlpartials

import (
	"maragu.dev/gomponents"
	htmx "maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/html"
)

func LoginForm() gomponents.Node {
	return html.Div(
		html.H2(
			html.Class("text-2xl mb-2"),
			gomponents.Text("Login"),
		),
		html.Form(
			htmx.Post("/auth/login"),
			html.Input(
				html.Class("input input-accent block"),
				html.Type("text"),
				html.Placeholder("Username"),
			),
			html.Input(
				html.Class("input input-accent block"),
				html.Type("password"),
				html.Placeholder("Password"),
			),
			html.Button(
				html.Class("btn btn-accent mt-4"),
				gomponents.Text("Login"),
			),
		),
	)
}
