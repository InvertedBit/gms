package components

import (
	"github.com/invertedbit/gms-plugins/components"
	"github.com/invertedbit/gms-plugins/plugins"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

var ComponentRenderer *Renderer

type Renderer struct {
	Components map[string]components.Component
}

func NewRenderer() *Renderer {
	return &Renderer{
		Components: map[string]components.Component{
			"container": {
				Name:        "Container",
				Description: "A basic container (div) component",
				Render:      RenderContainerComponent,
				Children:    []components.Component{},
			},
			// Add more components here as needed
		},
	}
}

func (r *Renderer) LoadComponentsFromPlugin(plugin plugins.Plugin) {
	if plugin.Components == nil {
		return
	}
	for key, component := range plugin.Components {
		r.Components[key] = component
	}
}

func (r *Renderer) PrintLoadedComponents() {
	for name, component := range r.Components {
		println("Loaded component:", name, "-", component.Name)
	}
}

func (r *Renderer) RenderComponent(vm *components.ComponentViewModel) gomponents.Node {
	if component, exists := r.Components[vm.Name]; exists {
		return component.Render(vm)
	} else {
		return html.Div(
			gomponents.Text("Unknown Component"),
		)
	}
}
