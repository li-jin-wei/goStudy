package main

import "fmt"

//接口的嵌套

type i1 interface {
	Test1()
}

type i2 interface {
	Test2()
}

type i3 interface {
	i1
	i2
	Test3()
}

// Person 结构体
type Person struct {
}

// 实现接口i3

func (p Person) Test1() {
	fmt.Println("I'm test1")
}
func (p Person) Test2() {
	fmt.Println("I'm test2")
}
func (p Person) Test3() {
	fmt.Println("I'm test3")
}

func main() {
	//person 四种形态 Person test1 test2 test3
	var person Person = Person{}
	person.Test1()
	person.Test2()
	person.Test3()

	var t i1 = person //向上转型
	t.Test1()

	//t.Test2()  //向上转型只能调用它自己对应的方法

	var t2 i2 = person
	t2.Test2()
}
