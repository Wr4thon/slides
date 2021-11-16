package main

import "github.com/go-playground/validator/v10"

//START OMIT
func main() {
	v := validator.New()
	input := struct {
		UserName string `validate:"required"`
	}{
		UserName: "",
	}

	if err := v.Struct(input); err != nil {
		println(err.Error())
		return
	}

	println("username is valid")
}

//END OMIT
