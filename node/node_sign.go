/**
 * @Author: Huyantian
 * @Date: 2021/3/7 下午1:13
 */

package node

import "flow/core"

const SignAudit core.NodeType = 2

func NewSignAudit(name string, nodes ...*core.Node) *core.Node {
	n := &core.Node{
		Name:     name,
		Type:     SignAudit,
		Status:   core.Future,
		Children: nodes,
		OnEnter: func(self *core.Node) {
			self.Status = core.Waiting
			for _, child := range self.Children {
				child.Enter()
			}
		},
		AfterChildLeave: func(self, _ *core.Node) {
			var count int
			for _, child := range self.Children {
				if child.Status == core.Complete {
					count++
				}
			}
			if count >= len(self.Children) {
				self.Status = core.Complete
				self.Leave()
			}
		},
	}
	return n
}
