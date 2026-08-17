package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/atotto/clipboard"
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
	toClipboard := flag.Bool("clip", false, "Copy generated password to clipboard")

	flag.Parse()

	// Generate the secure password
	password, err := generator.GeneratePassword(*length, *includeUpper, *includeNumbers, *includeSymbols, *customSymbols)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating password: %v\n", err)
		os.Exit(1)
	}

	// Copy to clipboard or print to screen
	if *toClipboard {
		err = clipboard.WriteAll(password)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to copy to clipboard: %v\n", err)
		} else {
			fmt.Fprintln(os.Stderr, "Password copied to clipboard.")
		}
	} else {
		fmt.Println(password)
	}
}
