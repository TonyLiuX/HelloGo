package main

// 结构体在内存中占用连续的地址!

// 什么时候使用指针类型接收者？
// 1、需要修改接收者的数据
// 2、接收者进行拷贝代价比较大的对象
// 3、保持一致性，如果某个方法使用了指针接收者，那么其他方法也应该使用指针接收者

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	name string
	city string
	age  int8
}
type person1 struct {
	name, city string
	age        int8
	dreams     []string
}

func (p *person1) setDreams(d []string) {
	p.dreams = d
}

func (p *person1) setDreams1(d []string) {
	p.dreams = make([]string, len(d))
	copy(p.dreams, d)
}

// 构造函数
func newPerson(name, city string, age int8) *Person {
	return &Person{
		name: name,
		city: city,
		age:  age,
	}
}

// method，方法，属于特定的类型，函数不属于任何类型
// 值类型的接受者
func (p Person) dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

// 指针类型的接受者
func (p *Person) setAge(age int8) {
	p.age = age
}

// 值类型的方法不会修改原始数据，只会修改副本
func (p Person) setAge2(age int8) {
	p.age = age
}

// 结构体嵌套
type Address struct {
	province string
	city     string
}

type User struct {
	Id      int `json:"id"`
	Name    string
	Gender  string
	Address Address
}

func main() {
	// 常规创建结构体实例
	var p1 Person
	fmt.Printf("p1:%#v\n", p1)
	p1.name = "张三"
	p1.city = "上海"
	p1.age = 18
	fmt.Printf("p1:%#v\n", p1)

	// 结构体指针来创建实例，使用new创建结构体，结构体指针可以直接访问结构体成员
	p2 := new(Person)
	p2.name = "李四"
	p2.city = "北京"
	p2.age = 30
	fmt.Printf("type of p2:%T, p2:%v\n", p2, p2)

	// 取结构体地址来创建实例，相当于使用new操作符实例化结构体
	p3 := &Person{}
	p3.name = "七米"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("type of p3:%T, p3:%v\n", p3, p3)

	// 未初始化的结构体，各成员为零值
	var p4 Person
	fmt.Printf("p4:%#v\n", p4)

	// 键值对初始化结构体
	p5 := Person{
		name: "王武",
		city: "武汉",
		age:  30,
	}
	fmt.Printf("p5:%#v\n", p5)

	// 构造函数
	p6 := newPerson("嘿嘿", "常州", 10)
	fmt.Printf("p6:%#v\n", p6)

	// 方法
	p5.dream()
	p5.setAge(51)
	fmt.Printf("p5修改年纪后:%v\n", p5)
	p5.setAge2(40)
	fmt.Printf("p5修改年纪后:%v\n", p5)

	// 结构体嵌套
	user1 := &User{
		Name:   "用于1",
		Gender: "男",
		Address: Address{
			province: "江苏",
			city:     "苏州",
		},
	}
	fmt.Printf("user1:%#v\n", user1)

	// json序列化：将结构体专程json字符串，要求成员首字母大写才能序列化出数据
	j1, err := json.Marshal(user1)
	if err != nil {
		fmt.Println("json marshal failed")
	}
	fmt.Printf("user1 json后：%s\n", j1)

	// json反序列化：将json字符串转换成结构体
	str := `{"id":1,"Name":"用于json","Gender":"男","Address":{"province": "江苏"}}`
	ujson1 := &User{}
	err = json.Unmarshal([]byte(str), ujson1)
	if err != nil {
		fmt.Printf("json unmarshal failed!")
	}
	fmt.Printf("ujson1:%#v\n", ujson1)

	// 复制slice和map的注意点，直接复制底层指向数据的指针没有变
	pp1 := person1{
		name: "小王子",
		age:  18,
	}
	slice1 := []string{"吃饭", "睡觉", "打豆豆"}
	//pp1.setDreams(slice1)
	//fmt.Printf("pp1:%v, pp1的dreams的地址:%p, slice1:%v, slice1的地址:%p\n", pp1, &(pp1.dreams[0]), slice1, &slice1[0])
	//slice1[1] = "不睡觉"
	//fmt.Printf("pp1:%v, pp1的dreams的地址:%p, slice1:%v, slice1的地址:%p\n", pp1, &(pp1.dreams[0]), slice1, &slice1[0])

	pp1.setDreams1(slice1)
	fmt.Printf("pp1:%v, pp1的dreams的地址:%p, slice1:%v, slice1的地址:%p\n", pp1, &(pp1.dreams[0]), slice1, &slice1[0])
	slice1[1] = "不睡觉"
	fmt.Printf("pp1:%v, pp1的dreams的地址:%p, slice1:%v, slice1的地址:%p\n", pp1, &(pp1.dreams[0]), slice1, &slice1[0])
}
