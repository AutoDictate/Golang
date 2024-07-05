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

	for i, val := range intArray {
		
		fmt.Println(i , val)
	}

}

func InitializeArray() {

	x := [5]int{1,2,3,4}

	for _, v := range x {

		if v==0 {
			fmt.Println("Ended with no value")
		}else{
			fmt.Println(v)
		}
	}

}

func InitializeArrayWithoutSpecifyingLength() {

	x := [...] int {1,2,3,4,5}

	for _, val := range x {
		fmt.Println(val)
	}

	fmt.Println("Length of the Array (x) : ",len(x))

	fmt.Println(reflect.TypeOf(x).Kind())

}

func ItemExistsCheck() {

	ai_names := [] string {"surya", "yuvaraj", "sathya"}
	
	fmt.Println(itemExists(ai_names, "dharani"))
	fmt.Println(itemExists(ai_names, "sowmin"))
	fmt.Println(itemExists(ai_names, "aswin"))

	fmt.Println(itemExists(ai_names, "surya"))
	fmt.Println(itemExists(ai_names, "sathya"))
	fmt.Println(itemExists(ai_names, "aswin"))
	
}

func itemExists(arrayType []string, item string) bool {

	arr := arrayType

	for i:=0; i< len(arr); i++ {
		if arr[i]== item {
			return true
		}
	}

	return false
	
}