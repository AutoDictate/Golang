package structs

import "fmt"

type User struct {
	name    string
	age     int
	address string
	email   string

	Role struct {
		admin   bool
		manager bool
	}
}

func CreateUser() {

	var user1 User

	user1.name = "Surya"
	user1.age = 21
	user1.email = "surya@gmail.com"
	user1.address = "Salem"
	user1.Role.admin = true
	user1.Role.manager = false

	fmt.Println(user1)
}

func CreateUserUsingStructLiteral() {

	var user2 = User{
		name: "Surya",
		age: 21,
		address: "Chennai",
		email: "Zoro@gmail.com",
		Role: struct{admin bool; manager bool}{
			admin: true,
			manager: false,
		},
	}

	fmt.Println(user2)
}