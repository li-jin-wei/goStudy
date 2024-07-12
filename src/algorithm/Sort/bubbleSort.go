package main

import "fmt"

func main() {
	arr := []int{9, 18, 119, 29, 8, 97, 3, 89, 1}
	fmt.Println(arr)
	BubbleAsc(arr)
	BubbleDesc(arr)
}

func BubbleAsc(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

func BubbleDesc(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
