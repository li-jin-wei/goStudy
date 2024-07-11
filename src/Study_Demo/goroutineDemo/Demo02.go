package main

import (
	"fmt"
	"time"
)

func main() {
	go testgo1()
	go testgo2()
	for i := 0; i <= 5; i++ {
		fmt.Println("main函数执行", i)
	}
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main 函数结束")
}

func testgo1() {
	for i := 0; i < 10; i++ {
		fmt.Println("测试子goroutine1", i)
	}
}

func testgo2() {
	for i := 0; i < 10; i++ {
		fmt.Println("测试子goroutine2", i)
	}
}
