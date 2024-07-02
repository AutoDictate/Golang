package tvariables

import "fmt"

var (
	name    = "Surya"
	age     = 21
	address = "Salem"
	CGPA    = 8.8
)

func Test_Variable() {

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(address)
	fmt.Println(CGPA)

}

func Value() {
	
	var num int

	var name string

	var rupee float32

	var male bool

	fmt.Println(num)
	fmt.Println(name)
	fmt.Println(rupee)
	fmt.Println(!male)
}

const (
	cname = "Surya"
	cage = 21
	caddress = "Salem"
)
func ConstVariable() {

	fmt.Printf("Constant name : %v",cname)
	fmt.Println()
	fmt.Printf("Constant age : %v",cage)
	fmt.Println()
	fmt.Printf("Constant address : %v",caddress)
	fmt.Println()
	
}

func MultipleVariable() {

	var name, age string = "Kyros", "108"

	fmt.Println(name)
	fmt.Println(age)
}