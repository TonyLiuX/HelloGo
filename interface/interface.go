package main

import "fmt"

// interface，接口：数个方法组成
// 实现了接口中的所有方法，也就实现了这个接口
// go提倡面向结构编程
// 任意类型都实现了空接口

type ISayer interface {
	say()
}

type IMover interface {
	move()
}

type IAnimal interface {
	ISayer
	IMover
}

type Dog struct {
}

func (d Dog) say() {
	fmt.Println("汪汪汪")
}

func (d *Dog) move() {
	fmt.Println("狗跑了！")
}

type Cat struct {
}

func (c Cat) say() {
	fmt.Println("喵喵喵")
}
func (c Cat) move() {
	fmt.Println("猫跑了！")
}

// 结构体嵌套来实现某个接口
type IWashingMachine interface {
	wash()
	dry()
}

type Dryer struct {
}

func (d Dryer) dry() {
	fmt.Println("甩一甩")
}

type Haier struct {
	Dryer
}

func (h Haier) wash() {
	fmt.Println("洗刷刷")
}

func main() {
	var x ISayer
	d := Dog{}
	x = d
	x.say()
	c := Cat{}
	x = c
	x.say()

	// 值接收者实现接口，结构体指针可以赋值给值接收者的接口变量
	var y IMover
	c1 := Cat{}
	y = c1
	y.move()
	c2 := &Cat{}
	y = c2
	y.move()

	// 指针接收者实现接口，结构体类型不可以赋值给指针接收者的接口变量
	d1 := &Dog{}
	y = d1
	y.move()
	//d2 := Dog{}
	//y = d2
	//y.move()

	// 使用结构体嵌套来实现接口
	var wm IWashingMachine
	h1 := Haier{Dryer{}}
	wm = h1
	wm.dry()
}
