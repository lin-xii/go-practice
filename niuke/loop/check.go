package loop

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param material int整型一维数组 成品质量
 * @param standard int整型 质检标准
 * @return int整型一维数组
 */
func check(material []int, standard int) []int {
	// write code here
	result := make([]int, 0)
	for _, v := range material {
		if v < standard {
			continue
		}
		result = append(result, v)
	}
	return result
}

func ExecCheck() {
	quality := []int{1, 3, 4, 5, 6, 7, 8, 89}
	standard := 5
	fmt.Println(check(quality, standard))
}

// 描述

// 某工厂加工了一批材料，现将加工出来的成品抽出来一部分进行检查，给定一个标准，如果低于这个标准，则视为不合格。所有成品的质量用一个切片来表示。

// 知识点：

// golang continue跳出当前循环，继续下一次循环

// 示例1

// 输入：
// [2,2,3,4,6,6,3],3
// 复制
// 返回值：
// [3,4,6,6,3]
