package main

import "fmt"

func main() {
	continueDemo1()
	//continueDemo2()
}

func continueDemo1() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}

// continue后添加标签，表示开始标签对应的循环
func continueDemo2() {
forloop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}
