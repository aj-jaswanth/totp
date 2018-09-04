package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"github.com/aj-jaswanth/totp/base32"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	StepInterval = 30
)

func main() {
	stepsSinceEpoch := time.Now().Unix() / StepInterval

	decodedSecret := []byte(base32.Decode(secret()))
	hmacSha1 := computeHmacSha1(decodedSecret, uint64(stepsSinceEpoch))

	offset := hmacSha1[19] & 0x0f
	dynamicExtract := hmacSha1[offset : offset+4]

	var val uint32
	for _, v := range dynamicExtract {
		val = (val << 8) | uint32(v)
	}

	val = val & 0x7fffffff
	toChars := strconv.Itoa(int(val))
	fmt.Println(toChars[len(toChars)-6:])
}

func computeHmacSha1(secret []byte, stepsSinceEpoch uint64) []byte {
	hmacSha1 := hmac.New(sha1.New, secret)
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, stepsSinceEpoch)
	_, err := hmacSha1.Write(b)
	if err != nil {
		log.Fatal(err)
	}
	return hmacSha1.Sum(nil)
}

func secret() string {
	content, e := ioutil.ReadFile(os.Getenv("SLOC"))
	if e != nil {
		log.Fatal(e)
	}
	return string(content)
}
