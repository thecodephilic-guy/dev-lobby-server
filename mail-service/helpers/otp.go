package helpers

import (
	"crypto/rand"
	"io"
	"strconv"
)

func GenerateOTP() (string, error) {
	b := make([]byte, 3) //3 bytes can hold a number up to 2^24 - 1, which is enough for a 6-digit number

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}

	// Convert bytes to an integer
	num := int(b[0])<<16 | int(b[1])<<8 | int(b[2])

	// Ensure the number is within the 6-digit range (100,000 to 999,999)
	otp := num%900000 + 100000

	return strconv.Itoa(otp), nil
}
