package ibanparser

import (
	"errors"

	"github.com/arvryna/iban-validator/internal/util"
)

const IbanMinSize = 5
const IbanMaxSize = 34

var errorPrefix = "Invalid IBAN number, Reason: "
var ErrorInvalidIbanLength = errors.New(errorPrefix + " lengh should be between [5, 34]")
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
	isValidLengh := len(p.iban) >= IbanMinSize && len(p.iban) <= IbanMaxSize

	if !isValidLengh {
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
