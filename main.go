package main

import (
	"fmt"
	"github.com/aj-jaswanth/totp/decode/base32"
	"github.com/aj-jaswanth/totp/secret"
	"github.com/aj-jaswanth/totp/totp"
	"log"
)

func main() {
	fileSecret := secret.FileSecret{Decoder: new(base32.Base32)}
	otp := totp.Totp{Secret: fileSecret, Time: new(totp.UnixTime), StepInterval: 30}
	value, err := otp.Generate()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value[len(value)-6:])
}
