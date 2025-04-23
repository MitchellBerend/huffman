package huffman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	blob := []byte{0x3, 0x54, 0xfe, 0x0}

	decoded := Decode(blob, ac21)

	assert.Equal(t,
		content,
		string(decoded),
	)
}
