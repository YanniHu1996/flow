/**
 * @Author: Huyantian
 * @Date: 2021/3/8 下午3:57
 */

package core

type Nodes []*Node

func (ns *Nodes) Size() int {
	return len(*ns)
}

func (ns *Nodes) Add(node *Node) {
	*ns = append(*ns, node)
}

func (ns *Nodes) Remove(node *Node) {
	i := ns.Index(node)
	if i == -1 {
		return
	}
	*ns = append((*ns)[:i], (*ns)[i:]...)
}

func (ns Nodes) Index(node *Node) int {
	for i, n := range ns {
		if n == node {
			return i
		}
	}
	return -1
}

func (ns Nodes) First() *Node {
	if len(ns) == 0 {
		return nil
	}
	return ns[0]
}
