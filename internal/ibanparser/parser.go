package ibanparser

import (
	"errors"

	"github.com/arvryna/iban-validator/internal/util"
)

var errorPrefix = "Invalid IBAN number, Reason: "
var ErrorInvalidIbanLength = errors.New(errorPrefix + " length should be between [5, 34]")
var ErrorIbanNonAlphaNumeric = errors.New(errorPrefix + "contains non-alphanumeric chars")

type IBANParser interface {
	FormattedIban() string
	Validate() error
}

type ibanParser struct {
	iban string
}

func Init(iban string) IBANParser {
	return &ibanParser{iban: iban}
}

func (p *ibanParser) Validate() error {
	p.iban = util.TrimString(p.iban)

	isValidLength := len(p.iban) >= IbanMinSize && len(p.iban) <= IbanMaxSize

	if !isValidLength {
		return ErrorInvalidIbanLength
	}

	if !util.IsAlphaNumeric(p.iban) {
		return ErrorIbanNonAlphaNumeric
	}

	return nil
}

func (p *ibanParser) FormattedIban() string {
	return p.iban
}
