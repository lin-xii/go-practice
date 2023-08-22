package loop

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param score int整型一维数组 团队成员分数
 * @param target int整型 达标分数
 * @return bool布尔型
 */
func canPass(score []int, target int) bool {
	// write code here
	for _, v := range score {
		if v > target {
			return true
		}
	}
	return false
}

func ExecCanPass() {
	score := []int{1, 2, 3, 4, 5, 7, 8}
	target := 3
	fmt.Println(canPass(score, target))
}

// 描述

// 有一个团队闯关游戏，通过规则是团队中的某一个人的分数大于标准分数，这个团队就算通关。给出该团队所有人的分数，判断该团队是否能通关，能通关返回true,不能返回false。

// 知识点：

// golang中break 可以中断当前 for 循环或跳出 switch 语句。
// 示例1

// 输入：
// [1,2,3,4,8],7
// 复制
// 返回值：
// true
