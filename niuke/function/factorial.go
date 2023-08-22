package function

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param i int整型 数字
 * @return int整型
 */
func factorial(i int) int {
	// write code here
	fmt.Print()
	if i > 0 {
		return i * factorial(i-1)
	}
	return 1
}

func ExecFactorial() {
	params := 4
	fmt.Println(factorial(params))
}

// 描述

// 一个正整数的阶乘（factorial）是所有小于及等于该数的正整数的积，并且0的阶乘为1。自然数n的阶乘写作n!。

// 知识点：

//       递归，就是在运行的过程中调用自己。 一个函数调用自己，就叫做递归函数。构成递归需具备的条件：
//       子问题须与原始问题为同样的事，且更为简单。
//       不能无限制地调用本身，须有个出口，化简为非递归状况处理。
// 示例1

// 输入：
// 1
// 复制
// 返回值：
// 1
// 复制
// 示例2

// 输入：
// 2
// 复制
// 返回值：
// 2
