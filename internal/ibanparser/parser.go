package ibanparser

import "github.com/arvryna/iban-validator/internal/util"

const IbanMaxSize = 34

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
	if len(p.iban) > IbanMaxSize {
		return ErrorLongIban
	}

	if !util.IsAlphaNumeric(p.iban) {
		return ErrorIbanNonAlphaNumeric
	}
	return nil
}

func (p *ibanParser) FormattedIban() string {
	return p.iban
}
