package util

import "regexp"

func IsAlphaNumeric(str string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(str)
}
