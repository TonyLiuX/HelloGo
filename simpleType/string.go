package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")

	// 多行
	s1 := `第一行
第二行
第三行
`
	fmt.Printf("%s\n", s1)

	// len()， 求长度
	s2 := "abcdef"
	fmt.Printf("%s的长度为%d\n", s2, len(s2))

	// +或fmt.Sprintf，拼接
	s3 := "xyz"
	fmt.Printf("%s与%s拼接后%s\n", s2, s3, s2+s3)

	// strings.Split()，分割
	fmt.Printf("%s分割后%v\n", s2, strings.Split(s2, ""))

	// strings.contains，是否包含
	fmt.Printf("%s是否包含%s：%v\n", s2, "de", strings.Contains(s2, "de"))

	// strings.HasPrefix,strings.HasSuffix，前缀/后缀判断
	fmt.Printf("%s是否包含前缀%s：%v\n", s2, "ab", strings.HasPrefix(s2, "ab"))
	fmt.Printf("%s是否包含后缀%s：%v\n", s2, "ef", strings.HasSuffix(s2, "ef"))

	// strings.Index(),strings.LastIndex()，子串出现的位置
	s4 := "abcdefab"
	fmt.Printf("%s出现%s的第一个位置%d\n", s4, "ab", strings.Index(s4, "ab"))
	fmt.Printf("%s出现%s的最后一个位置%d\n", s4, "ab", strings.LastIndex(s4, "ab"))

	// strings.Join，合并字符串数组元素
	fmt.Printf("%v转字符串%s\n", []string{"a", "b", "c"}, strings.Join([]string{"a", "b", "c"}, ""))
}
