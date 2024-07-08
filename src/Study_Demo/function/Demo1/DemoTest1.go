package main

import "fmt"

func test() interface{} {
	return 1234
}

func main() {
	var i = test()
	fmt.Println(i)
}
