package main

import (
	"fmt"
	"sync"
)

//func3......
//func1......
//func2......
//fatal error: all goroutines are asleep - deadlock!
//使用同步等待组创建协程，主协程中设置同步等待组的数量为4，而只加进去了3条协程，
//最终都执行完成之后，还有一条未执行，当程序进入阻塞状态的时候无法解锁，就造成了死锁。

var wg sync.WaitGroup

//创建同步等待组

func main() {
	wg.Add(4)
	go Sale1()
	go Sale2()
	go Sale3()
	wg.Wait()
	fmt.Println("main end")
}

func Sale1() {
	fmt.Println("func1......")
	wg.Done()
}
func Sale2() {
	defer wg.Done()
	fmt.Println("func2......")
}
func Sale3() {
	defer wg.Done()
	fmt.Println("func3......")
}
