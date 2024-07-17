package main

import "fmt"

// 继承02
type Person struct {
	name string
	age  int
}

type Student struct {
	Person //匿名字段，实现了继承
	school string
}

func main() {
	//创建父类对象
	p1 := Person{name: "小华", age: 20}
	fmt.Println(p1)
	fmt.Println(p1.age, p1.name)

	//创建子类
	s1 := Student{Person: Person{name: "小明", age: 25}, school: "清华"}
	fmt.Println(s1)
	fmt.Println(s1.Person.age, s1.Person.name, s1.school)

	// 概念：提升字段, 只有匿名字段才可以做到
	// Person 在Student中是一个匿名字段， Person中的属性 name age 就是提升字段
	// 所有的提升字段就可以直接使用了，不同在通过匿名字段来点了
	var s2 Student
	s2.name = "小花"
	s2.age = 20
	s2.school = "浙大"
	fmt.Println(s2)
	fmt.Println(s2.age, s2.name, s2.school)
}
