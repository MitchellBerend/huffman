package node

import (
	"sort"
)

func BuildTreeFromBinary(stream []byte) *Node {
	if len(stream)%2 != 0 {
		panic("Tree binary malformed")
	}
	nodeLen := len(stream) / 2

	nodes := make([]*Node, 0, nodeLen)

	for idx := range nodeLen {
		modifier := idx * 2
		b := stream[modifier]
		weight := stream[modifier+1]

		nodes = append(nodes, NewNode([]byte{b}, nil, nil, int(weight)))
	}

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

func BuildTree(stream []byte) *Node {
	ir := make(map[byte]int64)

	for _, char := range stream {
		ir[char] = ir[char] + 1
	}

	nodes := []*internalNode{}

	for key := range ir {
		node := newInternalNode([]byte{key}, nil, nil, ir[key])
		nodes = append(nodes, node)
	}

	sortInternalNodes(nodes)

	// for _, node:= range nodes{
	// 	fmt.Printf("%s:%d\n", string(node.Value), node.Weight)
	// }

	for len(nodes) > 1 {
		first := nodes[0]
		second := nodes[1]
		vals := first.Value
		vals = append(vals, second.Value...)

		newNode := newInternalNode(vals, first, second, 0)

		nodes = nodes[2:]

		nodes = append(nodes, newNode)

		sortInternalNodes(nodes)
	}

	finalInternalNode := nodes[0]
	return finalInternalNode.toNode()
}

func sortNodes(nodes []*Node) {
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
