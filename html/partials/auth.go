package htmlpartials

import (
	"github.com/invertedbit/gms/models"
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
				html.Name("username"),
			),
			html.Input(
				html.Class("input input-accent my-2 block w-full"),
				html.Type("password"),
				html.Placeholder("Password"),
				html.Name("password"),
			),
			html.Button(
				html.Class("btn btn-accent mt-2"),
				gomponents.Text("Login"),
			),
		),
	)
}

func ProfileMenu(user *models.User) gomponents.Node {
	var content gomponents.Node
	if user == nil {
		content = html.A(
			html.Class("btn"),
			gomponents.Text("Login"),
			html.Href("/auth/login"),
		)
	} else {
		content = html.Div(
			html.Span(
				gomponents.Text("Hello, "+user.Email),
				html.Class("mr-2"),
			),
			html.A(
				html.Class("btn btn-ghost"),
				html.Href("/auth/logout"),
				gomponents.Text("Logout"),
			),
		)
	}

	return content
}
