package luhn

import (
	"testing"
)

func TestGetControlDigit(t *testing.T) {
	control, _ := GetControlDigit("012345")

	if control != 5 {
		t.Errorf("invalid control digit")
	}
}

func TestIsValid(t *testing.T) {
	valid, _ := IsValid("0123455")

	if !valid {
		t.Errorf("should be valid")
	}

	valid, _ = IsValid("0123456")

	if valid {
		t.Errorf("should be invalid")
	}
}
