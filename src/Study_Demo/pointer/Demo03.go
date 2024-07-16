package main

import "fmt"

func main() {
	//	指针数组

	arr := [4]int{1, 2, 3, 4}

	fmt.Println("arr =", arr)
	fmt.Printf("arr =%p\n", &arr)

	var p1 *[4]int
	p1 = &arr

	fmt.Printf("p1指向的地址：%p\n", p1)
	fmt.Printf("p1自己的地址:%p\n", &p1)
	fmt.Println("p1指向的地址的值：", *p1)

	(*p1)[0] = 100
	//通过指针修改数组
	fmt.Println("arr =", arr)
	fmt.Println("p1指向的地址的值:", *p1)

	p1[0] = 200
	fmt.Println("arr =", arr)
	fmt.Println("p1指向的地址的值:", *p1)
}
