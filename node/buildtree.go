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
		node := newNode([]byte{key}, nil, nil, ir[key])
		nodes = append(nodes, node)
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Weight < nodes[j].Weight
	})

	for len(nodes) > 1 {
		first := nodes[0]
		second := nodes[1]
		vals := first.Value
		vals = append(vals, second.Value...)

		newNode := newNode(vals, first, second, 0)

		nodes = nodes[2:]

		nodes = append(nodes, newNode)

		sort.Slice(nodes, func(i, j int) bool {
			return nodes[i].Weight < nodes[j].Weight
		})
	}

	finalNode := nodes[0]

	return finalNode
}
