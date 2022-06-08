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

func TestValidIbans(t *testing.T) {
	ibans := [6]string{
		"FR14 2004 1010 0505 0001 3M02 606",
		"SE45 5000 0000 0583 9825 7466",
		"DE89 3704 0044 0532 0130 00",
		"NL91 ABNA 0417 1643 00",
		"IT60 X054 2811 1010 0000 0123 456",
		"BR97 0036 0305 0000 1000 9795 493P 1",
	}

	for _, iban := range ibans {
		parser := Init(iban)
		err := parser.Validate()
		if err != nil {
			t.Error("Iban validation failed", err)
		}
	}
}

func TestInValidIbans(t *testing.T) {
	ibans := [6]string{
		"FR14 2004 1010 0505 0001 4M02 606",
		"SE45 5000 1100 0583 9825 7466",
		"DE89 3704 0044 0532 0120 00",
		"NL91 ABNA 0417 1533 00",
		"IT60 X054 2811 1010 0000 0133 456",
		"BR97 0036 0305 0000 1000 9885 493P 1",
	}

	for _, iban := range ibans {
		parser := Init(iban)
		err := parser.Validate()
		if err == nil {
			t.Error("Iban:", iban, "validation failed, reason: ", err)
		}
	}
}
