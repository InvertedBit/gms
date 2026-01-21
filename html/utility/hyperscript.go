package utility

import "maragu.dev/gomponents"

func Hyperscript(script string) gomponents.Node {
	return gomponents.Attr("_", script)
}
