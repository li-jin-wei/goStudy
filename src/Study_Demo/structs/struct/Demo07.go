package main

import "fmt"

//结构体嵌套

type Person struct {
	name    string
	age     int
	address Address
}

type Address struct {
	city, state string
}

func main() {
	var person = Person{}
	person.name = "小杰"
	person.age = 20
	person.address = Address{
		city:  "杭州市",
		state: "中国",
	}
	fmt.Println(person.name)
	fmt.Println(person.age)
	fmt.Println(person.address)
	fmt.Println(person.address.city)
}
