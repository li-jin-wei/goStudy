package main

import (
	"fmt"
	//"go-project/src/Demo/testPackage"
	"strconv"
)

var string1 string = "123456"

var u1 uint = 100

func main() {
	//fmt.Println(testPackage.A)
	//fmt.Println(testPackage.B)
	i1, _ := strconv.Atoi(string1)
	fmt.Println(i1)
	fmt.Printf("%T\n", i1)

	fmt.Printf("%T", string1)
	fmt.Println()
	fmt.Printf("%T", u1)
}
