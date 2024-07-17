package main

import "fmt"

//通过接口实现oop中的多态
//多态是指一种事务的多种形态，比如大象既是其本身，也是动物

type Animal interface {
	eat()
	sleep()
}

type Dog struct {
	name string
}

func (dog Dog) eat() {
	fmt.Println(dog.name, "dog eat")
}
func (dog Dog) sleep() {
	fmt.Println(dog.name, "dog sleep")
}

func main() {

	//多态的具体体现，dog既是Dog，也是Animal
	//运行结果
	//小黑球 dog eat
	//小黑球 dog sleep
	//小黑球 dog eat
	//小黑球 dog sleep

	dog1 := Dog{name: "小黑球"}
	dog1.eat()
	dog1.sleep()

	var animal Animal
	animal = dog1
	test2(animal)

}

func test2(animal Animal) {
	animal.eat()
	animal.sleep()
}
