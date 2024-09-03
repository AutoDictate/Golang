package looping

import "fmt"

func ForLoop() {

	fmt.Println("Type 1")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	
	fmt.Println("Type 2")
	var i = 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("Type 3")
	for i := 1; ; i++ {
		fmt.Println(i)
		if i == 10 {
			break
		}
	}

}

func ForRange() {

	for range "Hello" {
		fmt.Println("Hello")
	}

}
