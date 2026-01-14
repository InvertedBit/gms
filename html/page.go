package html

import (
	"io"
	"net/http"

	"github.com/invertedbit/gms/viewmodels"
	"maragu.dev/gomponents"
)

type PageRenderer interface {
	Render(w io.Writer) error
}

// type Layout interface {
// 	Render(w io.Writer, pageContent gomponents.Node) error
// }

type Page struct {
	Title           string
	PageContent     gomponents.Node
	LayoutViewModel *viewmodels.LayoutViewModel
}

func (p *Page) Render(w io.Writer) error {
	layout := Layout{
		LayoutViewModel: p.LayoutViewModel,
	}
	return layout.Render(w, p.PageContent)
}

func (p Page) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.Render(w)
}

// AdminPage represents a page using the admin layout
type AdminPage struct {
	Title                string
	PageContent          gomponents.Node
	AdminLayoutViewModel *viewmodels.AdminLayoutViewModel
}

func (p *AdminPage) Render(w io.Writer) error {
	layout := AdminLayout{
		AdminLayoutViewModel: p.AdminLayoutViewModel,
	}
	return layout.Render(w, p.PageContent)
}

func (p AdminPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.Render(w)
}
