package interfaces

import (
	"fmt"
)

type animals interface {
	sleep()
	eat()
}

type tiger struct{}
type chicken struct{}

func (t *tiger) sleep() {
	fmt.Println("sleep")
}

func (t *tiger) eat() {
	fmt.Println("eat")
}

func (c *chicken) eat() {
	fmt.Println("chicken eat")
}

func Animals() {
	// t1 := tiger{}
	var t1 animals = &tiger{}
	// var c1 animals = &chicken{}
	t1.sleep()
	t1.eat()
}

// 描述

// 定义一个动物接口，该接口有 sleep,eat 方法，定义老虎实现动物接口，老虎的sleep方法输出"tiger sleep",eat方法输出"tiger eat"，最后依次调用老虎的sleep,eat方法。

// 知识点：

//  在Go语言中接口（interface）是一种类型，一种抽象的类型。interface是一组method的集合，是duck-type programming的一种体现。接口做的事情就像是定义一个协议（规则）， 只要一台机器有洗衣服和甩干的功能，我就称它为洗衣机。不关心属性（数据），只关心行为（方法）。

// 接口（interface）定义了一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节。面向接口编程

// 接口是一个或多个方法签名的集合。

// 任何类型的方法集中只要拥有该接口'对应的全部方法'签名。就表示它 "实现" 了该接口，无须在该类型上显式声明实现了哪个接口。这称为Structural Typing。所谓对应方法，是指有相同名称、参数列表 (不包括参数名) 以及返回值。当然，该类型还可以有其他方法。

// 接口只有方法声明，没有实现，没有数据字段。

// 接口可以匿名嵌入其他接口，或嵌入到结构中。

// 对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，既无法修改复制品的状态，也无法获取指针。

// 只有当接口存储的类型和对象都为nil时，接口才等于nil。

// 接口调用不会做receiver的自动转换。

// 接口同样支持匿名字段方法。

// 接口也可实现类似OOP中的多态。

// 空接口可以作为任何类型数据的容器。

// 一个类型可实现多个接口。

// 接口命名习惯以 er 结尾。

// type 接口类型名 interface{
//         方法名1( 参数列表1 ) 返回值列表1
//         方法名2( 参数列表2 ) 返回值列表2
//         …
//     }
// 接口名：使用type将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，有字符串功能的接口叫Stringer等。接口名最好要能突出该接口的类型含义。
// 方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
// 参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。
// 输入描述：

// 无
// 输出描述：

// sleep
// eat
