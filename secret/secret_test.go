package secret

import (
	"os"
	"testing"
)

type DummyDecoder struct {
}

func (DummyDecoder) Decode(input []byte) ([]byte, error) {
	return input, nil
}

func Test_GetSecretFromFile(t *testing.T) {
	file, _ := os.Create("test.txt")
	file.WriteString("test secret")
	file.Close()
	defer os.Remove("test.txt")

	fileSecret := FileSecret{Decoder: new(DummyDecoder)}

	os.Setenv("SLOC", "test.txt")
	secret, e := fileSecret.Get()
	if e != nil || secret == nil {
		t.Fail()
	}
}

func Test_GetSecretWithInvalidFile(t *testing.T) {
	os.Setenv("SLOC", "blah.txt")

	fileSecret := FileSecret{Decoder: new(DummyDecoder)}

	_, e := fileSecret.Get()

	if e == nil {
		t.Fail()
	}
}
