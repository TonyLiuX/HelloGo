package main

import (
	"fmt"
)

// 通道channel，遵循先进先出规则，具有具体类型
// 通道是引用类型，声明后为空，使用时必须分配内存，可以使用make来创建内存
func main() {
	var ch1 chan int
	ch1 = make(chan int)
	fmt.Printf("ch1:%#v\n", ch1)

	// 发送、接收、关闭
	go rec(ch1)
	ch1 <- 10 // 发送一个数据，无缓冲通道也为阻塞通道、同步通道。这一步将阻塞当前线程，直到另一个线程有变量接收通道中的数据，所以可以执行rec函数

	// 有缓冲通道
	var ch2 chan string
	ch2 = make(chan string, 1) // 有缓存的通道，达到容量时才会阻塞当前线程
	ch2 <- "haha"
	s1 := <-ch2
	ch2 <- "heihei"
	fmt.Printf("s1接收到值：%s\n", s1)
}

func rec(c chan int) {
	fmt.Println("准备接收通道中的值")
	ret := <-c // 接收数据会阻塞所在线程，直到从通道中接收到数据
	fmt.Printf("接收到值:%d\n", ret)
}
