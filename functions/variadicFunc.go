package functions

import "fmt"

func VariadicFunc(){

	// VariadicExample("surya", "selva", "moha")

	// VariadicExample1()
	// VariadicExample1("your name", "weathering with you")
	// VariadicExample1("one piece", "naruto", "black clover")

	// for loop function 
		val := LoopFunc("Rectangle", 20, 30)
		fmt.Println(val)

}

func VariadicExample(s ...string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

func VariadicExample1(s ...string) {
	fmt.Println(s)
}

func LoopFunc(str string, y ...int) int {
	area := 1

	for _, val := range y {
		if str == "Rectangle" {
			area *= val;
		} else {
			area = val * val
		}
	}

	return area;
}