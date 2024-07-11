package main

import (
	"fmt"
	"sync"
)

var food = 10

// 同步等到组对象
var wg sync.WaitGroup

// 创建一把锁
var mutex sync.Mutex

func main() {
	wg.Add(4)
	go Relief("灾民1")
	go Relief("灾民2")
	go Relief("灾民3")
	go Relief("灾民4")
	wg.Wait()
}
func Relief(name string) {
	defer wg.Done()
	for {
		mutex.Lock()
		if food > 0 {
			food--
			fmt.Println(name, "抢到救济粮，还剩下", food, "个")
		} else {
			mutex.Unlock()
			fmt.Println(name, "别抢了 没有粮食了")
			break
		}
		mutex.Unlock()
	}
}
