package node

type Bit int

const (
	ZERO Bit = 0
	ONE  Bit = 1
)

func PushBit(b int, bit Bit) int {
	return (b << 1) | int(bit)
}
