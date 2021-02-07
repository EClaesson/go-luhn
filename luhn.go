package luhn

import (
	"fmt"
	"strconv"
)

func stringToDigits(str string) ([]int, error) {
	digits := make([]int, len(str))

	for i, val := range str {
		digit, err := strconv.Atoi(string(val))

		if err != nil || digit > 9 {
			return nil, fmt.Errorf("non-digit rune: '%s'", string(val))
		}

		digits[i] = digit
	}

	return digits, nil
}

// GetControlDigit calculates the control digit of a string of digits.
func GetControlDigit(num string) (int, error) {
	digits, err := stringToDigits(num)

	if err != nil {
		return 0, err
	}

	mult := 2
	valStr := ""

	for i := range digits {
		digit := digits[len(digits)-i-1]
		valStr += strconv.Itoa(digit * mult)

		if mult == 1 {
			mult = 2
		} else {
			mult = 1
		}
	}

	valDigits, _ := stringToDigits(valStr)

	valSum := 0

	for _, val := range valDigits {
		valSum += val
	}

	control := 10 - (valSum % 10)

	if control == 10 {
		control = 0
	}

	return control, nil
}

// IsValid checks whether the passed control digit (last digit in string) is correct.
func IsValid(num string) (bool, error) {
	control, _ := strconv.Atoi(string(num[len(num)-1]))
	number := num[:len(num)-1]

	correct, err := GetControlDigit(number)

	if err != nil {
		return false, err
	}

	return control == correct, nil
}
