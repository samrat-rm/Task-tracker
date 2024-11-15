package utils

import (
	"math/rand"
	"time"
)

// generateRandomID generates a random 5-digit integer ID
func GenerateRandomID() int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(90000) + 10000 // Return a number between 10000 and 99999
}
