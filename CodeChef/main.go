package main

import (
	"fmt"
)

func main() {

	res := SingleNumber([]int{1, 1, 2, 3, 3, 4})
	fmt.Println(res)
}

func SquareRoot(x int) int {

	if x < 2 {
		return x
	}

	left, right := 1, x

	for left <= right {

		mid := left + (right-left)/2

		square := mid * mid

		if square == x {
			return mid
		} else if square < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}

func SingleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	hMap := make(map[int]int)

	fmt.Println("Length of the map : ", len(hMap))
	fmt.Println(hMap[1])

	for _, val := range nums {
		hMap[val]++
	}

	for key, val := range hMap {
		if val == 1 {
			return key
		}
	}

	return -1
}
