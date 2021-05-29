package datagen

import (
	"math/rand"
	"time"
)

var letterRunes = []rune("thequickbrownfoxjumpedoverthelazydog")

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
	max := time.Date(2030, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	generatedTimeStamp := time.Unix(sec, 0)
	return generatedTimeStamp.Format(createdFormat)
}

const createdFormat = "2006-01-02 15:04:05" //"Jan 2, 2006 at 3:04pm (MST)"
