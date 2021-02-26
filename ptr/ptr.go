package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Printf("a:%v ptr:%p\n", a, b)
	fmt.Printf("b:%p type:%T ptr:%p\n", b, b, &b)
	c := *b
	fmt.Printf("c:%v ptr:%p\n", c, &c)

	// 指针传值
	modify1(a)
	fmt.Printf("a:%v ptr:%p\n", a, &a)
	modify2(&a)
	fmt.Printf("a:%v ptr:%p\n", a, &a)
}

// 值传参
func modify1(x int) {
	x = 100
}

// 指针传参
func modify2(x *int) {
	*x = 100
}
