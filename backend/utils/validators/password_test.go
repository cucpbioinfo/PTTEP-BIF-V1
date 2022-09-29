package validators

import (
	"fmt"
	"testing"
)

func TestValidPassword(t *testing.T) {
	const validPassword = "123123aA"
	isValidPassword := IsValidPassword(validPassword)

	if !isValidPassword {
		t.Error(fmt.Sprintf("%s should be valid password", validPassword))
	}
}

func TestInvalidPasswordLength(t *testing.T) {
	const length7Password = "123aAsd"
	const length17Password = "112233445566aaAAX"
	isLength7PasswordValid := IsValidPassword(length7Password)
	isLength17PasswordValid := IsValidPassword(length17Password)

	if isLength7PasswordValid {
		t.Error("Password of length 7 should be invalid")
	}

	if isLength17PasswordValid {
		t.Error("Password of length 17 should be invalid")
	}
}

func TestInvalidPasswordFormat(t *testing.T) {
	const numberPassword = "12345678"
	const lowerCasePassword = "abcdefgh"
	const upperCasePassword = "ABCDEFGH"
	const numberAndLowerCasePassword = "abcdefg1"
	const numberAndUpperCasePassword = "ABCDEFGH1"
	const lowerAndUpperCasePassword = "abcdEFGH"

	isNumberPasswordValid := IsValidPassword(numberPassword)
	isLowerCasePasswordValid := IsValidPassword(lowerCasePassword)
	isUpperCasePasswordValid := IsValidPassword(upperCasePassword)
	isNumberAndLowerCasePasswordValid := IsValidPassword(numberAndLowerCasePassword)
	isNumberAndUpperCasePasswordValid := IsValidPassword(numberAndUpperCasePassword)
	isLowerAndUpperCasePasswordValid := IsValidPassword(lowerAndUpperCasePassword)

	if isNumberPasswordValid {
		t.Error("Password with only number should be invalid")
	}

	if isLowerCasePasswordValid {
		t.Error("Password with only lower case should be invalid")
	}

	if isUpperCasePasswordValid {
		t.Error("Password with only upper case should be invalid")
	}

	if isNumberAndLowerCasePasswordValid {
		t.Error("Password with number and lower case should be invalid")
	}

	if isNumberAndUpperCasePasswordValid {
		t.Error("Password with number and upper case should be invalid")
	}

	if isLowerAndUpperCasePasswordValid {
		t.Error("Password with lower and upper case should be invalid")
	}

}
