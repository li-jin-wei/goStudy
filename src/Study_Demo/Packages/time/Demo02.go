package main

import (
	"fmt"
	"math/rand"
	"time"
)

//随机数

func main() {
	////获取一个随机数
	//num1 := rand.Int()
	//fmt.Println("num1:", num1)
	//
	//num2 := rand.Intn(100)
	//fmt.Println("num2:", num2)

	timestamp := time.Now().Unix()
	rand.Seed(timestamp)
	Demo()

}

func Demo() {
	for i := 0; i < 5; i++ {
		num1 := rand.Intn(100)
		num2 := rand.Intn(10)
		if i == num2 {
			fmt.Println(num2, 10)
			continue
		}
		fmt.Println(num1)
	}
}
