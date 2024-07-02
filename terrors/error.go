package terrors

import (
	"errors"
	"fmt"
	"log"
)

func Hello(name string) (string, error) {

	if(name == "") {
		return "Error", errors.New("empty name");
	}

	// logging
	log.SetPrefix("greetings: ")

	message:=fmt.Sprintf("Hi, %v. You are escaped from error", name);

	return message, nil;
}