package main

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

func main() {
	v := validator.New()
	//START OMIT
	v.RegisterValidation("usernameLen", usernameLen, false)
	v.RegisterValidation("usernameCharacterSet", usernameCharacterSet, false)
	input := struct {
		UserName string `validate:"required,usernameLen,usernameCharacterSet"`
	}{
		UserName: "",
		//END OMIT
	}

	if err := v.Struct(input); err != nil {
		println(err.Error())
		return
	}

	println("username is valid")
}

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
