package util_test

import (
	"math/rand"
	"time"
)

func RandLetters(n int) string {
	rs1Letters := []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	rand.NewSource(time.Now().UTC().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = rs1Letters[rand.Intn(len(rs1Letters))] //#nosec G404
	}
	return string(b)
}

func RandLettersNumbersOnly(n int) string {
	rs1Letters := []rune("0123456789")
	rand.NewSource(time.Now().UTC().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = rs1Letters[rand.Intn(len(rs1Letters))] //#nosec G404
	}
	return string(b)
}
