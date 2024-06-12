package main

import "fmt"

func sum(a, b int) (sum1, c int) {
	sum1 = a + b
	c = a - b

	return sum1, c

}
func main() {
	fmt.Println(sum(100, 50))

}
