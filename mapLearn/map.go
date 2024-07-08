package maplearn

import (
	"fmt"
	"sort"
)

func MapInitialize() {

	var employee = map[string]string{
		"firstName": "Surya",
		"lastName":  "A",
		"address":   "Salem",
		"email":     "surya@gmail.com"}

	fmt.Println(len(employee))
	EmptyMap()
	MapDeclareUsingMake()
	IterateMap(employee)
	TruncateMap(employee)

}

func EmptyMap() {

	var student = map[string] int {}

	fmt.Println(student)
}

func MapDeclareUsingMake()  {
	
	var student = make(map[string]int)

	student["rollno"] = 39
	student["age"] = 21

	fmt.Println(student)

	// Prints the length of the Map function
	fmt.Println(len(student))
	// DeleteKeyInMap(student)
}

func DeleteKeyInMap(student map[string]int)  {
	
	delete(student, "age")
	fmt.Println(student)
}

func IterateMap(employee map[string]string)  {
	
	for key, value := range employee {
		fmt.Printf("%v : %v \n", key, value)
	}

}

func TruncateMap(employee map[string]string) {
	
	for key := range employee {
		delete(employee, key)
	}

	fmt.Println(employee)
}

func SortMapKeys() {
	
	unSortedMap := map[string]int{"India": 20, "Canada": 70, "Germany": 15}
	
	keys := make([] string, 0, len(unSortedMap))
	
	for k := range unSortedMap {
		keys = append(keys, k)
	}
	
	sort.Strings(keys)
	
	for _, k := range keys {
		fmt.Println(unSortedMap[k])
	}
}

func SortMapValues() {
	unSortedMap := map[string]int{"India": 20, "Canada": 70, "Germany": 15}
	

	values := make([]int, 0, len(unSortedMap))

	for _, val := range unSortedMap {
		values = append(values, val)
	}

	sort.Ints(values)

	for i := 0; i < len(values); i++ {
		fmt.Println(values[i])
	}
}