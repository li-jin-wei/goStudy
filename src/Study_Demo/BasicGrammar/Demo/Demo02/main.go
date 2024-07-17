package main

import "fmt"

var a int = 123

func main() {
	//	就近原则
	fmt.Println(a, &a)
	a := 456
	fmt.Println(a, &a)

}
