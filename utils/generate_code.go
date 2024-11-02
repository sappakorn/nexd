package utils

import (
	"math/rand"
	"strings"
	"time"
)

func GenerateCodeNumberAndString() string {
	rand.Seed(time.Now().UnixNano())

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomString2 := ""
	for i := 0; i < 4; i++ {
		randomIndex := rand.Intn(len(charset))
		randomString2 += string(charset[randomIndex])
	}
	return strings.ToUpper(randomString2)
}


func GenerateOTP() int {
	// Seed the random number generator to get different results each time
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 100000 and 999999 (6 digits)
	randomNumber := rand.Intn(900000) + 100000
	return randomNumber
}
