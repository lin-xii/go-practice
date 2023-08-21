package control

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param age int整型 年龄
 * @return string字符串
 */
func getAge(age int) string {
	// write code here
	fmt.Println()
	if age < 1 {
		return "婴儿"
	} else if age < 5 {
		return "幼儿"
	} else if age < 12 {
		return "儿童"
	} else if age < 19 {
		return "少年"
	} else if age < 36 {
		return "青年"
	} else if age < 60 {
		return "中年"
	} else {
		return "老年"
	}
}

func ExecGetAge() {
	fmt.Println(getAge(1))
}

// 描述

// 已知年龄的分段如下， 婴儿(出生0-1岁)、幼儿(1-4岁)包含1岁、儿童(5-11)包含5岁、少年(12-18)包含12岁、青年(19-35)包含19岁、中年(36-59)包含36岁、老年(60以上)包含60岁 ，输入一个人的年龄，返回相应的年龄段。

// 知识点：

// if语句的语法为：
// if condition {
// }
// 如果 condition 为 true，那么就执行 { 和 } 之间的代码。

// if 语句后面可以接可选的 else if 和 else 语句：
// if condition {
// } else if condition {
// } else {
// }

// if 后面可以接任意数量的 else if 语句。condition 的求值由上到下依次进行，直到某个 if 或者 else if 中的 condition 为 true 时，执行相应的代码块。如果没有一个 conditon 为 true，则执行 else 中的代码块
