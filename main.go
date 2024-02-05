package main

import "fmt"

// 一些练习
func main() {
	fmt.Println(addition(1, 2))
}

func addition(a, b int) int {
	return a + b
}

func primitiveType() {
	var aInt8 int8 = 128
	var aUint8 uint8 = 256
	var aFloat float32 = 1.1
	var aString string = "hello"
	var aBool bool = true
}
