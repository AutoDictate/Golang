package conditionalstatements

import (
	"fmt"
	"time"
)


func SwitchCase() {

	today := time.Now()

	fmt.Println("Day : ",today.Day())
	fmt.Println("Month : ",today.Month())
	fmt.Println("Year : ",today.Year())

	switch today.Day() {
		
	case 5, 7, 8:
		fmt.Println("Today is 5th. Clean your house.")
		fallthrough
	case 1:
		fmt.Println("Today is 1st. Join the Company.")
		fallthrough
	case 10:
		fmt.Println("Today is 10th. Buy some wine.")
		fallthrough
	case 15:
		fmt.Println("Today is 15th. Visit a doctor.")
		fallthrough
	case 25:
		fmt.Println("Today is 25th. Buy some food.")
		fallthrough
	case 31:
		fmt.Println("Party tonight.")
		fallthrough
	default:
		fmt.Println("No information available for that day.")
	}

}

