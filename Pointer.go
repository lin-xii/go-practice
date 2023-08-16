package main

import "fmt"

func Pointer() {
	var a int = 1
	fmt.Println(a, &a, *(&a))
	fmt.Println("--------------------")

	b := &a
	// 可以看出来，b和&a，指向的都是a的内存地址
	fmt.Println(b)
	fmt.Println(&b)
	// 其实，就是*了两次，才拿到的值
	fmt.Println(**(&b))
}
