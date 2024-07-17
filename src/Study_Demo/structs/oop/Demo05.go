package main

import "fmt"

//方法的重写

type Animal struct {
	name string
	age  int
}

func (a Animal) eat() {
	fmt.Println(a.name, "正在吃.....")
}

func (a Animal) sleep() {
	fmt.Println(a.name, "正在sleep.....")
}

type Dog struct {
	Animal
}

func (d Dog) wang() {
	fmt.Println("wangwangwanwg~~~")
}

type Cat struct {
	Animal
	color string
}

// 重写父类的方法
func (c Cat) eat() {
	fmt.Println(c.name, "正在吃....")
}

func main() {
	dog := Dog{Animal{name: "旺财", age: 20}}
	dog.eat()
	dog.wang()
	cat := Cat{Animal{name: "煤球", age: 3}, "黑色"}
	cat.eat()
	fmt.Println(cat.color)

	a := Animal{name: "大黄", age: 20}
	a.eat()
	a.sleep()
}
