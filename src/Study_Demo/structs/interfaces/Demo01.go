package main

import "fmt"

//接口

// 接口，方法的集合

type USB interface {
	input()
	output()
}

type Mouse struct {
	name string
}

// 结构体实现了接口的所有方法才算真正实现了接口
func (mouse Mouse) output() {
	fmt.Println(mouse.name, "鼠标输出")
}

func (mouse Mouse) input() {
	fmt.Println(mouse.name, "鼠标输入")
}

type KeyBoard struct {
	name string
}

func (key KeyBoard) input() {
	fmt.Println(key.name, "键盘输出")
}
func (key KeyBoard) output() {
	fmt.Println(key.name, "键盘输入")
}

// 函数参数需要时USB接口类型或者他的实现结构体，
// 如果Mouse或者KeyBoard 没有全部实现USB接口的方法，使用此方法传入其实例化对象就会报错
// cannot use m1 (variable of type Mouse) as USB value in argument to test: Mouse does not implement USB (missing method input)
func test(u USB) {
	u.input()
	u.output()
}
func main() {
	m1 := Mouse{name: "罗技"}
	test(m1)

	k1 := KeyBoard{name: "雷蛇"}
	test(k1)

}
