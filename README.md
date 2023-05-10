# randomizer

[![Coverage Status](https://coveralls.io/repos/github/imusmanmalik/randomizer/badge.svg?branch=main)](https://coveralls.io/github/imusmanmalik/randomizer?branch=main) [![CI Build](https://github.com/imusmanmalik/randomizer/actions/workflows/build.yaml/badge.svg)](https://github.com/imusmanmalik/randomizer/actions/workflows/build.yaml) [![CI Test](https://github.com/imusmanmalik/randomizer/actions/workflows/test.yaml/badge.svg)](https://github.com/imusmanmalik/randomizer/actions/workflows/test.yaml) [![CI Scan](https://github.com/imusmanmalik/randomizer/actions/workflows/scan.yaml/badge.svg)](https://github.com/imusmanmalik/randomizer/actions/workflows/scan.yaml) [![CI Release](https://github.com/imusmanmalik/randomizer/actions/workflows/release.yaml/badge.svg)](https://github.com/imusmanmalik/randomizer/actions/workflows/release.yaml) [![Go Reference](https://pkg.go.dev/badge/github.com/imusmanmalik/randomizer.svg)](https://pkg.go.dev/github.com/imusmanmalik/randomizer)

This is a GoLang library for generating cryptographically secure random numbers using the crypto/rand package. The library provides a simple API for generating random integers, bytes, and strings.

## Installation

To install the library, use the go get command:

```shell
go get github.com/imusmanmalik/randomizer
```

## Usage

To use the library, import it in your Go code:

```go
import (
    "fmt"
    "github.com/imusmanmalik/randomizer"
)

func main() {
	// Generate a random integer between 1 and 100
	randomInt, err := randomizer.RandomInt(1, 100)
	if err != nil {
		panic(err)
	}
	fmt.Println(randomInt)

	// Generate a random slice of bytes with length 16
	randomBytes, err := randomizer.RandomBytes(16)
	if err != nil {
		panic(err)
	}
	fmt.Println(randomBytes)

	// Generate a random string with length 32
	randomString, err := randomizer.RandomString(32)
	if err != nil {
		panic(err)
	}
	fmt.Println(randomString)
}
```

# Testing

```shell
go test
```

## Contributing

Contributions are welcome! If you find a bug or have an idea for a new feature, please open an issue or submit a pull request on GitHub.

## License

This library is licensed under Apache 2.0 License. See the [LICENSE file](https://github.com/imusmanmalik/randomizer/blob/main/LICENSE) for details.

## Acknowledgments

This library was inspired by the math/rand package in the Go standard library, and the github.com/Pallinder/go-randomdata library.
