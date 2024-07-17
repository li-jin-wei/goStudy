package main

import "fmt"

//异常 panic
//defer 延迟调用
//panic 异常
//recover 恢复

func main() {
	defer fmt.Println("main--1")
	defer fmt.Println("main--2")
	fmt.Println("main--3")
	testPainc(1)
	defer fmt.Println("main--4")
	fmt.Println("main--5")
}

//func testPainc(num int) {
//	defer fmt.Println("testPainc--1")
//	defer fmt.Println("testPainc--2")
//	fmt.Println("testPainc--3")
//
//	if num == 1 {
//		panic("出现预定的异常----panic")
//	}
//	defer fmt.Println("testPainc--4")
//	fmt.Println("testPainc--5")
//}

func testPainc(num int) {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println("recover执行了...panic msg:", msg)
			fmt.Println("---------程序已恢复-------------")
		}
	}()
	defer fmt.Println("testPainc--1")
	defer fmt.Println("testPainc--2")

	fmt.Println("testPainc--3")

	if num == 1 {
		panic("出现预定的异常----panic")
	}
	defer fmt.Println("testPainc--4")
	fmt.Println("testPainc--5")
}
