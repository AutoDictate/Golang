package datatypeconversion

import (
	"fmt"
	"reflect"
	"strconv"
)

func FunctionExec()  {

	// StringToInteger()
	// StringToIntegerUsingParseInt()
	// StringToIntegerUsingSscan()
	StringToBoolean()
	
}

func StringToInteger() {

	strVar := "9999"

	intVar, err := strconv.Atoi(strVar)

	fmt.Println(intVar, err)
	fmt.Println("After the Atoi is executed: ",reflect.TypeOf(intVar))

}

func StringToIntegerUsingParseInt() {

	strVar := "9999"

	intVar, err := strconv.ParseInt(strVar, 0, 64)

	fmt.Println("---------StringToIntegerUsingParseInt---------")
	fmt.Println(intVar, err)
	fmt.Println("After the Atoi is executed: ",reflect.TypeOf(intVar))

}

func StringToIntegerUsingSscan()  {
	
	strVar := "200.00"

	intVar := 0

	_, err := fmt.Sscan(strVar, &intVar)

	fmt.Println(intVar, err, reflect.TypeOf(intVar))
}

func StringToBoolean() {

	strVar := "true"

	boolVar, _ := strconv.ParseBool(strVar)

	fmt.Printf("%T, %v", boolVar, boolVar)
}