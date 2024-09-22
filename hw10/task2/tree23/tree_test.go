package tree23

import (
	"testing"
)

func TestRemove(t *testing.T) {
	tree := NewTree()

	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(25)
	tree.Insert(35)

	tree.Delete(20)
	if !isValidTree(tree.root) {
		t.Errorf("Tree is not valid after removing 20")
	}
	if tree.Search(20) != nil {
		t.Errorf("Key 20 was not removed")
	}

	tree.Delete(10)
	if !isValidTree(tree.root) {
		t.Errorf("Tree is not valid after removing 10")
	}
	if tree.Search(10) != nil {
		t.Errorf("Key 10 was not removed")
	}

	tree.Delete(30)
	if !isValidTree(tree.root) {
		t.Errorf("Tree is not valid after removing 30")
	}
	if tree.Search(30) != nil {
		t.Errorf("Key 30 was not removed")
	}

	tree.Delete(5)
	if !isValidTree(tree.root) {
		t.Errorf("Tree is not valid after removing 5")
	}
	if tree.Search(5) != nil {
		t.Errorf("Key 5 was not removed")
	}

	tree.Delete(15)
	if !isValidTree(tree.root) {
		t.Errorf("Tree is not valid after removing 15")
	}
	if tree.Search(15) != nil {
		t.Errorf("Key 15 was not removed")
	}

	tree.Delete(25)
	if !isValidTree(tree.root) {
		t.Errorf("Tree is not valid after removing 25")
	}
	if tree.Search(25) != nil {
		t.Errorf("Key 25 was not removed")
	}

	tree.Delete(35)
	if !isValidTree(tree.root) {
		t.Errorf("Tree is not valid after removing 35")
	}
	if tree.Search(35) != nil {
		t.Errorf("Key 35 was not removed")
	}
}

// Helper function to check if a node is a valid 2-3 node
func isValidNode(node *Node) bool {
	if node == nil {
		return true
	}
	if len(node.vals) < 1 || len(node.vals) > 3 {
		return false
	}
	for i := 0; i < len(node.vals)-1; i++ {
		if node.vals[i] >= node.vals[i+1] {
			return false
		}
	}
	return true
}

// Helper function to check if the entire tree is a valid 2-3 tree
func isValidTree(node *Node) bool {
	if node == nil {
		return true
	}
	if !isValidNode(node) {
		return false
	}
	for _, child := range node.children {
		if child != nil {
			if child.parent != node {
				return false
			}
			if !isValidTree(child) {
				return false
			}
		}
	}
	return true
}
