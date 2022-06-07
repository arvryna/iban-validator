package ibanparser

import "errors"

var errorPrefix = "Invalid IBAN number, Reason: "
var ErrorLongIban = errors.New(errorPrefix + " lengh > 34 ")
var ErrorIbanNonAlphaNumeric = errors.New(errorPrefix + "contains non-alphanumeric chars")
