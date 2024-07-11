package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	go Relief1()
	go Relief2()
	go Relief3()
	wg.Wait()
	fmt.Println("main end....")
}

func Relief1() {
	fmt.Println("func1...")
	wg.Done()
}

func Relief2() {
	defer wg.Done()
	fmt.Println("func2...")
}

func Relief3() {
	defer wg.Done()
	fmt.Println("func3...")
}
