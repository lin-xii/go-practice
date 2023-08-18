package niuke

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s int整型一维数组 身高
 * @param index int整型 出队索引
 * @return int整型一维数组
 */
func deleteElement(s []int, index int) []int {
	// write code here
	result := make([]int, 0)
	for i, v := range s {
		if i == index {
			continue
		}
		result = append(result, v)
	}
	return result
}

func ExecDeleteElement() {
	temp := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(deleteElement(temp, 2))
}

// 描述

// 学生们都排成了一队，有一个切片表示相应学生们的身高，现随机喊某个位置的人出队，返回出队后的这个切片。比如[2,3,4,5],索引为1的位置的出队，出队后切片为[2,4,5]

// 知识点：
// s|n|    切片s中索引位置为n的项
// s|:|    从切片s的索引位置0到len(s)-1 处所获得的切片
// s|low:|   从切片s的索引位置 low 到len(s)-1 处所获得的切片
// s|:high| 从切片s的索引位置 0到high 处所获得的切片，len=high
// s|low: high|  从切片s的素引位置 Iow 到high 处所获得的切片，len-high-low
// s|low: high:max|  从切片s的素引位置 low 到high 处所获得的切片，len-high-low, cap=max-low
