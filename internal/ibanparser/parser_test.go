package ibanparser

import "testing"

func TestIbanLength(t *testing.T) {
	iban := Init("12345123451234512345123451234512345123451234512345123451234512345")
	err := iban.Validate()
	if err != ErrorInvalidIbanLength {
		t.Error("Iban length validation failed")
	}
}

func TestIbanAlphaNumeric(t *testing.T) {
	iban := Init("大あ Дракад")
	err := iban.Validate()
	if err != ErrorIbanNonAlphaNumeric {
		t.Error("Iban alphanumeric test failed", err)
	}

}
