package conditionalstatements

import "fmt"

func FunctionExec() {

	ifInitialization()

}

func ifInitialization() {

	if x := 100; x > 10 {
		fmt.Println("Greater")
	}else {
		fmt.Println("Smaller")
	}

}