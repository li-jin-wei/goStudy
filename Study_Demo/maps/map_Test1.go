package main

import "fmt"

func main() {
	maps1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
	fmt.Println(maps1)

	maps2 := make(map[int]int)
	maps2[1] = 1
	maps2[100] = 100
	fmt.Println(maps2)
	for nums := range maps1 {
		fmt.Println(maps1[nums])
	}
	for key, value := range maps1 {
		fmt.Println(key, value)
	}
}
