package main

import "fmt"

//深浅拷贝
//深拷贝是指将值类型的数据进行拷贝的时候，拷贝的是数值本身，所以值类型的数据默认都是深拷贝。
//浅拷贝指的是拷贝的引用地址，修改拷贝过后的数据,原有的数据也被修改。
//那么如何做到引用类型的深拷贝？也就是需要将引用类型的值进行拷贝。修改拷贝的值不会对原有的值造成影响。

func main() {
	//slice1 := []int{1, 2, 3, 4}
	//slice2 := make([]int, 0)
	//for _, v := range slice1 {
	//	slice2 = append(slice2, v)
	//}
	//fmt.Println(slice1)
	//fmt.Println(slice2)
	//fmt.Printf("%p,%p\n", slice1, slice2)

	s2 := []int{1, 2, 3, 4}
	s3 := []int{7, 8, 9}
	//	深拷贝,使用copy函数
	//copy(s2, s3)
	//fmt.Println(s2)
	//fmt.Println(s3)

	copy(s3, s2[2:])
	fmt.Println(s2)
	fmt.Println(s3)

}
