package verificationfunction

import (
	"strings"
)

/*
This function takes 1 argument:
	- a string

The objective of this function is to look if the argument is correctly formated.

The function gonna return:
	- a boolean
*/
func IsValidMessage(s string) bool {
	tmp := s

	tmp = strings.ReplaceAll(tmp, " ", "")
	tmp = strings.ReplaceAll(tmp, string(rune(10)), "")
	tmp = strings.ReplaceAll(tmp, string(rune(13)), "")

	return tmp != ""
}
