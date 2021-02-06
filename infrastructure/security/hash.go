package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func createHash(secret []byte, plainText []byte, channel chan string) {
	h := hmac.New(sha256.New, secret)

	h.Write(plainText)

	channel <- hex.EncodeToString(h.Sum(nil))
}

func CreateHash(secret []byte, plainText []byte) string {
	channel := make(chan string)

	go createHash(secret, plainText, channel)

	result := <-channel

	close(channel)

	return result
}
