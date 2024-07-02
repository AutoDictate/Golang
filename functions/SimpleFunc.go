package functions

import "fmt"

func SimpleFunction() {

	var a int = 10
	var b int = 10

	fmt.Println(a+b)

}

func ParameterizedFunction(a int, b int) {

	fmt.Println(a+b)
}

func ReturnFunction() int {
	return 10+10
}

func ReturnFunctionString () string {
	return "Hello, bro !!"
}

func ReturnFunctionNamed() (name string) {
	
	name = "Surya"

	return
}

func ReturnMultipleValues() (fname string, lname string) {
	fname = "Joseph"
	lname = "kurmila"

	return
}

var (

	// Anonymous Function
	area = func(a int, b int ) int {
		return a*b
	}
)

func AssignFunctionOnVariable() {
	fmt.Println(area(10, 8))
}

func SelfArguments() {

	func (a int, b int) {
		fmt.Println(a * b)
	}(10,20)
	
}



