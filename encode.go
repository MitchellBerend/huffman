package huffman

import (
	"huffman/node"
)

func Encode(stream []byte, tree *node.Node) []byte {
	var direction node.Bit

	buf := []byte{}
	root := tree
	current := tree
	counter := 0
	b := 0

	for _, char := range stream {
		for !current.IsLeaf() {
			if counter == 8 {
				buf = append(buf, byte(b))
				counter = 0
				b = 0
			}
			current, direction = current.ChildContains([]byte(string(char)))
			b = node.PushBit(b, direction)
			counter++
		}
		current = root
	}

	buf = append([]uint8{uint8(8 - counter)}, buf...)

	for counter > 0 && counter < 8 {
		counter++
		b = node.PushBit(b, node.ZERO)
	}

	buf = append(buf, byte(b))

	return buf
}
