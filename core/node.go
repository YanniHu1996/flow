/**
 * @Author: Huyantian
 * @Date: 2021/3/7 上午11:37
 */

package core

type NodeType int

type NodeStatus int

const (
	Ready NodeStatus = iota
	Complete
	Future
	Waiting
	Skip
)

type Node struct {
	Flow *Flow
	Name string

	Type   NodeType
	Status NodeStatus

	Precursor *Node
	Successor *Node

	Parent   *Node
	Children []*Node

	BeforeChildEnter, AfterChildEnter func(self, child *Node)
	BeforeChildLeave, AfterChildLeave func(self, child *Node)

	Handler Handler

	OnEnter, OnLeave func(self *Node)
}

func (n *Node) Enter() {
	if p := n.Parent; p != nil && p.BeforeChildEnter != nil {
		n.Parent.BeforeChildEnter(n.Parent, n)
	}
	if n.OnEnter != nil {
		n.OnEnter(n)
	}
	if p := n.Parent; p != nil && p.AfterChildEnter != nil {
		n.Parent.AfterChildEnter(n.Parent, n)
	}
}

func (n *Node) Leave() {
	if p := n.Parent; p != nil && p.BeforeChildLeave != nil {
		n.Parent.BeforeChildLeave(n.Parent, n)
	}
	if n.OnLeave != nil {
		n.OnLeave(n)
	}
	if p := n.Parent; p != nil && p.AfterChildLeave != nil {
		n.Parent.AfterChildLeave(n.Parent, n)
	}
}

func (n *Node) Complete() {
	n.Status = Complete
}

func Link(nodes []*Node) {
	if len(nodes) <= 1 {
		return
	}
	for i, node := range nodes {
		Link(node.Children)
		if i != 0 {
			nodes[i].Precursor = nodes[i-1]
		}
		if i != len(nodes)-1 {
			nodes[i].Successor = nodes[i+1]
		}
	}
}

func SetParent(parent *Node, children []*Node) {
	parent.Children = children
	for _, child := range children {
		SetParent(child, child.Children)
		child.Parent = parent
	}
}

func SetFlow(flow *Flow, nodes ...*Node) {
	for _, node := range nodes {
		SetFlow(flow, node.Children...)
		node.Flow = flow
	}
}
