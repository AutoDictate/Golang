package arrays

import "fmt"

func Slice() {

	names := []string{"1. surya", "2. yuvaraj", "3. sowmin", "4. aswin"}

	// Here, Array doesn't has specified with Size.
	// So, It is considered as a "SLICE"

	// ASCENDING ORDER
	fmt.Println("Printing names in a Chronological Order")
	for _, value := range names {
		fmt.Println(value)
	}

	// DESCENDING ORDER
	fmt.Println("Printing names in a Reverse Order")
	for _, value := range names {
		defer fmt.Println(value)
	}
}