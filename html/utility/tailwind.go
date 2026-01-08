package utility

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

type ClassBuilder struct {
	classes []string
}

func GetClassBuilder(classes string) *ClassBuilder {
	return &ClassBuilder{
		classes: []string{classes},
	}
}

func (cb *ClassBuilder) Class(classes string) {
	if cb.classes == nil {
		cb.classes = []string{}
	}
	cb.classes = append(cb.classes, classes)
}

func (cb *ClassBuilder) GetNode() gomponents.Node {
	classString := ""
	for i := len(cb.classes) - 1; i >= 0; i-- {
		classString += cb.classes[i]
		if i > 0 {
			classString += " "
		}
	}
	return html.Class(classString)
}
