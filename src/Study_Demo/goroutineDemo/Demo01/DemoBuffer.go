package main

import "fmt"

func main() {
	ch1 := make(chan int, 5)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	ch1 <- 4
	ch1 <- 5 //此时缓冲区已经满 如果再加入 则会进入阻塞状态

	//继续添加时会造成死锁 因为缓冲区满了 一直没有读取
	ch1 <- 6
	fmt.Println("main end")
}
