package service

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand" // genrating random numbers
	"strconv"   // for converting between strings and numbers
)

func GenerateID(movieName string) string {
	hash := sha256.Sum256([]byte(movieName))
	hashHex := hex.EncodeToString(hash[:]) // convert to hex string

	first6chars := hashHex[:6]
	last6chars := hashHex[len(hashHex)-6:]
	randomNum := strconv.Itoa(rand.Intn(100)) // generate's rand number between 0-99
	id := first6chars + randomNum + last6chars

	return id
}
