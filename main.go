package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
)

func main() {

	// cipher key
	key := "thisis32bitlongpassphraseimusing"

	// plaintext
	pt := "This is a secret"

	c := EncryptAES([]byte(key), pt)

	// ciphertext
	fmt.Println(c)

}

func EncryptAES(key []byte, plaintext string) string {

	c, err := aes.NewCipher(key)
	CheckError(err)

	out := make([]byte, len(plaintext))

	c.Encrypt(out, []byte(plaintext))

	return hex.EncodeToString(out)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}