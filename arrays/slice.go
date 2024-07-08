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

func RemoveSlice() {

	names := []string{"1. surya", "2. yuvaraj", "3. sowmin", "4. aswin"}

	names = RemoveIndex(names, 3)

	fmt.Println(names)
}

func RemoveIndex(s []string, i int) []string {

	return append(s [:i], s [i+1:]...)
	
}

func AppendToExistingSlice() {

	var slice1 = []string {"ironman", "antman", "captain amaerica"}
	var slice2 = []string {"deadpool", "wolverine"}

	slice1 = append(slice1, slice2...)
	
	for i, value := range slice1 {
		fmt.Printf("%d. %v\n",i+1, value)
	}
}

func CheckIfItemExists() {

	var slice1 = []int {1,2,3,4,5}

	var check int = 60000
	var isHere bool = false
	
	for _, value := range slice1{
		if  check == value{
			isHere = true
		}
	}

	if isHere {
		fmt.Println("The value is here !!!")
	}else {
		fmt.Println("Oops !!! value not found")
	}
}