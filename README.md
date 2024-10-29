# ValidCNPJ Function

This project contains a Go implementation of the `ValidCNPJ` function, which validates Brazilian CNPJ numbers. CNPJ (Cadastro Nacional da Pessoa Jurídica) is a unique identifier issued to Brazilian companies, similar to a business tax ID number in other countries. The function is designed to determine whether a given CNPJ is valid, regardless of whether it contains formatting characters such as dots, dashes, or slashes.

## Code Overview

The code consists of two functions:

1. **`ValidCNPJ(cnpj string) bool`**: The main function to validate a CNPJ number.
2. **`onlyLetterNumber(cnpj string) string`**: A helper function to remove any formatting characters from the input.

### `ValidCNPJ` Function

The `ValidCNPJ` function takes a CNPJ as a string input and performs the following steps:

1. **Remove Formatting Characters**: It uses the helper function `onlyLetterNumber` to remove all non-numeric characters from the input string. This ensures that the validation process only works with the numerical part of the CNPJ.

2. **Check Length**: It verifies that the cleaned CNPJ has a length of 14 characters, as a valid CNPJ must contain 14 digits.

3. **Calculate First Verification Digit**:
   - The first 12 digits are used to calculate the first verification digit.
   - The sum of the product of each digit and a weight is calculated, with the weight starting at 5 and cycling down to 2, then restarting at 9.
   - The remainder of the sum divided by 11 is calculated, and the first verification digit (`digit1`) is determined based on this remainder.

4. **Calculate Second Verification Digit**:
   - The first 13 digits (including the calculated `digit1`) are used to calculate the second verification digit (`digit2`), following a similar process as for `digit1`, but with the weight starting at 6.

5. **Compare Verification Digits**: Finally, the function compares the calculated verification digits (`digit1` and `digit2`) with the last two digits of the input CNPJ. If both digits match, the function returns `true`, indicating that the CNPJ is valid; otherwise, it returns `false`.

### `onlyLetterNumber` Function

The `onlyLetterNumber` function is a helper that removes any formatting characters from the CNPJ string, such as periods (`.`), slashes (`/`), and dashes (`-`). This ensures that the validation function only works with the numerical digits.

## Usage

To use this code to validate a CNPJ:

1. Import the `ValidCNPJ` function in your Go project.
2. Call the function with a CNPJ string (formatted or unformatted).

Example:

```go
package main

import (
	"fmt"
)

func main() {
	cnpj := "12.345.678/0001-95"
	isValid := ValidCNPJ(cnpj)
	if isValid {
		fmt.Println("The CNPJ is valid.")
	} else {
		fmt.Println("The CNPJ is invalid.")
	}
}
```

## Notes

- The CNPJ validation relies on the algorithm used by the Brazilian Federal Revenue to validate CNPJs.
- The function handles both formatted and unformatted CNPJ strings by stripping out any non-numeric characters.
- Make sure that the input CNPJ has 14 digits, otherwise the function will return `false`.

## Conclusion

This project provides a utility for validating Brazilian CNPJ numbers, using a simple and efficient Go implementation. It helps ensure that the CNPJ is correctly formatted and contains valid verification digits according to the official validation rules.
