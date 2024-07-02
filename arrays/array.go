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

func CreateArray() {
	
	// Create and Initialize the Array 

	var intArray [5] int

	intArray[0] = 15    
	intArray[1] = 14    
	intArray[2] = 13    
	intArray[3] = 12    
	intArray[4] = 11

	for _, val := range intArray {
		fmt.Println(val)
	}

}