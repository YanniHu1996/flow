/**
 * @Author: Huyantian
 * @Date: 2021/3/7 下午1:05
 */

package node

import "flow/core"

const Audit core.NodeType = 1

func NewNormalAudit(name string, approver core.Handler) *core.Node {
	return &core.Node{
		Name:    name,
		Type:    Audit,
		Status:  core.Future,
		Handler: approver,
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
