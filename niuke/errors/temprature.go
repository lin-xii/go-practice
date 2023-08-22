package errors

import (
	"errors"
	"fmt"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param t double浮点型 体温
 * @return string字符串
 */
func temperature(t float64) string {
	// write code here
	fmt.Println()
	if t > 37.5 {
		errors.New("体温异常")
		return "体温异常"
	}
	return ""
}

func ExecTemprature() {
	params := 38.0
	fmt.Println(temperature(params))
}

// 描述

// 实现一个函数，该函数的功能是 给定一个float类型变量表示某个人的体温，如果有人体温大于37.5，抛出"体温异常"，并输出。

// 知识点：

// Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。

// 示例1

// 输入：
// 38.000000
// 复制
// 返回值：
// "体温异常"
