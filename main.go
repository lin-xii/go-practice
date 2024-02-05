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
	var aInt int = 2 ^ 32 + 1
	// 0 ~ 127
	var aInt8 int8 = 127
	// -256 ~ 255
	var aUint8 uint8 = 255
	var aFloat float32 = 1.1
	var aString string = "hello"
	var aBool bool = true

	fmt.Println(aInt, aInt8, aUint8, aFloat, aString, aBool)
}
