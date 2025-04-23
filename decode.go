package huffman

import (
	"huffman/node"
)

func Decode(buf []byte, tree *node.Node) []byte {
	amount := int(buf[0])
	streamLen := len(buf) - 2
	root := tree
	current := tree
	_buf := buf[1:]
	br := false

	rv := make([]byte, 0, len(buf))

	for i := range _buf {
		_counter := 7
		if i == streamLen {
			br = true
		}

		for j := _counter; j >= 0; j-- {
			bit := (_buf[i] >> j) & 1
			_bit := node.Bit(bit)
			current = current.GetChild(_bit)

			if current.IsLeaf() {
				rv = append(rv, current.Value...)
				current = root
			}

			if j-amount == 0 && br {
				break
			}
		}
	}

	return rv
}
