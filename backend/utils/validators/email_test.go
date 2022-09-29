package validators

import (
	"fmt"
	"testing"
)

func TestValidEmail(t *testing.T) {
	const validEmail = "poom007.sk135@gmail.com"
	isValidEmail := IsValidEmail(validEmail)

	if !isValidEmail {
		t.Error(fmt.Sprintf("%s should be a valid email", validEmail))
	}
}

func TestInvalidEmailLength(t *testing.T) {
	const length2Email = "p2"
	const length255Email = "nrlsbcnhmfcfbdxlhksguqadqhqinarurkenaamghqlmrnnwgnahrnpxiuulqghyvnzofpgdsdxxlnsxxcznciyofdaoiumpqgbesojpdeosomukgatoevndawjbdfiiukobzrdffbfdhfmmyquwuugydoysdtcyqzlegamkmqecnqcqaiqzuubbcblcjursuysffupcadixkimuriehjbxwbndbkppvzdsmjhjyuotqdakuxulwx@gmail.com"

	isLength2EmailValid := IsValidEmail(length2Email)
	isLength255EmailValid := IsValidEmail(length255Email)

	if isLength2EmailValid {
		t.Error("Email of length 2 should be invalid")
	}
	if isLength255EmailValid {
		t.Error("Email of length 255 should be invalid")
	}
}

func TestInValidEmailFormat(t *testing.T) {
	const invalidEmailFormat = "po@e.@gmail.com"

	isValidEmail := IsValidEmail(invalidEmailFormat)

	if isValidEmail {
		t.Error("Email with invalid format should be invalid")
	}
}
