package rand

import (
	"math/rand"
	"time"
)

func GetRandomString(nSize int) string {
	chars := "123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(chars)
	value := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < nSize; i++ {
		value = append(value, bytes[r.Intn(len(bytes))])
	}
	return string(value)
}