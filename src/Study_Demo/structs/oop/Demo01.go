package main

import "fmt"

//封装

type Person struct {
	Name string
	Age  int
	Sex  string
}

func NewPerson(name string) *Person {
	return &Person{
		Name: name,
	}
}
func (p *Person) setName(name string) {
	p.Name = name
}
func (p *Person) setAge(age int) {
	p.Age = age
}
func (p *Person) setSex(sex string) {
	p.Sex = sex
}

func (p *Person) getName() string {
	return p.Name
}
func (p *Person) getAge() int {
	return p.Age
}
func (p *Person) getSex(name string) string {
	return p.Sex
}

func main() {
	p := NewPerson("王二")
	p.getName()
	p.setAge(18)
	p.setSex("女")
	p.setName("李四")
	fmt.Println(*p)
}
