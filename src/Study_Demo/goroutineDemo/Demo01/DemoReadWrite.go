package main

import "fmt"

func main() {
	ch1 := make(chan int)
	//ch2 := make(chan<- int) //只写通道
	ch3 := make(<-chan int) //只读通道

	go WriteOnly(ch1)
	data2 := <-ch1
	fmt.Println("data2:", data2)

	//go WriteOnly2(ch2)

	go ReadOnly(ch1)
	ch1 <- 20

	go ReadOnly(ch3)
	fmt.Println("end")
}

func ReadOnly(ch <-chan int) {
	data := <-ch
	fmt.Println("读取到的通道数", data)
}

func WriteOnly(ch chan<- int) {
	ch <- 10
	fmt.Println("只写通道结束")
}
