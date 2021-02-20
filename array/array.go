package main

import (
	"fmt"
)

func main() {
	var testArr [3]int
	numArr := [3]int{1, 2}
	strArr := [...]string{"a", "b", "c", "d"}
	fmt.Println(testArr)
	fmt.Println(numArr)
	fmt.Println(strArr)

	// 遍历
	for i := 0; i < len(strArr); i++ {
		fmt.Println(strArr[i])
	}
	for index, value := range numArr {
		fmt.Printf("%d,%d\n", index, value)
	}

	// 多维数组
	arr := [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(arr)

	// 求和
	arr1 := [6]int{1, 3, 5, 7, 8, 7}
	fmt.Printf("%v元素之和：%d\n", arr1, sum(arr1))

	// 求两个元素之和为8的下标
	twoSum(arr1, 8)
}

// 求数组各元素之和
func sum(arr [6]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

// 求两个元素之和为某个值的下标
func twoSum(arr [6]int, someValue int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if arr[i]+arr[j] == someValue {
				fmt.Printf("位置%d和位置%d之和等于%d\n", i, j, someValue)
				continue
			}
		}
	}
}
