package main

import (
	"errors"
	"fmt"
)

// 自己定义一个错误
// 1、errors.New("xxxxx")
// 2、fmt.Errorf()
// 都会返回  error 对象， 本身也是一个类型

func main() {
	age_err := setAge(-1)
	if age_err != nil {
		fmt.Println(age_err)
	}
	fmt.Printf("%T\n", age_err)

	errInfo1 := fmt.Errorf("我是一个错误信息:%d\n", 500)

	if errInfo1 != nil {
		//处理这个错误
		fmt.Println(errInfo1)
	}
}

func setAge(age int) error {
	if age < 0 {
		age = 3

		//抛出一个错误errprs
		return errors.New("年龄不合法")
	}
	fmt.Println("年龄设置成功：age=", age)
	return nil
}
