package tree23

import (
	"sort"
)

type Node struct {
	vals     []int
	parent   *Node
	children []*Node
}

func (n *Node) isLeaf() bool {
	return len(n.children) == 0
}

func (n *Node) insertVal(val int) {
	n.vals = append(n.vals, val)
	n.sortVals()
}

func (n *Node) sortVals() {
	sort.Ints(n.vals)
}

func (n *Node) findChild(val int) *Node {
	if n.isLeaf() {
		return nil
	}

	for i, v := range n.vals {
		if val < v {
			return n.children[i]
		}
	}

	return n.children[len(n.children)-1]
}

func (n *Node) GetChildren() []*Node {
	return n.children
}

func (n *Node) GetVals() []int {
	return n.vals
}

func (n *Node) deleteFromLeaf(val int) {
	for i, v := range n.vals {
		if v == val {
			n.vals = append(n.vals[:i], n.vals[i+1:]...)
			break
		}
	}

	if len(n.vals) == 0 && n.parent != nil {
		n.parent.removeChild(n)
	}
}

func (n *Node) removeChild(child *Node) {
	for i, c := range n.children {
		if c == child {
			n.children = append(n.children[:i], n.children[i+1:]...)
			break
		}
	}
}

func (n *Node) hasVal(val int) bool {
	for _, v := range n.vals {
		if v == val {
			return true
		}
	}
	return false
}

func (n *Node) replaceVal(old, new int) {
	for i, v := range n.vals {
		if v == old {
			n.vals[i] = new
			n.sortVals()
			return
		}
	}
}

func (n *Node) getPredecessor(val int) *Node {
	for i, v := range n.vals {
		if v == val {
			child := n.children[i]
			for !child.isLeaf() {
				child = child.children[len(child.children)-1]
			}
			return child
		}
	}
	return nil
}
