package util

import (
	"math/rand"
	"time"
)

func RandomString(length int) string {
	charset := "0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(result)
}
