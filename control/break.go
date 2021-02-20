package main

import "fmt"

func main() {
	breakDemo1()
	breakDemo2()
}

func breakDemo1() {
	i := 0
	for ; i < 10; i++ {
		if i < 5 {
			fmt.Println(i)
		} else {
			break
		}
	}
}

// break后可以添加标签，表示退出标签对应的代码块，标签要求必须定义在对应的for、switch、select的代码块上
func breakDemo2() {
BREAKTAG:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKTAG
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
}
