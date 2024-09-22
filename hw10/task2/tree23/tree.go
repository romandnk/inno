package tree23

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Insert(val int) {
	if t.root == nil {
		t.root = &Node{vals: []int{val}}
		return
	}

	t.insert(t.root, val)

	if len(t.root.vals) == 3 {
		t.splitRoot()
	}
}

func (t *Tree) insert(node *Node, val int) {
	if node.isLeaf() {
		node.insertVal(val)
	} else {
		child := node.findChild(val)
		t.insert(child, val)
		if len(child.vals) == 3 {
			t.splitChild(node, child)
		}
	}
}

func (t *Tree) splitChild(parent, child *Node) {
	midIndex := 1
	midVal := child.vals[midIndex] // 2

	left := &Node{vals: []int{child.vals[0]}, parent: parent}  // [1] - [5 9]
	right := &Node{vals: []int{child.vals[2]}, parent: parent} // [4] - [5 9]

	if len(child.children) > 0 {
		left.children = append(left.children, child.children[0], child.children[1])
		right.children = append(right.children, child.children[2], child.children[3])
		for _, c := range left.children {
			c.parent = left
		}
		for _, c := range right.children {
			c.parent = right
		}
	}

	for i, v := range parent.vals {
		if midVal < v {
			parent.vals = append(parent.vals[:i], append([]int{midVal}, parent.vals[i:]...)...)
			parent.children = append(parent.children[:i], append([]*Node{left, right}, parent.children[i+1:]...)...)
			return
		}
	}

	parent.vals = append(parent.vals, midVal)
	parent.children = append(parent.children[:len(parent.children)-1], left, right)
}

func (t *Tree) splitRoot() {
	oldRoot := t.root
	midIndex := 1
	midVal := oldRoot.vals[midIndex]

	left := &Node{vals: []int{oldRoot.vals[0]}, parent: nil}
	right := &Node{vals: []int{oldRoot.vals[2]}, parent: nil}

	if len(oldRoot.children) > 0 {
		left.children = append(left.children, oldRoot.children[0], oldRoot.children[1])
		right.children = append(right.children, oldRoot.children[2], oldRoot.children[3])
		for _, c := range left.children {
			c.parent = left
		}
		for _, c := range right.children {
			c.parent = right
		}
	}

	t.root = &Node{vals: []int{midVal}, children: []*Node{left, right}}
	left.parent = t.root
	right.parent = t.root
}

func (t *Tree) Search(val int) *Node {
	return t.search(t.root, val)
}

func (t *Tree) search(node *Node, val int) *Node {
	if node == nil {
		return nil
	}

	for _, v := range node.vals {
		if v == val {
			return node
		}
	}

	if node.isLeaf() {
		return nil
	}

	child := node.findChild(val)
	return t.search(child, val)
}

func (t *Tree) Delete(val int) {
	if t.root == nil {
		return
	}
	t.deleteVal(t.root, val)

	if len(t.root.vals) == 0 && len(t.root.children) > 0 {
		t.root = t.root.children[0]
		t.root.parent = nil
	}

	if len(t.root.vals) == 0 && len(t.root.children) == 0 {
		t.root = nil
	}
}

func (t *Tree) deleteVal(node *Node, val int) {
	if node.isLeaf() {
		node.deleteFromLeaf(val)
	} else {
		child := node.findChild(val)
		if node.hasVal(val) {
			predecessor := node.getPredecessor(val)
			node.replaceVal(val, predecessor.vals[len(predecessor.vals)-1])
			t.deleteVal(predecessor, predecessor.vals[len(predecessor.vals)-1])
		} else {
			t.deleteVal(child, val)
		}
		if len(child.vals) == 0 {
			t.rebalance(node, child)
		}
	}
}

func (t *Tree) rebalance(parent, child *Node) {
	childIndex := -1
	for i, c := range parent.children {
		if c == child {
			childIndex = i
			break
		}
	}

	if childIndex == -1 {
		return
	}

	if childIndex > 0 && len(parent.children[childIndex-1].vals) > 1 {
		leftSibling := parent.children[childIndex-1]
		child.vals = append([]int{parent.vals[childIndex-1]}, child.vals...)
		parent.vals[childIndex-1] = leftSibling.vals[len(leftSibling.vals)-1]
		leftSibling.vals = leftSibling.vals[:len(leftSibling.vals)-1]
		if !leftSibling.isLeaf() {
			child.children = append([]*Node{leftSibling.children[len(leftSibling.children)-1]}, child.children...)
			leftSibling.children = leftSibling.children[:len(leftSibling.children)-1]
		}
	} else if childIndex < len(parent.children)-1 && len(parent.children[childIndex+1].vals) > 1 {
		rightSibling := parent.children[childIndex+1]
		child.vals = append(child.vals, parent.vals[childIndex])
		parent.vals[childIndex] = rightSibling.vals[0]
		rightSibling.vals = rightSibling.vals[1:]
		if !rightSibling.isLeaf() {
			child.children = append(child.children, rightSibling.children[0])
			rightSibling.children = rightSibling.children[1:]
		}
	} else {
		if childIndex > 0 {
			leftSibling := parent.children[childIndex-1]
			leftSibling.vals = append(leftSibling.vals, parent.vals[childIndex-1])
			leftSibling.vals = append(leftSibling.vals, child.vals...)
			leftSibling.children = append(leftSibling.children, child.children...)
			parent.vals = append(parent.vals[:childIndex-1], parent.vals[childIndex:]...)
			parent.children = append(parent.children[:childIndex], parent.children[childIndex+1:]...)
		} else {
			rightSibling := parent.children[childIndex+1]
			child.vals = append(child.vals, parent.vals[childIndex])
			child.vals = append(child.vals, rightSibling.vals...)
			child.children = append(child.children, rightSibling.children...)
			parent.vals = append(parent.vals[:childIndex], parent.vals[childIndex+1:]...)
			parent.children = append(parent.children[:childIndex+1], parent.children[childIndex+2:]...)
		}
	}
	if len(parent.vals) == 0 && parent.parent != nil {
		t.rebalance(parent.parent, parent)
	}
}
