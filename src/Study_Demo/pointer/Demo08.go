package main

import "fmt"

func setName(name *string) {
	*name = "郭峰"
	fmt.Println("函数内部", *name, name)
}

func main() {
	name := "张丹"
	fmt.Println("函数外部", &name, name)
	setName(&name)
	fmt.Println(name)
}
