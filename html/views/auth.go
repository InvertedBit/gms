package htmlviews

import (
	htmlpartials "github.com/invertedbit/gms/html/partials"
	"github.com/invertedbit/gms/html/utility"
	"maragu.dev/gomponents"
)

func LoginPage() gomponents.Node {
	return htmlpartials.Container(
		utility.GetClassBuilder("border-2 w-2xl border-accent"),
		htmlpartials.LoginForm(),
	)
}
