package main

import "fmt"

func main() {
	var i int
	fmt.Println(&i)

	i = 10
	fmt.Println(&i)
}
