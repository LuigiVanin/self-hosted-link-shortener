package shared

import (
	"math/rand"
	"strings"
)

type Pair[T1 any, T2 any] struct {
	First  T1
	Second T2
}

type JSON map[string]any

func GenerateRandomCharacters(length int) string {
	var result strings.Builder
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	count := len(characters)

	for range length {
		index := rand.Intn(count)
		character := string(characters[index])
		result.WriteString(character)
	}

	return result.String()
}
