package main

import (
	"fmt"
)

// 递归函数
func main() {
	//num := getSum2(100)
	//fmt.Println(num)

	//fmt.Println(fb(5))
	fmt.Println(fib(3))
}

func getSum2(i int) int {
	fmt.Println(i)
	if i == 0 {
		return 0
	}
	return getSum2(i-1) + i
	//	100+99+98
}

// 使用递归函数解决斐波那契数列问题，但使用递归方法计算较大的n值时会非常低效，因为它会重复计算很多子问题
// 为了计算fb(5)，它会分别计算fb(4)和fb(3)，而为了计算fb(4)，它又会计算fb(3)和fb(2)，这导致fb(3)被计算了两次。
// 对于较大的n，这种重复计算会导致性能急剧下降。因此，在实际应用中，通常会使用动态规划或迭代方法来优化斐波那契数列的计算。
func fb(n int) int {
	if n <= 2 {
		return 1
	} else {
		return fb(n-1) + fb(n-2)
	}
}

// 使用动态规划解决斐波那数列问题
func fib(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	//由于切片是从0开始的，而斐波那契数列通常从1开始计数，因此这里使用了n+1的长度来确保能够存储到第n个数的值。
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 1000000007
		//	对1000000007（即10^9 + 7，一个常用的大质数）取模。这样做通常是为了防止整数溢出
	}
	return dp[n]
}
