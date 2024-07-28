package main

import "fmt"

func main() {
	go test()
	for a := 0; a < 5; a++ {
		fmt.Println(a)
	}
	fmt.Println("main 函数结束")
}

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("+测试goroutine", i)
	}
}
