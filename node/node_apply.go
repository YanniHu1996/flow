/**
 * @Author: Huyantian
 * @Date: 2021/3/7 下午1:05
 */

package node

import "flow/core"

const Apply core.NodeType = 0

func NewApply(name string, applier core.Handler) *core.Node {
	return &core.Node{
		Name:    name,
		Type:    Apply,
		Status:  core.Ready,
		Handler: applier,
		OnEnter: func(self *core.Node) {
			self.Flow.AddReady(self)
			self.Status = core.Ready
		},
		OnLeave: func(self *core.Node) {
			self.Flow.RemoveReady(self)
			self.Status = core.Complete
		},
	}
}
