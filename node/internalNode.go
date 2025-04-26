package node

import (
	"sort"
)

type internalNode struct {
	Value  []byte
	left   *internalNode
	right  *internalNode
	Weight int64
}

func (i internalNode) toNode() *Node {
	nodes := i.ToArray()

	for len(nodes) > 1 {
		first := nodes[0]
		second := nodes[1]
		vals := first.Value
		vals = append(vals, second.Value...)

		newNode := NewNode(vals, first, second, 0)

		nodes = nodes[2:]

		nodes = append(nodes, newNode)

		sortNodes(nodes)
	}

	return nodes[0]
}

func (n internalNode) ToArray() []*Node {
	toCheck := []*internalNode{n.left, n.right}
	internalNodes := make([]*internalNode, 0, len(n.Value))
	rv := make([]*Node, 0, len(internalNodes))

	for len(toCheck) > 0 {
		current := toCheck[0]
		toCheck = toCheck[1:]

		if current.left == nil || current.right == nil {
			internalNodes = append(internalNodes, current)
		} else {
			toCheck = append(toCheck, current.left)
			toCheck = append(toCheck, current.right)
		}
	}

	sortInternalNodes(internalNodes)

	for idx, node := range internalNodes {
		rv = append(rv, NewNode(node.Value, nil, nil, idx+1))
	}

	return rv
}

func newInternalNode(value []byte, left *internalNode, right *internalNode, weight int64) *internalNode {
	_weight := weight

	if left != nil && right != nil {
		_weight = (left.Weight + right.Weight)
	}

	return &internalNode{value, left, right, _weight}
}

func sortInternalNodes(nodes []*internalNode) {
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].Weight != nodes[j].Weight {
			return nodes[i].Weight < nodes[j].Weight
		}
		if len(nodes[i].Value) != len(nodes[j].Value) {
			return len(nodes[i].Value) < len(nodes[j].Value)
		}
		return nodes[i].Value[0] < nodes[j].Value[0]
	})
}
