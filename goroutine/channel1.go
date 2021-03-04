package main

// chan默认是双通道，即可读可写
// 单向通道，chan<-：只写，<-chan：只读
// 函数传参和赋值中，都可以将双向通道转成单向通道，但是反过来是不可以的！！
import "fmt"

func counter(c chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("存数据：%d\n", i)
		c <- i
	}
	close(c)
}

func squarer(c1 <-chan int, c2 chan<- int) {
	for v := range c1 {
		fmt.Printf("取数据：%d\n", v)
		c2 <- v * v
	}
	close(c2)
}

// for range从通道中取数据，通道关闭时就会自动退出
func printer(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go counter(ch1)
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go squarer(ch1, ch2)
	// 在主goroutine中从ch2中接收值打印
	printer(ch2)
}
