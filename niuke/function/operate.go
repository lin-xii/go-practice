package function

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param a int整型
 * @param b int整型
 * @return int整型一维数组
 */
func operate(a int, b int) []int {
	// write code here
	fmt.Println()
	result := []int{
		a + b,
		a - b,
		a * b,
		a / b,
		a % b,
	}
	return result
}

func ExecOperate() {
	a, b := 10, 2
	fmt.Println(operate(a, b))
}

// 描述

// 定义一个函数，实现输入a，b两个数，返回两数的和，差，乘积，商，余数。然后依次存入切片中，最后返回。

// 知识点：

// golang函数支持多个返回值，但函数有多个返回值时，如果其中某个或某几个返回值不想使用，可以通过下划线_来丢弃这些返回值。例如下面的f1函数两个返回值，调用该函数时，丢弃了第二个返回值b，只保留了第一个返回值a赋值给了变量a。
// 示例1

// 输入：
// 2,2
// 复制
// 返回值：
// [4,0,4,1,0]
