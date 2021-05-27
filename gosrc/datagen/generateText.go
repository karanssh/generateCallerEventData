package datagen

import (
	"math/rand"
	"time"
)

var letterRunes = []rune("the quick brown fox jumped over the lazy dog but could not win the race")

// RandomStringFromRunes generates a random string of len n from given input
func RandomStringFromRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// RandomTime generates a random timestamp
func RandomTime() string {
	min := time.Date(2000, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	timeGen := time.Unix(sec, 0)
	return timeGen.String()
}
