package validators

import "regexp"

func IsEmailValid(email string) bool {
	var rxEmail = regexp.MustCompile("/^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$/g")

	if len(email) < 3 || len(email) > 254 || !rxEmail.MatchString(email) {
		return false;
	}

	return true
}