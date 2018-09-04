package base32

import (
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		input  string
		output []byte
	}{
		{"JBSWY3DPEHPK3PXP", []byte{'H', 'e', 'l', 'l', 'o', '!', 0xDE, 0xAD, 0xBE, 0xEF}},
		{"JBSWY3DPEBLW64TMMQQSASJAMFWSARTJNZSS4ICIN53SAYLSMUQHS33VEBSG62LOM4QHI2DFOJST6ICJEBUG64DFEB4W65JAMFZGKIDEN5UW4ZZAM5XW6ZBO", []byte{
			'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '!', ' ', 'I', ' ', 'a', 'm', ' ', 'F', 'i', 'n', 'e', '.', ' ', 'H', 'o', 'w', ' ', 'a', 'r', 'e', ' ', 'y', 'o', 'u', ' ', 'd', 'o', 'i', 'n', 'g', ' ', 't', 'h', 'e', 'r', 'e', '?', ' ', 'I', ' ', 'h', 'o', 'p', 'e', ' ', 'y', 'o', 'u', ' ', 'a', 'r', 'e', ' ', 'd', 'o', 'i', 'n', 'g', ' ', 'g', 'o', 'o', 'd', '.',
		}},
	}

	base32 := new(Base32)

	for _, test := range tests {
		output, _ := base32.Decode([]byte(test.input))
		if !equal(output, test.output) {
			fmt.Println(test, output, string(output))
			t.Fail()
		}
	}
}

func equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
