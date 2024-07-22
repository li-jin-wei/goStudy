package main

import "fmt"

func main() {

	var i = 5
	var intP *int //声明一个指向int的指针
	intP = &i     //将intP指向i的地址
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
}
