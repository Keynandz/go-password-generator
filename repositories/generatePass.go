package repositories

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	lettersDown  = "abcdefghijklmnopqrstuvwxyz"
	lettersUp    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits       = "0123456789"
	specialChars = "!@#$%^&*()"
)

func GeneratePassword(limit, up, low, digit, symbol int) (string, error) {
	if limit < 1 {
		return "", fmt.Errorf("password length must be at least 1")
	}

	if up+low+digit+symbol == 0 {
		return "", nil
	}

	var password strings.Builder

	if up != 0 {
		password.WriteByte(lettersUp[rand.Intn(len(lettersUp))])
	}

	if low != 0 {
		password.WriteByte(lettersDown[rand.Intn(len(lettersDown))])
	}

	if digit != 0 {
		password.WriteByte(digits[rand.Intn(len(digits))])
	}

	if symbol != 0 {
		password.WriteByte(specialChars[rand.Intn(len(specialChars))])
	}

	for password.Len() < limit {
		switch rand.Intn(4) {
		case 0:
			if up != 0 {
				password.WriteByte(lettersUp[rand.Intn(len(lettersUp))])
			}
		case 1:
			if low != 0 {
				password.WriteByte(lettersDown[rand.Intn(len(lettersDown))])
			}
		case 2:
			if digit != 0 {
				password.WriteByte(digits[rand.Intn(len(digits))])
			}
		case 3:
			if symbol != 0 {
				password.WriteByte(specialChars[rand.Intn(len(specialChars))])
			}
		}
	}

	result := password.String()

	rand.Seed(time.Now().UnixNano())
	runeSlice := []rune(result)
	rand.Shuffle(len(runeSlice), func(i, j int) {
		runeSlice[i], runeSlice[j] = runeSlice[j], runeSlice[i]
	})

	shuffledPassword := string(runeSlice)
	return shuffledPassword, nil
}

func PasswordStrength(password string) string {
	var complexity int
	length := len(password)

	uniqueChars := make(map[rune]bool)
	for _, char := range password {
		uniqueChars[char] = true
	}
	uniqueCount := len(uniqueChars)

	if length >= 12 && uniqueCount >= 4 {
		complexity = 3
	} else if length >= 8 && uniqueCount >= 3 {
		complexity = 2
	} else {
		complexity = 1
	}

	switch complexity {
	case 3:
		return "Strong"
	case 2:
		return "Medium"
	default:
		return "Weak"
	}
}
