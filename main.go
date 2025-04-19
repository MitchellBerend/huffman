package main

import (
	"fmt"
	"huffman/node"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fullTree := node.BuildTree(content)

	compressed := encode(content, fullTree)

	_ = os.WriteFile("compressed", compressed, 0666)

	decoded := decode(compressed, fullTree)

	if !(string(content) == string(decoded)) {
		node.PrintTree(fullTree, "    ", false)
		fmt.Printf("encoded:\n%s\n", content)
		fmt.Printf("decoded:\n%s\n", decoded)
	} else {
		fmt.Println("correct round trip")
	}
}

func encode(stream []byte, tree *node.Node) []byte {
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

func decode(buf []byte, tree *node.Node) []byte {
	amount := int(buf[0])
	streamLen := len(buf)
	root := tree
	current := tree
	_buf := buf[1:]

	rv := make([]byte, 0, len(buf))

	for i := range _buf {
		_counter := 7
		if i == streamLen-1 {
			_counter = amount
		}

		for j := _counter; j >= 0; j-- {
			bit := (_buf[i] >> j) & 1
			_bit := node.Bit(bit)
			current = current.GetChild(_bit)

			if current.IsLeaf() {
				rv = append(rv, current.Value...)

				current = root
			}
		}
	}

	return rv
}
