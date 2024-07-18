package Demo01

//封装

type person struct {
	name string
	age  int
	sex  string
}

func NewPerson(name string, age int, sex string) *person {
	return &person{
		name: name,
		age:  age,
		sex:  sex,
	}
}

func (p *person) SetName(name string) {
	p.name = name
}
func (p *person) SetAge(age int) {
	p.age = age
}
func (p *person) SetSex(sex string) {
	p.sex = sex
}

func (p *person) GetName() string {
	return p.name
}
func (p *person) GetAge() int {
	return p.age
}
func (p *person) GetSex(name string) string {
	return p.sex
}
