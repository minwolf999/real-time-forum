package verificationfunction

/*
This function takes 1 argument:
	- a string

The objective of this function is to look if the argument is correctly formated.

The function gonna return:
	- a boolean
*/
func PasswordVerif(password string) bool {
	var isLongEnought bool = false
	var containUpper bool = false
	var containSpeChar bool = false
	var containNumber bool = false
	if len(password) >= 8 {
		isLongEnought = true
	}
	for _, r := range password {
		if r >= 'A' && r <= 'Z' {
			containUpper = true
		} else if r >= '0' && r <= '9' {
			containNumber = true
		} else if r < 'a' || r > 'z' {
			containSpeChar = true
		}
	}
	if isLongEnought && containNumber && containSpeChar && containUpper {
		return true
	}
	return false
}
