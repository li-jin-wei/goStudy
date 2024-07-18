package main

import (
	"fmt"
	"go-project/src/Study_Demo/structs/oop/Demo01"
)

func main() {
	var p1 = Demo01.NewPerson("张三", 20, "女")
	fmt.Println(*p1)
	fmt.Println(p1.GetName())
	p1.SetAge(-1)
	fmt.Println(p1.GetAge())
}
