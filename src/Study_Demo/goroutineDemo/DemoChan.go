package main

import "fmt"

func main() {
	ch1 := make(chan int)

	go func() {
		fmt.Println("子goroutine start")
		for i := 0; i < 10; i++ {
			ch1 <- i //往通道中放数据
		}
		close(ch1)
	}()
	//for {
	//	v, ok := <-ch1
	//	if !ok {
	//		fmt.Println("子协程通道已关闭")
	//		break
	//	}
	//	fmt.Println(v)
	//}
	//fmt.Println("主携程结束")

	//for range遍历所有数据，当没有数据的时候也就循环结束，可以替代 v,ok:=<-ch1
	for v := range ch1 {
		fmt.Println("读取到的通道的数据", v)
	}
	fmt.Println("主协程结束")
}
