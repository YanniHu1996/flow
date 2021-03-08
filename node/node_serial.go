/**
 * @Author: Huyantian
 * @Date: 2021/3/7 下午3:33
 */

package node

import "flow/core"

const Serial core.NodeType = 3

func NewSerial(name string, nodes ...*core.Node) *core.Node {
	n := &core.Node{
		Name:     name,
		Type:     Serial,
		Status:   core.Future,
		Children: nodes,
		OnEnter: func(self *core.Node) {
			self.Status = core.Waiting
			if len(self.Children) > 0 {
				child := self.Children[0]
				child.Enter()
			}
		},
		AfterChildLeave: func(self, child *core.Node) {
			if child.Successor == nil {
				self.Status = core.Complete
				self.Leave()
				return
			}
			child.Successor.Enter()
		},
	}
	return n
}
