package functions

import "fmt"

func DeferFunction() {

	for i := 0; i < 11; i++ {
		defer fmt.Println(i)
	}
}