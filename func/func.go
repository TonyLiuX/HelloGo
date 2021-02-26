package main

import "fmt"

type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func main() {
	fmt.Println(intSum(1, 2))
	fmt.Println(intSum1(1, 2, 3, 4, 5))
	fmt.Println(calc(10, 9))
	fmt.Println(calc1(10, 9))

	// 函数类型
	var c calculation
	c = add
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("1 + 2 = %d\n", c(1, 2))

	f := sub
	fmt.Printf("type if f:%T\n", f)
	fmt.Printf("20 - 10 = %d\n", f(20, 10))

	//
	fmt.Printf("2 + 3 = %d\n", calc2(2, 3, add))

	//
	f1 := adder(10)
	fmt.Printf("f(10):%d\n", f1(10))
	fmt.Printf("f(20):%d\n", f1(20))
	fmt.Printf("f(30):%d\n", f1(30))

	// defer，在defer归属的函数即将返回时，延迟执行语句，多个defer语句逆序执行。defer注册要延迟执行的函数时该函数所有的参数都需要确定其值！
	fmt.Println("start")
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//defer fmt.Println(3)
	fmt.Println("end")

	fmt.Println(def1())
	fmt.Println(def2())
	fmt.Println(def3())
	fmt.Println(def4())

	def51()

	// 训练题
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}

func intSum(x, y int) int {
	return x + y
}

// 可变参数，通常要作为函数的最后一个参数
func intSum1(x int, y ...int) int {
	fmt.Println(x, y) //x是一个切片
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}

// 多返回值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 返回值提前命名
func calc1(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

// 函数作为参数
func calc2(x, y int, op calculation) int {
	return op(x, y)
}

// 闭包
func adder(x int) func(y int) int {
	return func(y int) int {
		x += y
		return x
	}
}

// defer执行循序
func def1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func def2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func def3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func def4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func def5(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func def51() {
	x := 1
	y := 2
	defer def5("AA", x, def5("A", x, y))
	x = 10
	defer def5("BB", x, def5("B", x, y))
	y = 20
}

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	total := 0
	for _, u := range users {
		num := 0
		for _, c := range u {
			switch c {
			case 'e', 'E':
				num += 1
			case 'i', 'I':
				num += 2
			case 'o', 'O':
				num += 3
			case 'u', 'U':
				num += 4
			}
		}
		distribution[u] = num
		total += num
	}
	fmt.Println(distribution)
	return coins - total
}
