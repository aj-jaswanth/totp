package base32

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

func Decode(input string) []byte {
	inputLength := len(input)
	est := estimatedOutputLength(inputLength)

	output := make([]byte, est)
	var holder uint64
	for i := 0; i < inputLength/8; i++ {
		holder = 0
		for j := 0; j < 8; j++ {
			holder = (holder << 5) | uint64(alphaTobase32[input[(8*i)+j]])
		}
		var slider uint64 = 0xff << 32
		for k := 0; k < 5; k++ {
			output[(5*i)+k] = byte((holder & slider) >> (32 - uint(k*8)))
			slider = slider >> 8
		}
	}
	return output
}

func estimatedOutputLength(inputLength int) int {
	return (inputLength * 5) / 8
}
