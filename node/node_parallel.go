/**
 * @Author: Huyantian
 * @Date: 2021/3/7 下午6:52
 */

package node

import "flow/core"

const Parallel core.NodeType = 4

func NewParallel(name string, nodes ...*core.Node) *core.Node {
	return &core.Node{
		Name:     name,
		Type:     Parallel,
		Status:   core.Future,
		Children: nodes,
		OnEnter: func(self *core.Node) {
			self.Status = core.Waiting
			for _, child := range self.Children {
				child.Enter()
			}
		},
		AfterChildLeave: func(self, _ *core.Node) {
			self.Status = core.Complete
			self.Leave()
		},
	}
}
