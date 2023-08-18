package niuke

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param x int整型 参加活动人数
 * @return bool布尔型
 */
func odevity(x int) bool {
	// write code here
	// result := x % 2
	// 奇数：false
	// 1
	// 偶数：true
	// 0
	return x%2 == 0
}

func ExecOdevity() {
	fmt.Println(odevity(9))
}

// 描述

// 某公司举办了一个联谊活动，现在要统计参加活动人数的单双，如果是单数，返回false，偶数返回true

// 知识点：

// golang中，%为取余操作，奇偶性的判断可以通过取余来判断，如果对2取余，余数为0则为偶数，如果为1，则为奇数。
