package main

import "fmt"

func main() {
	// 十进制
	var a int = 10
	fmt.Printf("%d\n", a)
	fmt.Printf("%b\n", a)

	// 二进制
	b := 0b1011
	fmt.Printf("%b\n", b)
	fmt.Printf("%d\n", b)

	// 八进制
	c := 0o77
	fmt.Printf("%o\n", c)
	fmt.Printf("%d\n", b)

	// 十六进制
	d := 0xff
	fmt.Printf("%x\n", d)
	fmt.Printf("%d\n", d)
}
