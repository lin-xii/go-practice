package loop

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param array int整型二维数组
 * @return int整型一维数组
 */
func transform(array [][]int) []int {
	// write code here
	result := []int{}
	m := len(array)
	for ix, x := range array {
		n := len(x)
		fmt.Println("n:", n)
		for yx := range x {
			result = append(result, ix*n+yx)
			// fmt.Println(y)
		}
	}
	fmt.Println("m:", m)
	return result
}

func ExecTransform() {
	// [3][2]int{}
	// params := [][]int{
	// 	{11, 12},
	// 	{21, 22},
	// 	{31, 32},
	// }
	params := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	// fmt.Println(params)
	fmt.Println(transform(params))
}

// 输入：
// [[1,2,3],[4,5,6],[7,8,9]]
// 复制
// 返回值：
// [1,2,3,4,5,6,7,8,9]

// 描述

// 已知一个m*n二维数组，二维数组中的元素的索引（x，y）可以表示为一个二维坐标,现将这个二维坐标转换为一维坐标,一维坐标=x*n+y。返回这个一维数组。

// 知识点：

// 数组：是同一种数据类型的固定长度的序列。

// 数组定义：var a [len]int，比如：var a [5]int，数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变。

// 长度是数组类型的一部分，因此，var a[5] int和var a[10]int是不同的类型。

// 数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1
//     for i := 0; i < len(a); i++ {
//     }
//     for index, v := range a {
//     }

// 访问越界，如果下标在数组合法范围之外，则触发访问越界，会panic

// 输入描述：

// 输入任意一种类型的变量
// 返回值描述：

// 返回该具体类型
