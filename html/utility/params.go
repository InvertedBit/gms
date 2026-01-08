package utility

import "maragu.dev/gomponents"

func MergeParams(params []gomponents.Node, newParams ...gomponents.Node) []gomponents.Node {
	params = append(params, newParams...)

	return params
}
