package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/casuncio/passmith/pkg/generator"
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

	// Generate the secure password
	password, err := generator.GeneratePassword(*length, *includeUpper, *includeNumbers, *includeSymbols, *customSymbols)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating password: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(password)
}
