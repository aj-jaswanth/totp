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
	}

	for _, test := range tests {
		output := Decode(test.input)
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
