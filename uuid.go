package gengoutil

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"strings"
	"time"
)

func GenerateUUIDFromTimestamp(timestamp time.Time) string {
	bigInt := new(big.Int).SetInt64(timestamp.UnixNano())

	// Optionally pad the number to ensure it's 128 bits
	// Shift left to create a 128-bit representation
	bigInt.Lsh(bigInt, 64)

	// Convert to hexadecimal
	hexNum := fmt.Sprintf("%x", bigInt)

	// Output the result
	UUID := fmt.Sprintf("%s-%s-%s-%s-%s",
		hexNum[0:8],
		hexNum[8:12],
		hexNum[12:16],
		hexNum[16:20],
		hexNum[20:])

	return UUID
}

func GenerateUUIDFromString(input string) string {
	// Create a SHA-256 hash of the string
	hash := sha256.New()
	hash.Write([]byte(input))
	hashBytes := hash.Sum(nil)

	// Extract the first 16 bytes for UUID (128 bits)
	uuidBytes := hashBytes[:16]

	// Convert the bytes to the UUID format: 8-4-4-4-12
	var uuid strings.Builder
	for i, b := range uuidBytes {
		if i == 4 || i == 6 || i == 8 || i == 10 {
			uuid.WriteByte('-')
		}
		uuid.WriteString(fmt.Sprintf("%02x", b))
	}

	return uuid.String()
}
