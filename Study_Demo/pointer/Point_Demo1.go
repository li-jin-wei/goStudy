package main

import "fmt"

var i int = 100
var p1 *int = &i

func main() {
	fmt.Println(i, &i)
	fmt.Println(*p1, &p1)
}
