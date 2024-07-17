package main

import "fmt"

// Person 创建结构体Person
type Person struct {
	name    string
	age     int
	sex     string
	address string
}

func main() {

	//实例化Person结构体
	p := Person{}

	//给实例化对象赋值
	p.name = "张三"
	p.age = 10
	p.sex = "男"
	p.address = "陕西省"
	//fmt.Println(&p)
	fmt.Println(p)
	fmt.Println(p.name, p.age, p.address, p.sex)

	//内置函数 new 创建对象。  new 关键字创建的对象，都返回指针，而不是结构体对象。
	p1 := new(Person)
	p1.age = 22
	fmt.Println(*p1)
	fmt.Println(p1.age)

}
