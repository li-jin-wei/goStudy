package main

import "fmt"

type Person01 struct {
	name string
	age  int
	sex  string
}

// 初始化方法，返回一个指针类型
func newPerson(name string, age int, sex string) *Person01 {
	return &Person01{
		name: name,
		age:  age,
		sex:  sex,
	}
}

// SetName 对属性的封装，修改属性
func SetName(name string) {
	fmt.Println(name)
}

// GetName 对属性的封装，获取属性
func GetName() (name string) {
	return name
}

func main() {
	p := newPerson("张三", 30, "男")
	//fmt.Println(p)
	fmt.Println(*p)

}
