package datagen

import "math/rand"

// RandomNumberInRange generates a random number between a given range
func RandomNumberInRange(low, high int) int {
	return low + rand.Intn(high-low)
}
