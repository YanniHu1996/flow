package flow

import (
	"flow/core"
	"flow/node"
)

func NewFlow(nodes ...*core.Node) *core.Flow {
	f := &core.Flow{Status: core.Running}
	if len(nodes) == 0 {
		return f
	}
	nodes = append(nodes, node.NewEnd())
	f.RootNode = nodes[0]
	main := node.NewSerial("主流程", nodes...)

	core.SetFlow(f, main)
	core.SetFlow(f, nodes...)
	core.Link(nodes)
	core.SetParent(main, nodes)
	return f
}
