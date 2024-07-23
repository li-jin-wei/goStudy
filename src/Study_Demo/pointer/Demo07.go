package main

import "fmt"

func add1(n int) {
	n = n + 100
	fmt.Println("add1函数结束时:", n, &n)
}

func add2(n *int) {
	*n++
	fmt.Println("add2函数结束时:", n, *n)
}

func main() {
	var y = 200
	var yy = &y
	add1(y)
	fmt.Println("调用add1函数之后:", y, &y)

	add2(yy)
	fmt.Println("调用add2函数之后:", y, &y)
}

//最后输出值改变，此次传入的是y的址，址对应的值被改变了，所以外部输出y的值必然也改变；
//大概这样理解了
//总结：这也说明函数默认是值传递（即函数默认传递值到不同的址，原址上的值不变，
//除非标明&传址进而改值）
