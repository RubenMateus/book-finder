package random

import "math/rand"

// RandomInt generates a random int, based on a min and max values
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
