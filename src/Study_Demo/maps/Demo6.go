package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "how do you do"

	words := strings.Split(s, " ")

	//fmt.Println(words)
	//fmt.Println(words[0])

	a := make(map[string]int, 10)

	for _, word := range words {
		v, ok := a[word]
		if ok {
			a[word] = v + 1
		} else {
			a[word] = 1
		}
	}
	fmt.Println(a)
}
