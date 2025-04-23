package node

import (
	"sort"
)

func BuildTree(stream []byte) *Node {
	ir := make(map[byte]int64)

	for _, char := range stream {
		ir[char] = ir[char] + 1
	}

	nodes := []*Node{}

	for key := range ir {
		node := NewNode([]byte{key}, nil, nil, ir[key])
		nodes = append(nodes, node)
	}

	sortNodes(nodes)

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

	finalNode := nodes[0]

	return finalNode
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
