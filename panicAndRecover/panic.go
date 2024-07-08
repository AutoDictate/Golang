package panicandrecover

import "fmt"

func Panic() {

	fmt.Println("This is first statement")

	panic("This is the panic statement")

	// The statement after the panic statement will not be executed
	// fmt.Println("This is the last statement")
}

func Division(num1 int, num2 int) (result int){

	defer HandlePanic()
	
	if num2== 0 {
		panic("Cannot divide by zero")
	}
	result = num1 / num2
	return
}

func HandlePanic() {
	
	a:= recover()

	if a!=nil {
		fmt.Println("Recover : ", a)
	}
}