package main

import (
	"fmt"
	"time"
)

// 递归函数实例1
func recusive(n int) {
	fmt.Println("hello")
	time.Sleep(time.Second)
	if n > 10 {
		return
	}
	recusive(n + 1)

}

// 递归函数实例-2 (5*4*3*2*1=120) n的阶乘
func calc(n int) int {
	if n == 1 {
		return 1
	}
	return calc(n-1) * n
	// 1*2=2
	// 2*3=6
	// 6*4=24
	// 24*5=120

}

// 递归函数实例-3 斐波那契数
func fab(n int) int {
	if n <= 1 {
		return 1
	}
	return fab(n-1) + fab(n-2)
	// n=2 1 + 1 =2
	// n=3 2 + 1 =3
	//
}

func main() {
	fmt.Println("1234")
	// recusive(0)

	n := calc(5)
	fmt.Println(n)

	for i := 0; i < 10; i++ {
		fmt.Println(fab(i))
	}
}
