package datatypeconversion

import (
	"fmt"
	"strconv"
)

func IntLimitCheck() {

	var a int = 2399999999999999999

	b := strconv.Itoa(a);
	fmt.Println(len(b));

	for i := 1; i <= len(b); i++ {
		fmt.Println(i)
	}
}