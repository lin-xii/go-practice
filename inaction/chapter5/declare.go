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

func ComplexTypeDeclare() {
	var aArray [3]int
	fmt.Println(aArray, aArray == [3]int{0, 0, 0}, cap(aArray), len(aArray))
	// fmt.Println(aArray, aArray == [3]int{0, 0, 0}, aArray == nil, cap(aArray), len(aArray))

	var aSlice []int
	fmt.Println(aSlice, aSlice == nil, cap(aSlice), len(aSlice))

	var aMap map[string]int
	fmt.Println(aMap, aMap == nil, len(aMap))

	var aPerson Person
	fmt.Println(aPerson)
	// fmt.Println(aPerson, aPerson == nil)

	var aDuration Duration
	fmt.Println(aDuration, aDuration == 0)
	aDuration = Duration(1)
}

type Person struct {
	name   string
	age    uint8
	sex    bool
	height float32
}

type Duration int64
