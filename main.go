package main

import (
	"e/structs"
	"fmt"
)

func main() {

	user := &structs.User{}
	user = user.CreateUser()

	fmt.Println(*user)

}