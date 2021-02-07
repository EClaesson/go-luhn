# go-luhn
Luhn control digit calculation for Go.

## About the Luhn algorithm
> The Luhn algorithm or Luhn formula, also known as the "modulus 10" or "mod 10" algorithm, named after its creator, IBM scientist Hans Peter Luhn, is a simple checksum formula used to validate a variety of identification numbers, such as credit card numbers, IMEI numbers, National Provider Identifier numbers in the United States, Canadian Social Insurance Numbers, Israeli ID Numbers, South African ID Numbers, Greek Social Security Numbers (ΑΜΚΑ), and survey codes appearing on McDonald's, Taco Bell, and Tractor Supply Co. receipts. It is described in U.S. Patent No. 2,950,048, filed on January 6, 1954, and granted on August 23, 1960.
>
> The algorithm is in the public domain and is in wide use today. It is specified in ISO/IEC 7812-1.[1] It is not intended to be a cryptographically secure hash function; it was designed to protect against accidental errors, not malicious attacks. Most credit cards and many government identification numbers use the algorithm as a simple method of distinguishing valid numbers from mistyped or otherwise incorrect numbers. 
>
> – Wikipedia

## Usage
```go
package main

import (
    luhn "github.com/EClaesson/go-luhn"
)

func main() {
    control, err := luhn.GetControlDigit("012345")
    // control is 5

    valid, err := luhn.IsValid("0123455")
    // valid is true

    valid, err := luhn.IsValid("0123456")
    // valid is false
}
```
