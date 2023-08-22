package structs

import "fmt"

func Students2() {
	type student struct {
		name  string
		sex   bool
		age   uint
		score uint
	}

	xiaoming := student{
		name:  "小明",
		sex:   true,
		age:   23,
		score: 88,
	}
	// fmt.Println(xiaoming)
	fmt.Println(xiaoming.name)
	fmt.Println(xiaoming.sex)
	fmt.Println(xiaoming.age)
	fmt.Println(xiaoming.score)
}

// 描述

// 每个学生都有如下信息：姓名name，性别sex，年龄age，分数score，定义一个结构体Student，来表示小明的信息，小明的信息如下：姓名：小明，性别：true，年龄：23，分数：88
// 输出并打印小明信息
// 知识点：

// Go语言中的基础数据类型可以表示一些事物的基本属性，但是当我们想表达一个事物的全部或部分属性时，这时候再用单一的基本数据类型明显就无法满足需求了，Go语言提供了一种自定义数据类型，可以封装多个基本数据类型，这种数据类型叫结构体，英文名称struct。 也就是我们可以通过struct来定义自己的类型了。

// Go语言中通过struct来实现面向对象。

// 结构体由一系列命名的元素组成，这些元素又被称为字段，每个字段都有一个名称和一个类型。

// 结构体的目的就是把数据聚集在一起，以便能够更加便捷地操作这些数据。

// 结构体的定义：结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体中有一个或多个成员。type 语句设定了结构体的名称。结构体的格式如下：
// type 类型名 struct {
//         字段名 字段类型
//         字段名 字段类型
//         …
//     }
//      类型名：标识自定义结构体的名称，在同一个包内不能重复。
//      字段名：表示结构体字段名。结构体中的字段名必须唯一。
//      字段类型：表示结构体字段的具体类型。

// 访问结构体的成员：如果要访问结构体成员，需要使用点号 . 操作符

// 输入描述：

// 无
// 输出描述：

// 小明
// true
// 23
// 88

func Students3() {
	type address struct {
		city     string
		province string
	}
	type student struct {
		name  string
		sex   bool
		age   uint
		score uint
		address
	}
	xiaoming := student{
		name:  "小明",
		sex:   true,
		age:   23,
		score: 88,
		address: address{
			city:     "长沙市",
			province: "湖南省",
		},
	}
	fmt.Println(xiaoming.name)
	fmt.Println(xiaoming.sex)
	fmt.Println(xiaoming.age)
	fmt.Println(xiaoming.score)
	fmt.Println(xiaoming.province)
	fmt.Println(xiaoming.city)
}

// 描述

// 每个学生都有如下信息：姓名name，性别sex，年龄age，分数score，地址信息address,其中address地址信息又包含城市city,省份province等信息， 定义一个结构体Student，来表示小明的信息，小明的信息如下：姓名：小明，性别：true，年龄：23，分数：88，province：湖南省，city：长沙市。
// 依次输出打印小明信息

// 知识点：

// Go语言中的基础数据类型可以表示一些事物的基本属性，但是当我们想表达一个事物的全部或部分属性时，这时候再用单一的基本数据类型明显就无法满足需求了，Go语言提供了一种自定义数据类型，可以封装多个基本数据类型，这种数据类型叫结构体，英文名称struct。 也就是我们可以通过struct来定义自己的类型了。

// Go语言中通过struct来实现面向对象。

// 结构体由一系列命名的元素组成，这些元素又被称为字段，每个字段都有一个名称和一个类型。

// 结构体的目的就是把数据聚集在一起，以便能够更加便捷地操作这些数据。

// 结构体的定义：结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体中有一个或多个成员。type 语句设定了结构体的名称。结构体的格式如下：
// type 类型名 struct {
//         字段名 字段类型
//         字段名 字段类型
//         …
//     }
//       类型名：标识自定义结构体的名称，在同一个包内不能重复。
//       字段名：表示结构体字段名。结构体中的字段名必须唯一。
//       字段类型：表示结构体字段的具体类型。

// 访问结构体的成员：如果要访问结构体成员，需要使用点号 . 操作符

// 一个结构体中可以嵌套包含另一个结构体或结构体指针。

// 一个结构体中可以嵌套包含另一个结构体或结构体指针。
// 输入描述：

// 无
// 输出描述：

// 小明
// true
// 23
// 88
// 湖南省
// 长沙市
