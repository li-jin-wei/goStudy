package main

import "fmt"

func main() {
TestLabel:
	for a := 0; a < 10; a++ {
		if a == 5 {
			a += 1
			goto TestLabel
		}
		fmt.Println(a)
	}
}
