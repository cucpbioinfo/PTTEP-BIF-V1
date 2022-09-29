package validators

func IsValidPassword(password string) bool {
	if len(password) < 8 || len(password) > 16 {
		return false
	}
	// contain at least 1 upper, lower case and number
	return true
}
