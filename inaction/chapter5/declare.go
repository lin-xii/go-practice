package chapter5

import "fmt"

func BasicTypeDeclare() {
	// 基本类型的zero value
	var aInt int
	fmt.Println(aInt, aInt == 0)

	var aString string
	fmt.Println(aString, aString == "")

	var aBool bool
	fmt.Println(aBool, aBool == false)

	var aFloat float64
	fmt.Println(aFloat, aFloat == 0.0)
}
