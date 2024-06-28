package main

import "fmt"

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)
	s1 := a[6]
	fmt.Printf("%T", s1)
}
