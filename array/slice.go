package main

// 判断切片是否为空，需要使用len(s)方法，而不是s != nil
// 切片的拷贝是指针拷贝！

import (
	"fmt"
	"sort"
)

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

	// append，扩容策略：1、首先所需容量超过原始容量两倍，则扩容至所需容量；2、原始长度小于1024，则扩容至原始容量两倍；3、原始长度大于或等于1024，则每次增加25%原始容量，直到达到所需容量；4、否则最终扩容到所需容量
	var s3 []int
	s3 = append(s3, 1, 2, 3)
	fmt.Printf("切片s3%v, len(s3):%d, cap(s3):%d\n", s3, len(s3), cap(s3))

	// copy
	s4 := []int{1, 2, 3}
	s5 := make([]int, 3, 3)
	fmt.Printf("s4:%v, s5:%v\n", s4, s5)
	copy(s5, s4)
	fmt.Printf("copy(s4, s5)后s4:%v, s5:%v\n", s4, s5)
	s5[0] = 1000
	fmt.Printf("修改s5[0] = 1000后，s4:%v, s5:%v\n", s4, s5)

	// 删除元素，go中没有专门的删除元素的方法
	s6 := []int{1, 2, 3, 4, 5, 6}
	s6 = append(s6[:2], s6[3:]...) // 删除第三个元素
	fmt.Printf("s6:[1, 2, 3, 4, 5, 6]删除第三个元素后：%v\n", s6)

	test1()
	test2()
}

// test
func test1() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
}

func test2() {
	var a = [...]int{3, 7, 8, 9, 1}
	var s = make([]int, len(a), len(a))
	copy(s, a[:])
	fmt.Println(s)
	sort.Ints(s)
	fmt.Printf("%v排序后%v\n", a, s)
}
