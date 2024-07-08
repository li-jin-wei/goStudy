package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 36
	//取数据类型对象
	fmt.Println(reflect.TypeOf(a))
	//取值对象
	fmt.Println(reflect.ValueOf(a))

	atype := reflect.TypeOf(a)
	fmt.Println(atype.Name())
	avalue := reflect.ValueOf(a)
	fmt.Println(avalue.Int())

}
