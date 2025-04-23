package node

import (
	"bytes"
	"fmt"
)

type Node struct {
	Value  []byte
	Left   *Node
	Right  *Node
	Weight int64
}

func (n Node) ToArray() []*Node {
	toCheck := []*Node{n.Left, n.Right}
	rv := make([]*Node, 0, len(n.Value))

	for len(toCheck) > 0 {
		current := toCheck[0]
		toCheck = toCheck[1:]

		if current.IsLeaf() {
			rv = append(rv, current)
		} else {
			toCheck = append(toCheck, current.Left)
			toCheck = append(toCheck, current.Right)
		}
	}

	sortNodes(rv)

	return rv
}

func (n Node) ToString() string {
	return fmt.Sprintf("%s%d", n.Value, n.Weight)
}

func (n Node) IsLeaf() bool {
	return n.Left == nil || n.Right == nil
}

func (n Node) ChildContains(value []byte) (*Node, Bit) {
	if bytes.Contains(n.Left.Value, value) {
		return n.Left, ONE
	}

	return n.Right, ZERO
}

func (n *Node) GetChild(bit Bit) *Node {
	if n.IsLeaf() {
		return n
	}

	var rv *Node
	switch bit {
	case ZERO:
		rv = n.Right

	case ONE:
		rv = n.Left
	}

	return rv
}

func NewNode(value []byte, left *Node, right *Node, weight int64) *Node {
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

	PrintTree(node.Left, newIndent, true)
	PrintTree(node.Right, newIndent, false)
}
