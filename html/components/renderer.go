package components

import (
	"github.com/invertedbit/gms/viewmodels"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

type RenderFunc func(*viewmodels.ComponentViewModel) gomponents.Node

type Component struct {
	Name   string
	Render RenderFunc
}

type Renderer struct {
	Components map[string]Component
}

func NewRenderer() *Renderer {
	return &Renderer{
		Components: map[string]Component{},
	}
}

func (r *Renderer) RenderComponent(vm *viewmodels.ComponentViewModel) gomponents.Node {
	if component, exists := r.Components[vm.Name]; exists {
		return component.Render(vm)
	} else {
		return html.Div(
			gomponents.Text("Unknown Component"),
		)
	}
}
