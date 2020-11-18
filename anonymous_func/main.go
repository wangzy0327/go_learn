package main

import (
	"fmt"
	"strconv"
)

//AddUpper 累加器  闭包
func AddUpper() func(int) int {
	var n int = 10
	var str = "hello"
	return func(x int) int {
		n = n + x
		str += strconv.Itoa(x)
		fmt.Println("str=", str)
		return n
	}
}

func main() {
	func1 := func(a int, b int) int {
		return a + b
	}(2, 3)
	fmt.Println("匿名函数值：", func1)

	func2 := func(a int, b int) int {
		return a * b
	}

	fmt.Println("匿名函数：", func2(2, 3))
	//使用前面的代码
	f := AddUpper()
	fmt.Println(f(1)) //  11
	fmt.Println(f(2)) //  13
	fmt.Println(f(3)) //  16
	fmt.Println(f(4)) //  16

	s1 := "hello"
	s2 := "hello"
	fmt.Println(s1 == s2)
}
