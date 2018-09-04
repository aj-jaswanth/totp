package totp

import "time"

type UnixTime struct {
}

func (*UnixTime) Now() uint64 {
	return uint64(time.Now().Unix())
}
