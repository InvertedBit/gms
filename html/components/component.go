package components

import (
	"github.com/invertedbit/gms/viewmodels"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func RenderComponent(vm *viewmodels.ComponentViewModel) gomponents.Node {
	switch vm.Name {
	case "Container":
		return RenderContainerComponent(vm.ComponentInstance)
	default:
		return html.Div(
			gomponents.Text("Unknown Component"),
		)
	}
}

func GetComponentList() map[string]string {
	return map[string]string{
		"container":     "Container",
		"stack":         "Stack",
		"row":           "Row",
		"column":        "Column",
		"table":         "Table",
		"list":          "List",
		"text":          "Text",
		"image":         "Image",
		"button":        "Button",
		"link":          "Link",
		"form":          "Form",
		"input":         "Input",
		"checkbox":      "Checkbox",
		"radio":         "Radio",
		"select":        "Select",
		"modal":         "Modal",
		"navbar":        "Navbar",
		"footer":        "Footer",
		"header":        "Header",
		"card":          "Card",
		"carousel":      "Carousel",
		"carouselitem":  "Carousel Item",
		"accordion":     "Accordion",
		"accordionitem": "Accordion Item",
		"tabs":          "Tabs",
		// Add more components here as needed
	}
}
