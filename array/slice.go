package main

// 判断切片是否为空，需要使用len(s)方法，而不是s != nil
// 切片的拷贝是指针拷贝！

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3]
	fmt.Printf("s:%v, len(s):%d, cap(s):%d\n", s, len(s), cap(s))
	fmt.Printf("s:%v, len(s):%d, cap(s):%d\n", a[1:], len(a[1:]), cap(a[1:]))
	fmt.Printf("s:%v, len(s):%d, cap(s):%d\n", a[:3], len(a[:3]), cap(a[:3]))
	fmt.Printf("s:%v, len(s):%d, cap(s):%d\n", a[:], len(a[:]), cap(a[:]))

	// make创建切片
	s1 := make([]int, 2, 10)
	fmt.Printf("s1:%v, len(s1):%d, cap(s1):%d\n", s1, len(s1), cap(s1))

	s2 := []int{1, 2}
	fmt.Printf("s2:%v, len(s2):%d, cap(s2):%d\n", s2, len(s2), cap(s2))
}
