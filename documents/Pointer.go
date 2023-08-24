package main

import "fmt"

type persion struct {
	name string
	age  int
}

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

	fmt.Println("--------------------")
	person1 := persion{"name1", 21}
	fmt.Println("person1", person1)

	person2 := &person1
	fmt.Println("person2:", person2)
	fmt.Println(*&person2.name)
	fmt.Println((*person2).name)
	fmt.Println(*&person2.age)
}
