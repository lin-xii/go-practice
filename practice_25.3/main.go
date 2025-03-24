package practice253

import "fmt"

// print后的会多一个%，不是print的问题。是terminal追加的输入符号
func VerifyPrint() {
	fmt.Print("hello go.")
	fmt.Println("hello go.")
	fmt.Print("hello go without newline: ")
	fmt.Println("hello go with newline")
	fmt.Print("another line without newline")
	fmt.Print(" still on the same line")
	fmt.Println() // 只打印一个换行符
	fmt.Println("new line after newline")
}
