package utils

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	randStr = rand.New(rand.NewSource(time.Now().Unix()))
	letters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPLKJHGFDSAZCXVBNM0123456789")
)

func GetTestKey(i int) []byte {
	return []byte(fmt.Sprintf("bitcask-go-%09d", i))
}

func RandomValue(n int) []byte {
	b := make([]byte, n)

	for i := range b {
		b[i] = letters[randStr.Intn(len(letters))]
	}
	return []byte("bitcask-go-value" + string(b))
}
