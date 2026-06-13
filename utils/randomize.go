package utils

import "math/rand"

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomString(length int) string {
	code := make([]rune, length)

	for index := range code {
		code[index] = rune(letters[rand.Intn(len(letters))])
	}

	return string(code)
}
