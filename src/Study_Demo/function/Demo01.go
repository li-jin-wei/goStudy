package main

import "fmt"

func main() {
	//fmt.Println(max(99999999, 10000000))
	//fmt.Println(Test(17.9, 18.3))
	//t, x, _, _ := Test2(17.9, 18.3)
	//fmt.Println(t, x)
	num := getSum(1, 2, 3, 4, 5)
	fmt.Println(num)

}

// max 函数，多参数一个返回值
func max(num1 int, num2 int) int {
	var result int
	if num2 > num1 {
		result = num2
	} else {
		result = num1
	}
	return result
}

// Test 函数，多参数多个返回值,如果return时不指定返回值顺序，则按照参数顺序返回
func Test(len, wid float64) (zc float64, area float64) {
	area = len * wid
	zc = (len + wid) / 2
	return
}

// Test2 函数，多参数多个返回值,返回值名称可以省略
func Test2(len, wid float64) (float64, float64, float64, float64) {
	area := len * wid
	zc := (len + wid) / 2

	return area, zc, 10, 20
}

// getSum 函数，可变参数
func getSum(nums ...int) int {
	sum := 0
	fmt.Println("可变参数长度为:", len(nums))
	for i := 0; i < len(nums); i++ {
		fmt.Println("可变参数", i, nums[i])
		sum += nums[i]
	}
	return sum
}
