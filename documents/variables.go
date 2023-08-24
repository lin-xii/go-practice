package main

import "fmt"

func Variables() {
	var a = "initial"
	fmt.Println(a)

	// go里面，会做自动类型转换
	var b, c int = 1, 2.0
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// zero value. go：仅初始化的变量，是有零值的
	// 基础类型的zero value：0、“”、false
	// 堆类型的zero value：==nil
	var e int
	var f float32
	var g bool
	var h string
	var i map[int]string
	var k []string
	fmt.Println(e, f, g, h, i, k)
	// fmt.Println(h == nil, i == nil)
	fmt.Println(i == nil, k == nil)

	l := "apple"
	m := 'm'
	fmt.Println(l, m)

	// 强类型，因为没有过度的隐式类型转换
	// 例如js的 1+‘string’。。。
	// fmt.Println(e + m)
}
