package main

import "fmt"

// 继承01
type Ouyangcrazy struct {
	Name string
	Age  int
	//Ability string
}

type YangGuo struct {

	//Ouyangcrazy 匿名字段，实现了继承
	Ouyangcrazy
	Address string
}

func (o *Ouyangcrazy) ToadKongfu() {
	fmt.Println(o.Name, "蛤蟆功")
}

func (y *YangGuo) NewKongfu() {
	fmt.Println(y.Name, "子类自己的新功夫！")
}

func (y *YangGuo) ToadKongfu() {
	fmt.Println(y.Name, "的新蛤蟆功")
}

func main() {
	o := &Ouyangcrazy{"欧阳疯", 70}
	o.ToadKongfu()
	y := &YangGuo{Ouyangcrazy{"杨过", 18}, "古墓"}
	y.NewKongfu()
	y.ToadKongfu()
}
