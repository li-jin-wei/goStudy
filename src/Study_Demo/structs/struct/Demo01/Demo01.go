package main

import "fmt"

type Person struct {
	name    string
	age     int
	sex     string
	address string
}

func main() {
	p := Person{}
	p.name = "张三"
	p.age = 10
	p.sex = "男"
	p.address = "陕西省"
	//fmt.Println(&p)
	fmt.Println(p)
	fmt.Println(p.name, p.age, p.address, p.sex)

	p1 := new(Person)
	p1.age = 22
	fmt.Println(*p1)
	fmt.Println(p1.age)

}
