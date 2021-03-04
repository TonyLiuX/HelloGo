package calc

import "fmt"

func init() {
	fmt.Println("package calc init.")
}

// 加
func Add(x, y int) int {
	return x + y
}

// 减
func Minus(x, y int) int {
	return x - y
}

// 乘
func Multi(x, y int) int {
	return x * y
}

// 除
func Divide(x, y int) int {
	return x / y
}
