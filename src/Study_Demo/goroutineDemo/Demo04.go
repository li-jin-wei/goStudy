package main

import (
	"fmt"
	"runtime"
	"time"
)

// 虽然说Go编译器将Go的代码编译成本地可执行代码。
// 不需要像java或者.net那样的语言需要一个虚拟机来运行，
// 但其实go是运行在runtime调度器上的，它主要负责内存管理、垃圾回收、栈处理等等。
// 也包含了Go运行时系统交互的操作，控制goroutine的操作，Go程序的调度器可以很合理的分配CPU资源给每一个任务。
//
// Go1.5版本之前默认是单核执行的。从1.5之后使用可以通过runtime.GOMAXPROCS()来设置让程序并发执行，提高CPU的利用率
func main() {
	fmt.Println("GOROOT", runtime.GOROOT())
	//	获取当前GOROOT目录
	fmt.Println("操作系统", runtime.GOOS)
	//	获取当前操作系统
	fmt.Println("逻辑CPU数量", runtime.NumCPU())

	//设置最大的可同时使用的CPU核数，取逻辑CPU数量
	n := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(n)

	go func() {
		fmt.Println("start...")
		runtime.Goexit() //终止当前groutine
		fmt.Println("end...")
	}()

	time.Sleep(time.Second) //主goroutine 休眠3秒 让子goroutine执行完
	fmt.Println("main_end...")
}
