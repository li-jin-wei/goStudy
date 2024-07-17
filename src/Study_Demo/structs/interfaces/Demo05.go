package main

import (
	"fmt"
)

//接口断言

type I interface{}

func testAsserts(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("变量为string类型")
	case int:
		fmt.Println("变量为int类型")
	case I:
		fmt.Println("变量为I类型")
	case nil:
		fmt.Println("变量为nil类型")
	case interface{}:
		fmt.Println("变量为interface{}类型")

	default:
		fmt.Println("未知类型")

	}
}

func main() {
	testAsserts(false)
	testAsserts("string")
	testAsserts(1)

	var i I
	var i2 I = 10
	testAsserts(i)
	testAsserts(i2)
}
