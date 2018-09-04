package totp

import (
	"fmt"
	"strings"
	"testing"
)

type DummyTime struct {
	DummyTime uint64
}

func (dummyTime DummyTime) Now() uint64 {
	return dummyTime.DummyTime
}

type DummySecret struct {
}

func (DummySecret) Get() ([]byte, error) {
	return []byte("12345678901234567890"), nil
}

// https://tools.ietf.org/html/rfc6238#appendix-B
func Test_TotpGeneration(t *testing.T) {
	tests := []struct {
		time  uint64
		value string
	}{
		{59, "94287082"},
		{1111111109, "07081804"},
		{1111111111, "14050471"},
		{1234567890, "89005924"},
		{2000000000, "69279037"},
		{20000000000, "65353130"},
	}

	for _, test := range tests {
		totp := Totp{StepInterval: 30, Time: DummyTime{DummyTime: test.time}, Secret: new(DummySecret)}
		value, err := totp.Generate()
		if err != nil || !strings.HasSuffix(value, test.value) {
			fmt.Println(test, value, err)
			t.Fail()
		}
	}
}
