package base32

type Base32 struct{}

func (*Base32) Decode(input []byte) ([]byte, error) {
	inputLength := len(input)
	est := estimatedOutputLength(inputLength)
	output := make([]byte, est)

	for chunk := 0; chunk < inputLength/8; chunk++ {
		window := 8 * chunk
		holder := squashBytes(input[window : window+8])
		transformToBytes(output, holder, chunk)
	}
	return output, nil
}

func transformToBytes(output []byte, holder uint64, chunk int) {
	var slider uint64 = 0xff << 32
	for index := 0; index < 5; index++ {
		output[(5*chunk)+index] = byte((holder & slider) >> (32 - uint(index*8)))
		slider = slider >> 8
	}
}

func squashBytes(bytes []byte) uint64 {
	var holder uint64
	for _, b := range bytes {
		value := alphaTobase32[b]
		holder = (holder << 5) | uint64(value)
	}
	return holder
}

func estimatedOutputLength(inputLength int) int {
	return (inputLength * 5) / 8
}

var alphaTobase32 = map[byte]byte{
	'A': 0,
	'a': 0,
	'B': 1,
	'b': 1,
	'C': 2,
	'c': 2,
	'D': 3,
	'd': 3,
	'E': 4,
	'e': 4,
	'F': 5,
	'f': 5,
	'G': 6,
	'g': 6,
	'H': 7,
	'h': 7,
	'I': 8,
	'i': 8,
	'J': 9,
	'j': 9,
	'K': 10,
	'k': 10,
	'L': 11,
	'l': 11,
	'M': 12,
	'm': 12,
	'N': 13,
	'n': 13,
	'O': 14,
	'o': 14,
	'P': 15,
	'p': 15,
	'Q': 16,
	'q': 16,
	'R': 17,
	'r': 17,
	'S': 18,
	's': 18,
	'T': 19,
	't': 19,
	'U': 20,
	'u': 20,
	'V': 21,
	'v': 21,
	'W': 22,
	'w': 22,
	'X': 23,
	'x': 23,
	'Y': 24,
	'y': 24,
	'Z': 25,
	'z': 25,
	'2': 26,
	'3': 27,
	'4': 28,
	'5': 29,
	'6': 30,
	'7': 31,
}
