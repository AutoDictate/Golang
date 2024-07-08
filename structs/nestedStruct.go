package structs

import "fmt"

type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	name          string
	age           int
	address       string
	MonthlySalary []Salary
}

func CreateEmployee() {

	employee1 := Employee{
		name:    "Yuvaraj",
		age:     21,
		address: "Chennai",
		MonthlySalary: []Salary{
			Salary{
				Basic: 12000,
				HRA:   2500,
				TA:    500,
			},
			Salary{
				Basic: 12500,
				HRA:   2000,
				TA:    500,
			},
		},
	}

	fmt.Println(employee1)
	
	for _, val := range employee1.MonthlySalary {
		fmt.Println(val.Basic)
		fmt.Println(val.HRA)
		fmt.Println(val.TA)
	}
}