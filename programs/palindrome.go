package programs

import "strings"

func IsPalindrome(str string) bool {
	cleanedStr := strings.ToLower(strings.ReplaceAll(str, " ", ""))

	length := len(cleanedStr)

	for i := 0; i < length/2; i++ {
		if cleanedStr[i] != cleanedStr[length-i-1] {
			return false
		}
	}

	return true
}
