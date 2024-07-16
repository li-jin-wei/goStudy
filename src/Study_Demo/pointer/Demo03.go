package main

import "fmt"

func main() {
	//	指针数组

	arr := [4]int{1, 2, 3, 4}

	fmt.Println("arr =", arr)
	fmt.Printf("arr =%p\n", &arr)

	var p1 *[4]int
	p1 = &arr

	fmt.Printf("p1指向的地址：%p")

}
