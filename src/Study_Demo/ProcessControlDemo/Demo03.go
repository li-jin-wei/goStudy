package main

import "fmt"

func main() {
	fmt.Println("1")
	defer prints("100")
	getData("10")
	fmt.Println("3")
}

func getData(s string) string {
	fmt.Println("getData", s)
	return s
}

func prints(n string) {
	fmt.Println(n)
}
