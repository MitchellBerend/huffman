package huffman

import (
	"huffman/node"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	a  *node.Node = node.NewNode([]byte("a"), nil, nil, 3)
	c  *node.Node = node.NewNode([]byte("c"), nil, nil, 1)
	_1 *node.Node = node.NewNode([]byte("1"), nil, nil, 6)
	_2 *node.Node = node.NewNode([]byte("2"), nil, nil, 2)

	c2  *node.Node = node.NewNode([]byte("c2"), c, _2, 0)
	ac2 *node.Node = node.NewNode([]byte("ac2"), a, c2, 0)

	ac21 *node.Node = node.NewNode([]byte("ac21"), _1, ac2, 0)
)

const (
	content = "aaac11111122"
	lorem   = "Lorem ipsum dolor sit amet consectetur adipiscing elit. Quisque faucibus ex sapien vitae pellentesque sem placerat. In id cursus mi pretium tellus duis convallis. Tempus leo eu aenean sed diam urna tempor. Pulvinar vivamus fringilla lacus nec metus bibendum egestas. Iaculis massa nisl malesuada lacinia integer nunc posuere. Ut hendrerit semper vel class aptent taciti sociosqu. Ad litora torquent per conubia nostra inceptos himenaeos."
)

func TestEncode(t *testing.T) {
	encoded := Encode([]byte(content), ac21)
	expectedNodes := []*node.Node{
		node.NewNode([]byte("c"), nil, nil, 1),
		node.NewNode([]byte("2"), nil, nil, 2),
		node.NewNode([]byte("a"), nil, nil, 3),
		node.NewNode([]byte("1"), nil, nil, 6),
	}
	tree := node.BuildTree([]byte(content))

	for idx, node := range tree.ToArray() {
		assert.Equal(t, expectedNodes[idx].Value, node.Value)
	}

	assert.Equal(t,
		[]byte{0x3, 0x54, 0xfe, 0x0},
		encoded,
	)
}

func TestEncodeLong(t *testing.T) {
	encoded := Encode([]byte(lorem), ac21)
	tree := node.BuildTree([]byte(lorem))
	expectedNodes := []*node.Node{
		node.NewNode([]byte("A"), nil, nil, 1),
		node.NewNode([]byte("L"), nil, nil, 1),
		node.NewNode([]byte("P"), nil, nil, 1),
		node.NewNode([]byte("Q"), nil, nil, 1),
		node.NewNode([]byte("T"), nil, nil, 1),
		node.NewNode([]byte("U"), nil, nil, 1),
		node.NewNode([]byte("x"), nil, nil, 1),
		node.NewNode([]byte("I"), nil, nil, 2),
		node.NewNode([]byte("f"), nil, nil, 2),
		node.NewNode([]byte("h"), nil, nil, 2),
		node.NewNode([]byte("b"), nil, nil, 4),
		node.NewNode([]byte("g"), nil, nil, 4),
		node.NewNode([]byte("q"), nil, nil, 4),
		node.NewNode([]byte("v"), nil, nil, 6),
		node.NewNode([]byte("."), nil, nil, 8),
		node.NewNode([]byte("d"), nil, nil, 10),
		node.NewNode([]byte("p"), nil, nil, 13),
		node.NewNode([]byte("m"), nil, nil, 16),
		node.NewNode([]byte("o"), nil, nil, 16),
		node.NewNode([]byte("c"), nil, nil, 17),
		node.NewNode([]byte("r"), nil, nil, 19),
		node.NewNode([]byte("l"), nil, nil, 21),
		node.NewNode([]byte("n"), nil, nil, 25),
		node.NewNode([]byte("t"), nil, nil, 25),
		node.NewNode([]byte("u"), nil, nil, 27),
		node.NewNode([]byte("a"), nil, nil, 32),
		node.NewNode([]byte("i"), nil, nil, 34),
		node.NewNode([]byte("s"), nil, nil, 35),
		node.NewNode([]byte("e"), nil, nil, 45),
		node.NewNode([]byte(" "), nil, nil, 63),
	}

	for idx, node := range tree.ToArray() {
		assert.Equal(t, expectedNodes[idx].Value, node.Value)
	}

	assert.Equal(t,
		[]byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x10, 0x0, 0x20, 0x0, 0x40, 0x0, 0x4, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x20, 0x80, 0x0, 0x0, 0x4, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x90, 0x10, 0x0, 0x0, 0x0, 0x40, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x10, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x8, 0x1, 0x0, 0x0, 0x0, 0x40, 0x0, 0x20, 0x0, 0x0, 0x0, 0x0, 0x0, 0x40, 0x0, 0x20, 0x0, 0x0, 0x0, 0x0, 0x80, 0x90, 0x0, 0x4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x9, 0x0, 0x0, 0x10, 0x10, 0x0, 0x0, 0x20, 0x0, 0x84, 0x4, 0x80, 0x10, 0x0, 0x0, 0x0, 0x0, 0x20, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x10, 0x80, 0x10, 0x0, 0x0, 0x24, 0x0, 0x0, 0x20, 0x0, 0x0, 0x0, 0x0, 0x0, 0x40, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x1, 0x0, 0x0, 0x10, 0x1, 0x0, 0x0, 0x0, 0x0, 0x20, 0x0},
		encoded,
	)
}
