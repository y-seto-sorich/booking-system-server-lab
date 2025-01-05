package util

import (
	"math/rand"
	"time"
)

func GenerateLoginPassword() string {
	rs1Letters := []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!#$%&'*+-/=?^_`{|}~.")
	rand.NewSource(time.Now().UTC().UnixNano())
	b := make([]rune, 10)
	for i := range b {
		b[i] = rs1Letters[rand.Intn(len(rs1Letters))] //#nosec G404
	}
	return string(b)
}
