package verificationfunction

/*
This function takes 2 arguments:
	- a array of string
	- a string

The objective of this function is to look if the string is in the array.

The function gonna return:
	- a boolean
*/
func TabNotContain(tab []string, s string) bool {
	for _, v := range tab {
		if s == v {
			return false
		}
	}

	return true
}
