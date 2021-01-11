package main

import (
	"fmt"
	"time"
)
import "awesomeProject/myMath"

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

type BInterface interface {
	say()
}

func (b *Books) say() {
	b.title = "修改" + b.title
	fmt.Printf("书本：%v\n", b.title)
}

func main() {
	a := 1
	var b float32
	var c bool
	var d string
	b, bb := 2, 2
	fmt.Printf("%v %v %v %q %v\n", a, b, c, d, bb)

	fmt.Println("This is test!")
	fmt.Printf("1 + 3 = %v\n", myMath.Add(1, 3))
	fmt.Printf("3 - 1 = %v\n", myMath.Sub(3, 1))
	fmt.Printf("123、456最大值： %v\n", myMath.Max(123, 456))

	const Width, Height = 100, 100
	fmt.Printf("Area is %v\n", Width*Height)

	const (
		j = 1 << iota
		k
	)
	fmt.Printf("j: %v, k: %v\n", j, k)

	var e = 2
	e++
	fmt.Printf("e++: %v\n", e)

	// 数组
	var g = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Printf("数组：%v\n", g)
	var g1 = g
	g1[0] = 100
	fmt.Printf("数组: %v\n", g)

	// 指针
	var ptr *int
	ptr = &a
	fmt.Printf("a的值：%v, 存储地址 %v\n", *ptr, ptr)

	const MAX = 3
	h := []int{10, 100, 200}
	var ptr1 [MAX]*int // 指针数组
	for i := 0; i < MAX; i++ {
		ptr1[i] = &h[i]
	}
	for i := 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d, %v\n", i, *ptr1[i], ptr1[i])
	}

	// 结构体
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", bookId: 6495407})

	var book1 Books /* 声明 Book1 为 Books 类型 */
	var book2 Books /* 声明 Book2 为 Books 类型 */
	var book11 Books

	book1.title = "Go 语言"
	book1.author = "www.runoob.com"
	book1.subject = "Go 语言教程"
	book1.bookId = 6495407
	printBook(&book1)

	book11 = book1
	book11.title = "不是Go 语言"
	printBook(&book1)

	book2.title = "Python 教程"
	book2.author = "www.runoob.com"
	book2.subject = "Python 语言教程"
	book2.bookId = 6495700

	printBook(&book2)

	// 切片
	var slice1 = make([]int, 4, 20)
	fmt.Println(slice1, len(slice1), cap(slice1))
	slice1 = []int{1, 2, 3}
	fmt.Println(slice1, len(slice1), cap(slice1))
	slice2 := h[1:]
	fmt.Println(slice2)
	slice3 := slice1[1:]
	fmt.Println(slice3, len(slice3), cap(slice3))

	// Map
	var map1 = make(map[string]string)
	fmt.Printf("map1: %v\n", map1)
	map1["France"] = "巴黎"
	map1["Italy"] = "罗马"
	map1["Japan"] = "东京"
	fmt.Printf("map1: %v\n", map1)
	delete(map1, "France")
	fmt.Printf("删除 France 后，map1: %v\n", map1)

	// Range
	for i, num := range slice1 {
		fmt.Printf("slice[%d]: %d\n", i, num)
	}
	for key, value := range map1 {
		fmt.Printf("%s -> %s\n", key, value)
	}

	// interface
	book1.say()
	book2.say()

	// 线程
	go say("world")
	say("hello")

	m := []int{7, 2, 8, -9, 4, 0}
	n := make(chan int)
	go sum(m[:len(m)/2], n)
	go sum(m[len(m)/2:], n)
	m1, m2 := <-n, <-n
	fmt.Printf("通道传参：%v，%v\n", m1, m2)

	n2 := make(chan int, 2)
	n2 <- 1
	n2 <- 2
	fmt.Printf("%d,", <-n2)
	fmt.Printf("%d\n", <-n2)

	n3 := make(chan int, 10)
	go fibonacci(10, n3)
	for v := range n3 {
		fmt.Println(v)
	}

}

func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book bookId : %d\n", book.bookId)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
		time.Sleep(1000 * time.Millisecond)
	}
	close(c)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
