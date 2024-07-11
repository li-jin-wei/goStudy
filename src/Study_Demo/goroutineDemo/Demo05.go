package main

import (
	"fmt"
	"math/rand"
	"time"
)

var food = 10

func main() {
	go Relief("灾民1")
	go Relief("灾民2")
	go Relief("灾民3")
	go Relief("灾民4")

	time.Sleep(5 * time.Second)
}
func Relief(name string) {
	for {
		if food > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			//随机休眠时间
			food--
			fmt.Println(name, "抢到救济粮，还剩下", food, "个")
		} else {
			fmt.Println(name, "别抢了 没有粮食了")
			break
		}
	}
}
