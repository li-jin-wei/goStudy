package main

import "fmt"

func main() {
	fmt.Println(max(99999999, 10000000))
}

func max(num1 int, num2 int) int {
	var result int
	if num2 > num1 {
		result = num2
	} else {
		result = num1
	}
	return result
}
