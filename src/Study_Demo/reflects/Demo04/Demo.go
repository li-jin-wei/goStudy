package main

import (
	"fmt"
	"reflect"
)

type mystruct struct {
	Name string
	Sex  int
	Age  int `json:"age"`
}

func main() {

	typeofmystruct := reflect.TypeOf(mystruct{})
	fmt.Println(typeofmystruct.Name())
	for i := 0; i < typeofmystruct.NumField(); i++ {
		fieldtype := typeofmystruct.Field(i)
		fmt.Println(fieldtype.Name)
		fmt.Println(typeofmystruct.FieldByName(fieldtype.Name))
	}
}
