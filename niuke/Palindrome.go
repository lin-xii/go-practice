package niuke

import (
	"fmt"
	"strconv"
)

func IsPalindrome(x int) bool {
	// write code here
	var result bool

	xstr := strconv.Itoa(x)
	var temp string
	for i := len(xstr) - 1; i >= 0; i-- {
		temp += string(xstr[i])
	}
	result = xstr == temp

	return result
}

func ExecIsPalindrome() {
	fmt.Println(IsPalindrome(12321))
}
