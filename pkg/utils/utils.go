package utils

import (
	"crypto/rand"
	"math/big"
)

var charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var charSetLength = len(charSet)

func RandomBase62(len int) (string, error) {
	result := make([]byte, len)

	for i := range result {
		ind, err := rand.Int(rand.Reader, big.NewInt(int64(charSetLength)))
		if err != nil {
			return "", err
		}

		result[i] = charSet[ind.Int64()]
	}

	return string(result), nil
}
