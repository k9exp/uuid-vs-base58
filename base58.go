package main

import (
	"crypto/rand"
)

// BASE-58
var alphabets = []rune("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func New(size uint16) (string, error) {
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)

	if err != nil {
		return "", err
	}

	id := make([]rune, size)

	for i := uint16(0); i < size; i++ {
		id[i] = alphabets[bytes[i]&57]
	}
	return string(id[:size]), nil
}
