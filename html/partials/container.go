package htmlpartials

import (
	"github.com/invertedbit/gms/html/utility"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

type ContainerType int

const (
	ContainerTypeDefault ContainerType = iota
	ContainerTypeBordered
	ContainerTypeBackgroundBase
	ContainerTypeBackgroundPrimary
)

func Container(classBuilder *utility.ClassBuilder, children ...gomponents.Node) gomponents.Node {
	if classBuilder == nil {
		classBuilder = &utility.ClassBuilder{}
	}

	classBuilder.Class("container w-lg mx-auto rounded-lg bg-base-300 p-5")

	params := utility.MergeParams(children, classBuilder.GetNode())

	return html.Div(
		params...,
	)
}
