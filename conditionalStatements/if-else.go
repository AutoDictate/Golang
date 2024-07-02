package conditionalstatements

import "fmt"

func Vote() {

	var age int = 1

	const name string = "Surya"

	var isEligible bool = false

	if age > 18 {
		isEligible = true
	}

	if isEligible {
		fmt.Println("You are ready to Vote")
	} else {
		fmt.Println("You are not ready to Vote")
	}
}