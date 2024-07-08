package structs

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) greet() {
	fmt.Println("Hello, ", p.name)
}

func StructInitialize() {

	person1 := Person{"Surya", 21}

	fmt.Println(person1)

	var person2 = Person{
		"Yuvaraj",
		21,
	}

	fmt.Println(person2)
	AccessStructs(person1)
	AccessStructs(person2)
}

func AccessStructs(person Person) {

	fmt.Println("Person Name : ", person.name)
	fmt.Println("Person age : ", person.age)
	person.greet()
}