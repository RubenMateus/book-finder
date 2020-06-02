package random

import (
	crypto "crypto/rand"
	"fmt"
	"io"
	math "math/rand"
)

// RandomInt generates a random int, based on a min and max values
func RandomInt(min, max int) int {
	return min + math.Intn(max-min)
}

// Generate a random UUID according to RFC 4122
func NewUUID() (string, error) {
	uuid := make([]byte, 16)

	n, err := io.ReadFull(crypto.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}

	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
