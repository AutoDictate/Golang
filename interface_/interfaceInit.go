package interface_

import "fmt"

/*
1. An interface is an abstract type
2. Interface describes all the methods of a method set and
	provides the signatures for each method.
*/

type Student interface {
	PrintName(name string)
	PrintAddress(i int)
	PrintSalary(b int, t int) int
}

type Stu struct{} // user-defined Type -> Class that is going to implement Student interface

// PrintAddress implements Student.
func (e Stu) PrintAddress(i int) {
	panic("unimplemented")
}

func (e Stu) PrintName(name string) {
	fmt.Println("value of e : ", e)
	fmt.Println("Value of name : ", name)
}

func (e Stu) PrintSalary(b int, t int) int {
	var Salary = (b * t) / 100
	return b - Salary
}

func CreateInterface() {
	var s1 Stu
	s1.PrintName("surya")
	res := s1.PrintSalary(3500, 1000)
	fmt.Println(res)

}
