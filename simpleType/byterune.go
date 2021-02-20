package main

import "fmt"

func main() {
	s1 := "abc"
	// byte,rune
	byteS1 := []byte(s1)
	fmt.Printf("%s转byte:%v\n", s1, byteS1)

	s2 := "中文a"
	runeS1 := []rune(s2)
	fmt.Printf("%s转rune：%v\n", s2, runeS1)

	chinese(s2)
}

func chinese(s string) {
	count := 0
	for _, r := range []rune(s) {
		if r > 255 {
			count += 1
		}
	}
	fmt.Printf("%s中有%d个中文\n", s, count)
}
