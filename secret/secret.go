package secret

import (
	"io/ioutil"
	"os"
)

type FileSecret struct {
	Decoder Decoder
}

func (secret FileSecret) Get() ([]byte, error) {
	content, e := ioutil.ReadFile(os.Getenv("SLOC"))
	if e != nil {
		return nil, e
	}
	return secret.Decoder.Decode(content)
}

type Decoder interface {
	Decode([]byte) ([]byte, error)
}
