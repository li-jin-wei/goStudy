package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	s1 := Student{"小红", 19}
	fmt.Println(s1)

	//匿名结构体，直接可以在函数内部创建出来，创建后就需要赋值使用
	s2 := struct {
		name string
		age  int
	}{
		name: "张三",
		age:  3,
	}
	fmt.Println(s2.name)
	fmt.Println(s1.age)

	//匿名字段，使用默认数据类型当做字段名称
	t1 := Teacher{"宋", 18}
	fmt.Println(t1)
	fmt.Println(t1.string)
	fmt.Println(t1.int)
}

// Teacher 结构体中的匿名字段，没有名字的字段，这个时候属性类型不能重复
type Teacher struct {
	string
	int
}
