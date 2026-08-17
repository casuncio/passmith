# 🛡️ Passmith

A lightweight, cryptographically secure password generator forged in Go.

Passmith is designed to be a fast, reliable CLI tool for generating high-entropy passwords directly in your terminal. It uses Go's `crypto/rand` package to ensure that your passwords are suitable for production security requirements.

## Features

*   **Cryptographically Secure**: Uses `crypto/rand` for robust randomness.
*   **Character Type Guarantee**: Ensures at least one character from every enabled category (uppercase, numbers, symbols) appears in every generated password.
*   **Fast & Portable**: Compiled into a single, lightning-fast binary.
*   **Simple Syntax**: Designed to be quick to use in any terminal workflow.

## Installation

### Prerequisites

*   [Go](https://go.dev/doc/install) (1.18 or later recommended)

### Install via Go

If you have your Go environment configured, you can install the latest version directly:

```bash
go install [github.com/casuncio/passmith@latest](https://github.com/casuncio/passmith@latest)
```

Ensure your `~/go/bin` directory is in your system `PATH`.

## Usage

Run `passmith` with your preferred options:

```bash
# Generate a default 16-character password
passmith

# Generate a 24-character password without symbols
passmith -length=24 -symbols=false

# Generate a 12-character password using only letters and numbers
passmith -length=12 -symbols=false -upper=true

```

### Available Flags

| Flag | Default | Description |
| --- | --- | --- |
| `-length` | `32` | Total length of the generated password |
| `-upper` | `true` | Include uppercase letters (A-Z) |
| `-numbers` | `true` | Include numbers (0-9) |
| `-symbols` | `true` | Include special symbols (!@#$%^&*) |
| `-custom-symbols` | `""` | Override default symbols with custom set |

> **Note**: When a character type is enabled, the generated password is guaranteed to contain at least one character of that type. The password length must be at least equal to the number of enabled character types.

## Development

To build the tool from source:

1. Clone the repository:
```bash
git clone [https://github.com/casuncio/passmith.git](https://github.com/casuncio/passmith.git)
cd passmith

```


2. Build the binary:
```bash
go build -o passmith main.go

```


3. Run it:
```bash
./passmith

```


## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
