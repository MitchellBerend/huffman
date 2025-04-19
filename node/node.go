package node

import (
	"bytes"
	"fmt"
)

type Node struct {
	Value  []byte
	left   *Node
	right  *Node
	Weight int64
}

func (n Node) ToString() string {
	return fmt.Sprintf("%s%d", n.Value, n.Weight)
}

func (n Node) IsLeaf() bool {
	return n.left == nil || n.right == nil
}

func (n Node) ChildContains(value []byte) (*Node, Bit) {
	if bytes.Contains(n.left.Value, value) {
		return n.left, ONE
	}

	return n.right, ZERO
}

func (n *Node) GetChild(bit Bit) *Node {
	if n.IsLeaf() {
		return n
	}

	var rv *Node
	switch bit {
	case ZERO:
		rv = n.right

	case ONE:
		rv = n.left
	}

	return rv
}

func newNode(value []byte, left *Node, right *Node, weight int64) *Node {
	_weight := weight

	if left != nil && right != nil {
		_weight = (left.Weight + right.Weight)
	}

	return &Node{value, left, right, _weight}
}

// Chat gpt generated
func PrintTree(node *Node, indent string, isLeft bool) {
	if node == nil {
		return
	}

	prefix := "├── "
	if !isLeft {
		prefix = "└── "
	}

	fmt.Printf("%s%sLeaf(Value: '%s', Weight: %d)\n", indent, prefix, node.Value, node.Weight)

	newIndent := indent
	if isLeft {
		newIndent += "│   "
	} else {
		newIndent += "    "
	}

	PrintTree(node.left, newIndent, true)
	PrintTree(node.right, newIndent, false)
}
