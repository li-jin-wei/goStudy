package main

import "fmt"

//方法
//可以理解为函数多了一个调用者
//方法可以重名，不同的对象，调用的结果是不一样的

type Dog struct {
	name string
	age  int
}

func (d Dog) eat() {
	fmt.Println("Dog eating...")
}

func (dog Dog) sleep() {
	fmt.Println("Dog sleeping...")
}

type Cat struct {
	name string
	age  int
}

func (cat Cat) eat() {
	fmt.Println(cat.name, "eating...")
}

func (cat Cat) sleep() {
	fmt.Println(cat.name, "sleeping...")
}

func main() {
	dog := Dog{
		name: "旺财",
		age:  20,
	}
	dog.eat()

	cat := Cat{name: "喵喵", age: 1}
	cat.eat()
}
