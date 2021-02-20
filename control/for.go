package main

import "fmt"

func main() {
	//forDemo()
	//forDemo2()
	//forDemo3()
	forDemo4()
}

func forDemo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func forDemo2() {
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}

func forDemo3() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

// 九九乘法表
func forDemo4() {
FORTAG:
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if j == i {
				fmt.Printf("%d*%d=%d\n", j, i, i*j)
				continue FORTAG
			}
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
	}
}
