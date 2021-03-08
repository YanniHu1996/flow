/**
 * @Author: Huyantian
 * @Date: 2021/3/7 上午11:37
 */

package core

type Status int

const (
	Running Status = iota
	End
)

type Flow struct {
	Status   Status
	RootNode *Node

	ReadyNodes Nodes
}

func (f *Flow) AddReady(node *Node) {
	f.ReadyNodes.Add(node)
}

func (f *Flow) RemoveReady(node *Node) {
	f.ReadyNodes.Remove(node)
}

func (f *Flow) FirstReady() *Node {
	return f.ReadyNodes.First()
}

func (f *Flow) IsEnd() bool {
	if f.Status == End || f.ReadyNodes.Size() == 0 {
		return true
	}
	return false
}

func (f *Flow) Start() {
	f.RootNode.Enter()
}
