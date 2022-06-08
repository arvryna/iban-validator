package ibanparser

import (
	"errors"
	"math/big"
	"strconv"
	"unicode"

	"github.com/arvryna/iban-validator/internal/util"
)

var errorPrefix = "Invalid IBAN number, Reason: "
var ErrorInvalidIbanLength = errors.New(errorPrefix + " length should be between [5, 34]")
var ErrorIbanNonAlphaNumeric = errors.New(errorPrefix + "contains non-alphanumeric chars")
var ErrorInvalidIban = errors.New(errorPrefix + "remainder check failed.")

type IBANParser interface {
	FormattedIban() string
	Validate() error
}

type ibanParser struct {
	iban    string
	isValid bool
}

func Init(iban string) IBANParser {
	return &ibanParser{
		iban: util.TrimString(iban),
	}
}

func (p *ibanParser) Validate() error {
	isValidLength := len(p.iban) >= IbanMinSize && len(p.iban) <= IbanMaxSize

	if !isValidLength {
		return ErrorInvalidIbanLength
	}

	if !util.IsAlphaNumeric(p.iban) {
		return ErrorIbanNonAlphaNumeric
	}

	if p.ComputeIbanModulus() != ExpectedRemainder {
		p.isValid = false
		return ErrorInvalidIban
	}

	return nil
}

// https://www.wikiwand.com/en/International_Bank_Account_Number#/Algorithms
func (p *ibanParser) ComputeIbanModulus() int {
	iban := p.iban
	iban = iban[4:] + iban[0:4]
	res := ""
	for _, ch := range iban {
		if unicode.IsLetter(ch) {
			val := ch - 55
			res += strconv.Itoa(int(val))
		} else {

			res += string(ch)
		}
	}

	val, ok := new(big.Int).SetString(res, 10)
	if !ok {
		return -1
	}

	return int((new(big.Int).Mod(val, big.NewInt(97))).Int64())
}

func (p *ibanParser) FormattedIban() string {
	return p.iban
}
