package niuke

import (
	"fmt"
	"unicode/utf8"
)

func CountString(s string) int {
	// var length int

	length := utf8.RuneCountInString(s)
	// length = len(s)

	return length
}

func ExecCountString() {
	length := CountString("中国pro")
	fmt.Println(length)
}
