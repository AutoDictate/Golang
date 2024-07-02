package arrays

import (
	"fmt"
	"reflect"
)

func SimpleArray() {

	var intArray [5]int
	var stringArray [5]string

	fmt.Println(reflect.TypeOf(intArray).Kind())
	fmt.Println(reflect.TypeOf(stringArray).Kind())
}