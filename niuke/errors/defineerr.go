package errors

import "fmt"
import "errors"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param ping int整型 网络延迟值
 * @return string字符串
 */
func defineerr(ping int) string {
	// write code here
	if ping > 100 {
		return "网络延迟"
	}
	return ""
}

// type error interface {
// 	Error()
// }

type networkErr struct{}

func (e *networkErr) Error() error {
	return errors.New("网络延迟")
}

func ExecDefineerr() {
	params := 333
	fmt.Println(defineerr(params))
}

// 描述

// 实现erro接口，自定义一个错误，该错误抛出"网络延迟"错误。输入网络的延迟数，如果延迟数大于100则认为网络延迟，并返回

// 知识点：
//         错误：Go 语言通过内置的错误接口提供了非常简单的错误处理机制。
//         type error interface {
//         Error() string
//         }

//        实现 error 接口类型来生成错误信息，函数通常在最后的返回值中返回错误信息。使用errors.New 可返回一个错误信息
// 示例1

// 输入：
// 150
// 复制
// 返回值：
// "网络延迟"
