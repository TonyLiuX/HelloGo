package main

import (
	"fmt"
	"math/rand"
)

func getOneRand(c chan<- int64) {
	v := rand.Int63()
	c <- v
}

func sumRand(x int64) int64 {
	var sum int64 = 0
	for x > 0 {
		a := x % 10
		x /= 10
		sum += a
	}
	return sum
}

func sum(id int, c <-chan int64, r chan<- int64) {
	for v := range c {
		randomnum := sumRand(v)
		fmt.Printf("id:%d, %d各位数之和：%d\n", id, v, randomnum)
		r <- randomnum
	}
}

func main() {
	jobChan := make(chan int64)
	resultChan := make(chan int64)

	num := 10 // 随机数的个数

	go func() {
		for i := 0; i < num; i++ {
			getOneRand(jobChan)
		}
		close(jobChan)
	}()

	for i := 0; i < 24; i++ {
		go sum(i, jobChan, resultChan)
	}

	for i := 0; i < num; i++ {
		v := <-resultChan
		fmt.Println(v)
	}

}
