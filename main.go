package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
)

const (
	lowerChars  = "abcdefghijklmnopqrstuvwxyz"
	upperChars  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars = "0123456789"
	symbolChars = "!@#$%^&*()_+-=[]{}|;:,.<>?"
)

func main() {
	// Usage message
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of passmith:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  A secure password generator forged in Go.\n\n")
		flag.PrintDefaults()
		fmt.Fprintf(flag.CommandLine.Output(), "\nNote: Each enabled character type is guaranteed to appear at least once.\n")
	}

	// Define command-line flags
	length := flag.Int("length", 32, "Length of the password")
	includeUpper := flag.Bool("upper", true, "Include uppercase letters")
	includeNumbers := flag.Bool("numbers", true, "Include numbers")
	includeSymbols := flag.Bool("symbols", true, "Include special symbols")
	customSymbols := flag.String("custom-symbols", "", "Override default symbol set i.e. \"$&_-+=\"")

	flag.Parse()

	if *length <= 0 {
		fmt.Fprintln(os.Stderr, "Error: Password length must be greater than 0")
		os.Exit(1)
	}

	// Build the character pool and required sets based on flags
	chars := lowerChars
	requiredSets := []string{lowerChars}
	if *includeUpper {
		chars += upperChars
		requiredSets = append(requiredSets, upperChars)
	}
	if *includeNumbers {
		chars += numberChars
		requiredSets = append(requiredSets, numberChars)
	}
	if *customSymbols != "" {
		chars += *customSymbols
		requiredSets = append(requiredSets, *customSymbols)
	} else if *includeSymbols {
		chars += symbolChars
		requiredSets = append(requiredSets, symbolChars)
	}

	if *length < len(requiredSets) {
		fmt.Fprintf(os.Stderr, "Error: Password length must be at least %d to include all selected character types\n", len(requiredSets))
		os.Exit(1)
	}

	// Generate the secure password
	password, err := generatePassword(*length, chars, requiredSets)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating password: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(password)
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
