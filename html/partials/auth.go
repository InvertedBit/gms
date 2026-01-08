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
			html.Div(
				html.Class("login-error-message hidden border-2 p-2 my-2 border-error rounded-lg"),
			),
			html.Input(
				html.Class("input input-accent my-2 block w-full"),
				html.Type("text"),
				html.Placeholder("Username"),
			),
			html.Input(
				html.Class("input input-accent my-2 block w-full"),
				html.Type("password"),
				html.Placeholder("Password"),
			),
			html.Button(
				html.Class("btn btn-accent mt-2"),
				gomponents.Text("Login"),
			),
		),
	)
}
