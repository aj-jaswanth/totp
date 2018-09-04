package totp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"strconv"
)

type Totp struct {
	Secret       Secret
	Time         Time
	StepInterval int
}

// https://tools.ietf.org/html/rfc4226#section-5
func (totp *Totp) Generate() (string, error) {
	stepsSinceEpoch := totp.Time.Now() / uint64(totp.StepInterval)
	secret, err := totp.Secret.Get()
	if err != nil {
		return "", err
	}
	hmacValue, err := computeHmacSha1(secret, stepsSinceEpoch)
	if err != nil {
		return "", nil
	}
	val := dynamicFourBytesSquashed(hmacValue)
	val = val & 0x7fffffff
	return strconv.Itoa(int(val)), nil
}

func dynamicFourBytesSquashed(hmacValue []byte) uint32 {
	offset := hmacValue[19] & 0x0f
	dynamicExtract := hmacValue[offset : offset+4]
	var val uint32
	for _, v := range dynamicExtract {
		val = (val << 8) | uint32(v)
	}
	return val
}

func computeHmacSha1(secret []byte, stepsSinceEpoch uint64) ([]byte, error) {
	hmacSha1 := hmac.New(sha1.New, secret)
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, stepsSinceEpoch)
	_, err := hmacSha1.Write(bytes)
	if err != nil {
		return nil, err
	}
	return hmacSha1.Sum(nil), nil
}

type Secret interface {
	Get() ([]byte, error)
}

type Time interface {
	Now() uint64
}
