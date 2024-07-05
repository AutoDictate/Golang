package reference

import (
	"fmt"
	"reflect"
)

func CopyArray() {

	names := [3]string{"Surya", "Yuvaraj", "Sowmiyan"}

	tk_names := names	// Here, tk_names directly takes the value from the "names" Array

	techno_names := &tk_names // Here, techno_names is a pointer holing the address of a variable "tk_names"

	for _, v := range names {
		fmt.Println(v)
	}

	fmt.Println(tk_names)
	fmt.Println(reflect.TypeOf(techno_names).Kind())
}