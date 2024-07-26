package main

import "strconv"

func TestDemo(s1 string, i2 int) (s2 string, i1 int) {
	s2 = strconv.Itoa(i2)
	i1, _ = strconv.Atoi(s1)
	println(i1, s2)
	return s1, i1
}
