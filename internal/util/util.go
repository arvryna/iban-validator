package util

import (
	"regexp"
	"strings"
)

func TrimString(str string) string {
	return strings.ReplaceAll(str, " ", "")
}

func IsAlphaNumeric(str string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(str)
}
