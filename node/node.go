package node

import (
	"bytes"
	"fmt"
)

type Node struct {
	Value  []byte
	left   *Node
	right  *Node
	Weight int
}

func (n Node) ToArray() []*Node {
	toCheck := []*Node{n.left, n.right}
	rv := make([]*Node, 0, len(n.Value))

	for len(toCheck) > 0 {
		current := toCheck[0]
		toCheck = toCheck[1:]

		if current.IsLeaf() {
			rv = append(rv, current)
		} else {
			toCheck = append(toCheck, current.left)
			toCheck = append(toCheck, current.right)
		}
	}

	sortNodes(rv)

	return rv
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

// This produces a byte array that has the following format:
// <byte><weight><byte><weight><byte><weight>
// Where the first byte is the actual byte value from a leaf in the tree and
// the second byte is the weight as a u8
func (n Node) ToBinary() []byte {
	bin := []byte{}

	nodes := n.ToArray()

	for _, node := range nodes {
		bin = append(bin, node.Value[0])
		bin = append(bin, byte(node.Weight))
	}

	return bin
}

func NewNode(value []byte, left *Node, right *Node, weight int) *Node {
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
