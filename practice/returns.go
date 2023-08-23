package practice

import "fmt"

// 命名返回值时，不需要在函数体中声明变量
func multiReturns() (first int, second float32) {
	defer func() {
		// 命名变量，在defer里可以修改，或许是作用域的问题吧？
		// 之前tonybai的课里也讲过，go也有shadow问题。
		second = 3.0
	}()
	first = 1
	second = 2.0
	return
}

// 如果返回值是值类型，defer修改是不生效的，得是pointer才行
// func multiReturns() string {
// 	single := "single"
// 	defer func() {
// 		single = "defer"
// 		// return single
// 	}()
// 	return single
// }

// pointer类型，defer修改值是生效的
// func multiReturns() *string {
// 	single := "single"
// 	defer func() {
// 		single = "defer"
// 		// return single
// 	}()
// 	return &single
// }

func RunMultiReturns() {
	fmt.Println(multiReturns())
}
