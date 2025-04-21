package node

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const lorem = "Lorem ipsum dolor sit amet consectetur adipiscing elit. Quisque faucibus ex sapien vitae pellentesque sem placerat. In id cursus mi pretium tellus duis convallis. Tempus leo eu aenean sed diam urna tempor. Pulvinar vivamus fringilla lacus nec metus bibendum egestas. Iaculis massa nisl malesuada lacinia integer nunc posuere. Ut hendrerit semper vel class aptent taciti sociosqu. Ad litora torquent per conubia nostra inceptos himenaeos."

// The order of this array should always be the same
func TestNodeToArrayOrder(t *testing.T) {
	tree := BuildTree([]byte(lorem))

	expected := []*Node{
		NewNode([]byte("A"), nil, nil, 1),
		NewNode([]byte("L"), nil, nil, 1),
		NewNode([]byte("P"), nil, nil, 1),
		NewNode([]byte("Q"), nil, nil, 1),
		NewNode([]byte("T"), nil, nil, 1),
		NewNode([]byte("U"), nil, nil, 1),
		NewNode([]byte("x"), nil, nil, 1),
		NewNode([]byte("I"), nil, nil, 2),
		NewNode([]byte("f"), nil, nil, 2),
		NewNode([]byte("h"), nil, nil, 2),
		NewNode([]byte("b"), nil, nil, 4),
		NewNode([]byte("g"), nil, nil, 4),
		NewNode([]byte("q"), nil, nil, 4),
		NewNode([]byte("v"), nil, nil, 6),
		NewNode([]byte("."), nil, nil, 8),
		NewNode([]byte("d"), nil, nil, 10),
		NewNode([]byte("p"), nil, nil, 13),
		NewNode([]byte("m"), nil, nil, 16),
		NewNode([]byte("o"), nil, nil, 16),
		NewNode([]byte("c"), nil, nil, 17),
		NewNode([]byte("r"), nil, nil, 19),
		NewNode([]byte("l"), nil, nil, 21),
		NewNode([]byte("n"), nil, nil, 25),
		NewNode([]byte("t"), nil, nil, 25),
		NewNode([]byte("u"), nil, nil, 27),
		NewNode([]byte("a"), nil, nil, 32),
		NewNode([]byte("i"), nil, nil, 34),
		NewNode([]byte("s"), nil, nil, 35),
		NewNode([]byte("e"), nil, nil, 45),
		NewNode([]byte(" "), nil, nil, 63),
	}

	nodes := tree.ToArray()

	for idx := range nodes {
		assert.Equal(t, expected[idx], nodes[idx])

	}
}
