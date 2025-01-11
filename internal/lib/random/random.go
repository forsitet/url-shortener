package random

import (
	"crypto/rand"
	"math/big"
	"strings"
)

func NewRandomString(aliasLength int) string {
	str := strings.Builder{}
	alph := []rune("abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789")
	for i := 0; i < aliasLength; i++ {
		inx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(alph))))
		str.WriteRune(alph[inx.Int64()])
	}
	return str.String()
}
