package niuke

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param a int整型
 * @param b int整型
 * @return int整型一维数组
 */
func bitOperate(a int, b int) []int {
	// write code here
	result := []int{
		a & b,
		a | b,
		a ^ b,
	}

	return result
}

func ExecBitOperation() {
	fmt.Println(bitOperate(8, 3))
}

// 描述

// 已知a,b两个int类型变量，求出这两个变量的与，或，异或值，将结果依次存入切片中，然后返回这个切片。

// 知识点：

// golang 中 位运算符&  按位与 是双目运算符。 其功能是参与运算的两数各对应的二进位相与。

// golang 中 位运算符&  按位或 是双目运算符。 其功能是参与运算的两数各对应的二进位相或。

// golang 中 位运算符&  按位异或 是双目运算符。 其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。
