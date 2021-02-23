package main

import (
	"fmt"
	"strings"
)

func main() {
	scoreMap := make(map[string]int, 10)
	scoreMap["张三"] = 80
	scoreMap["李四"] = 99
	fmt.Println(scoreMap)

	// 判断map中某个值是否存在
	_, ok := scoreMap["张三"]
	fmt.Printf("%v中是否存在\"张三\"：%v\n", scoreMap, ok)

	// 声明时填充数据
	userInfo := map[string]string{
		"name": "张三",
		"age":  "18",
	}
	fmt.Println(userInfo)

	// 遍历 for range
	for k, v := range scoreMap {
		fmt.Printf("%s, %d\n", k, v)
	}

	// 使用delete()删除键值对
	delete(userInfo, "name")
	fmt.Printf("map[age:18 name:张三]删除键name后：%v\n", userInfo)

	//
	mapSlice()

	sliceMap()

	test1()

	test2()
}

// 元素为map类型的切片
func mapSlice() {
	mapSlice := make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}

// 值为切片的map
func sliceMap() {
	sliceMap := make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "上海", "北京")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}

// 计算字符串中出现单词的个数，比如"how do you do"中how有1个，do有2个，you有1个
func test1() {
	target := "how do you do"
	a := strings.Split(target, " ")
	fmt.Println(a)
	m := make(map[string]int, len(a))
	for _, v := range a {
		_, ok := m[v]
		fmt.Printf("%v是否存在：%v\n", v, ok)
		if ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}
	fmt.Printf("%s单词出现的次数情况：%v\n", target, m)
}

// 引用和切片扩容
func test2() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...) // 扩容后切片指向了不同的内存地址
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
