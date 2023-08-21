package maps

import "fmt"

func score() {
	scores := map[string]int{"小明": 60, "小王": 70, "张三": 95, "李四": 98, "王五": 100, "张伟": 88}
	fmt.Println(scores)
}

func ExecScore() {
	score()
}

// 描述

// 某大学宿舍6人的数学成绩分别为 小明：60，小王：70，张三：95，李四：98，王五：100，张伟：88 ，现要将六人的成绩录入成绩表中，这个成绩表用一个map来表示，成绩表的键为宿舍成员的姓名，值为对应的分数 。打印该成绩表

// 知识点：
// map的声明：map[KeyType]ValueType KeyType:表示键的类型。ValueType:表示键对应的值的类型。map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：make(map[KeyType]ValueType, [cap]) 其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。

// map[key]=value 给指定的key赋指定的value

// 输入描述：

// 无
// 输出描述：

// map[小明:60 小王:70 张三:95 张伟:88 李四:98 王五:100]
