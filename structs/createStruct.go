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

func (u *User) CreateUser() *User {

	u.name = "Surya"
	u.age = 21
	u.email = "surya@gmail.com"
	u.address = "Salem"
	u.Role.admin = true
	u.Role.manager = false

	fmt.Println(u)

	return u
}

func CreateUserUsingStructLiteral() {

	user2 := User{
		name:    "Surya",
		age:     21,
		address: "Chennai",
		email:   "Zoro@gmail.com",
		Role: struct {
			admin   bool
			manager bool
		}{
			admin:   true,
			manager: false,
		},
	}

	fmt.Println(user2)
}
