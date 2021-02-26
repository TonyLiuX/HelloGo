package main

import "fmt"

// 结构体实现面向对象

type Address1 struct {
	province string
	city     string
}

// 匿名结构体
type User1 struct {
	name   string
	gender string
	Address1
}

type User2 struct {
	name    string
	gender  string
	address Address1
}

// 继承
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

type Dog struct {
	feet int
	*Animal
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

func main() {
	u1 := User1{
		name:   "用户1",
		gender: "女",
		Address1: Address1{
			province: "广东",
			city:     "广州",
		},
	}
	fmt.Printf("u1:%#v\n", u1)

	var u2 User1
	u2.name = "用户2"
	u2.gender = "男"
	u2.province = "安徽"
	u2.city = "合肥"
	fmt.Printf("u2:%#v\n", u2)

	var u3 User2
	u3.name = "用户3"
	u3.gender = "男"
	u3.address.province = "安徽"
	u3.address.city = "六安"
	fmt.Printf("u3:%#v\n", u3)

	// 实现继承
	d1 := &Dog{
		feet: 2,
		Animal: &Animal{
			name: "哈哈",
		},
	}
	d1.move()
	d1.wang()
}
