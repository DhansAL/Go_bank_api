package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// generates an random integer btw min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// generates ramdom string of length n
func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// generates a random owner name
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD", "IN"}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}
