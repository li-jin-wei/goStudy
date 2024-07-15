package main

import "fmt"

//*是获取指针所指向的地址的值

func main() {
	var a int = 10

	var b = &a
	fmt.Println("a变量的地址:", &a)
	fmt.Println("b存储的地址:", b)
	fmt.Println("b的地址", &b)
	fmt.Println("b变量指向的地址中的值", *b)

	*b = 20
	//通过指针变量b,操控了a的值
	fmt.Println(a)
}
