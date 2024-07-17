package main

import (
	"encoding/json"
	"fmt"
)

type Prescripts struct {
	Name     string
	Unit     string
	Additive *Prescripts
}

func main() {
	p := &Prescripts{}
	p.Name = "鹤顶红"
	p.Unit = "1.2kg"
	p.Additive = &Prescripts{
		Name: "砒霜",
		Unit: "0.5kg",
	}
	//结构体专json字符串
	marshal, err := json.Marshal(p)
	if err != nil {
		return
	}
	fmt.Println("json=", string(marshal))
}
