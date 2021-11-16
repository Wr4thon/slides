package main

import "regexp"

func main() {
	var userNameRegex = regexp.MustCompile(`^[\w\d]{7,32}$`)

	input := struct {
		UserName string
	}{
		UserName: "",
	}

	if input.UserName == "" {
		println("username must be set")
		return
	}

	if !userNameRegex.Match([]byte(input.UserName)) {
		println("username does not match requirements")
		return
	}

	println("username is valid")
}
