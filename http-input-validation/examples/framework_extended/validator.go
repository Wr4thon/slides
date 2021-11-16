package main

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

//START OMIT
func usernameLen(fl validator.FieldLevel) bool {
	l := len(fl.Field().String())
	return l >= 7 && l < 32
}

func usernameCharacterSet(fl validator.FieldLevel) bool {
	val := fl.Field().String()

	for i := range val {
		v := rune(val[i])
		if !unicode.IsDigit(v) && !unicode.IsLetter(v) {
			return false
		}
	}

	return true
}

//END OMIT
