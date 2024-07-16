package main

import "fmt"

// 值传递和引用传递
func main() {

	//定义数组
	arr1 := [4]int{0, 1, 2, 3}
	arr2 := arr1
	fmt.Println(arr1, arr2)
	arr1[2] = 200
	fmt.Println(arr1, arr2)

	//	数组为值传递，arr1赋值给arr2，只是将arr1的值赋给arr2，修改arr1的值不会影响arr2
	fmt.Println("------------------分割线--------------------------")
	//定义切片
	slice1 := []int{0, 1, 2, 3}
	slice2 := slice1
	fmt.Println(slice1, slice2)

	slice1[2] = 200
	fmt.Println(slice1, slice2)

	fmt.Printf("%p,%p\n", slice1, slice2)
	//0x140000b40a0,0x140000b40a0
	//切片引用的底层数组的地址

	fmt.Printf("%p,%p\n", &slice1, &slice2)
	//0x140000a0018,0x140000a0030
	//切片本身的地址

	//切片是引用传递，修改slice1的值其实就是在修改slice1指向的底层数组的值，所以slice2的值也会改变

}
