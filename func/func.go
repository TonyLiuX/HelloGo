package main

import "fmt"

type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func main() {
	fmt.Println(intSum(1, 2))
	fmt.Println(intSum1(1, 2, 3, 4, 5))
	fmt.Println(calc(10, 9))
	fmt.Println(calc1(10, 9))

	// 函数类型
	var c calculation
	c = add
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("1 + 2 = %d\n", c(1, 2))

	f := sub
	fmt.Printf("type if f:%T\n", f)
	fmt.Printf("20 - 10 = %d\n", f(20, 10))

	//
	fmt.Printf("2 + 3 = %d\n", calc2(2, 3, add))

	//
	f1 := adder(10)
	fmt.Printf("f(10):%d\n", f1(10))
	fmt.Printf("f(20):%d\n", f1(20))
	fmt.Printf("f(30):%d\n", f1(30))
}

func intSum(x, y int) int {
	return x + y
}

// 可变参数，通常要作为函数的最后一个参数
func intSum1(x int, y ...int) int {
	fmt.Println(x, y) //x是一个切片
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}

// 多返回值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 返回值提前命名
func calc1(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

// 函数作为参数
func calc2(x, y int, op calculation) int {
	return op(x, y)
}

// 闭包
func adder(x int) func(y int) int {
	return func(y int) int {
		x += y
		return x
	}
}
