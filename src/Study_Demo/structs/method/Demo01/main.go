package main

import "fmt"

type Role struct {
	Name    string
	Ability string //技能
	Level   int
	kill    float64
}

func (r Role) Kungfu() {
	fmt.Printf("我是:%s,我的武功:%s,已经练到%d级了，杀伤力%.1f\n", r.Name, r.Ability, r.Level, r.kill)
}

func (r *Role) Kungfu2() {
	fmt.Printf("我是:%s,我的武功:%s,已经练到%d级了，杀伤力%.1f\n", r.Name, r.Ability, r.Level, r.kill)
}
func main() {
	rwx := Role{"任我行", "吸星大法", 10, 9.9}
	//rwx.Kungfu()
	rwx.Kungfu2()
}
