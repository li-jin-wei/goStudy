package main

import "fmt"

/*
1、定义指针变量
2、为指针变量赋值  &
3、访问指针变量中地址所指向的值   *
* ：在指针类型前面加上 * 号，就是来获取指针所指向的地址的值
*/

func main() {
	var a int = 10
	fmt.Println(a)
	fmt.Printf("%p\n", &a)

	var p *int
	//定义int类型的指针p
	p = &a //指针变量赋值

	fmt.Printf("p存储的指针地址:%p\n", p)
	fmt.Printf("p变量自己的地址:%p\n", &p)
	fmt.Printf("p存储变量的指针地址指向的值:%d\n", *p)

	a = 20
	fmt.Printf("a:%d\n", a)
	fmt.Printf("*p的值:%d\n", *p)

	//	指针的嵌套
	var ptr **int
	ptr = &p
	fmt.Printf("ptr变量存储的指针的地址:%p\n", ptr)
	fmt.Printf("ptr变量自己的地址:%p\n", &ptr)
	fmt.Printf("*ptr变量存储的地址:%p\n", *ptr)
	fmt.Printf("*ptr变量存储的地址中的值:%d\n", *ptr)

	**ptr = 111
	fmt.Println(a)
}
