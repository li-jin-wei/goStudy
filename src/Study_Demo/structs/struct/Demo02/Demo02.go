package main

import "fmt"

type Person01 struct {
	name string
	age  int
	sex  string
}

func newPerson(name string, age int, sex string) *Person01 {
	return &Person01{
		name: name,
		age:  age,
		sex:  sex,
	}
}
func SetName(name string) {
	fmt.Println(name)
}

func GetName() (name string) {
	return name
}

func main() {
	p := newPerson("张三", 30, "男")
	//fmt.Println(p)
	fmt.Println(*p)

}
