/**
 * @Author: Huyantian
 * @Date: 2021/3/7 下午6:58
 */

package node

import "flow/core"

const EndNode core.NodeType = 5

func NewEnd() *core.Node {
	return &core.Node{
		Type:   EndNode,
		Status: core.Future,
		OnEnter: func(self *core.Node) {
			self.Flow.Status = core.End
		},
	}
}
