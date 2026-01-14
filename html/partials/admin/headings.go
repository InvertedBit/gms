package adminpartials

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func PageHeading(title string) gomponents.Node {
	return html.H1(
		html.Class("text-3xl font-bold text-accent mt-2"),
		gomponents.Text(title),
	)
}
