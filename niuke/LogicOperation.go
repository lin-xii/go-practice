package niuke

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param a bool布尔型
 * @param b bool布尔型
 * @return bool布尔型一维数组
 */
func logicalOperation(a bool, b bool) []bool {
	// write code here
	result := []bool{
		a && b,
		a || b,
		!a,
		!b,
	}
	return result
}

func ExecLogicOperation() {
	fmt.Println(logicalOperation(true, true))
}

// 描述

// 给定两个bool类型变量a,b，求出这两个bool类型变量的逻辑and，or，not a,not b的值，将依次存入一个切片中，然后返回这个切片。

// 知识点：

// golang中，&& 表示逻辑 AND 运算符。 如果两边的操作数都是 True，则为 True，否则为 False。

// golang中，｜｜ 表示逻辑 OR 运算符。 如果两边的操作数有一个 True，则为 True，否则为 False。

// golang中，！ 表示逻辑 NOT 运算符。 如果条件为 True，则为 False，否则为 True。
