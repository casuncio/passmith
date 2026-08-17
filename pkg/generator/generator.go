package generator

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	lowerChars  = "abcdefghijklmnopqrstuvwxyz"
	upperChars  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars = "0123456789"
	symbolChars = "!@#$%^&*()_+-=[]{}|;:,.<>?"
)

// GeneratePassword generates a secure password based on specified options.
func GeneratePassword(length int, useUppers bool, useNumbers bool, useSymbols bool, customSymbols string) (string, error) {

	// Build the character pool and required sets based on flags
	chars := lowerChars
	requiredSets := []string{lowerChars}
	if useUppers {
		chars += upperChars
		requiredSets = append(requiredSets, upperChars)
	}
	if useNumbers {
		chars += numberChars
		requiredSets = append(requiredSets, numberChars)
	}
	if customSymbols != "" {
		chars += customSymbols
		requiredSets = append(requiredSets, customSymbols)
	} else if useSymbols {
		chars += symbolChars
		requiredSets = append(requiredSets, symbolChars)
	}

	if length <= len(requiredSets) {
		return "", fmt.Errorf("Password length must be at least %d to include all selected character types\n", len(requiredSets))
	}

	return generatePassword(length, chars, requiredSets)
}

// generatePassword uses crypto/rand for cryptographically secure randomness
// and guarantees at least one character from each required set.
func generatePassword(length int, chars string, requiredSets []string) (string, error) {
	bytes := make([]byte, length)
	charLen := big.NewInt(int64(len(chars)))

	// Guarantee one character from each required set
	used := make(map[int]bool)
	for _, set := range requiredSets {
		setLen := big.NewInt(int64(len(set)))
		num, err := rand.Int(rand.Reader, setLen)
		if err != nil {
			return "", err
		}
		pos, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
		if err != nil {
			return "", err
		}
		p := int(pos.Int64())
		// Avoid collisions with already-placed guaranteed characters
		for used[p] {
			p = (p + 1) % length
		}
		used[p] = true
		bytes[p] = set[num.Int64()]
	}

	// Fill remaining positions from the full pool
	for i := 0; i < length; i++ {
		if used[i] {
			continue
		}
		num, err := rand.Int(rand.Reader, charLen)
		if err != nil {
			return "", err
		}
		bytes[i] = chars[num.Int64()]
	}

	// Fisher-Yates shuffle with crypto/rand
	for i := length - 1; i > 0; i-- {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return "", err
		}
		bytes[i], bytes[j.Int64()] = bytes[j.Int64()], bytes[i]
	}

	return string(bytes), nil
}
