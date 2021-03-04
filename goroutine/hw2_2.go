package main

import (
	"fmt"
	"os"
)

// 使用接口的方式实现一个既可以往终端写日志也可以往文件写日志的简易日志库

type Writer1 interface {
	consoleWrite()
	fileWrite()
}

type logger1 struct {
	content  string
	fileName string
}

func (l logger1) consoleWrite() {
	fmt.Println("终端写日志", l.content)
}

func (l logger1) fileWrite(c chan bool) {
	file, err := os.OpenFile(l.fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("打开文件%s失败！\n", l.fileName)
		return
	}
	file.Seek(0, 2)
	_, err = file.WriteString(l.content + "\n")
	if err != nil {
		fmt.Printf("写入文件失败！\n")
	} else {
		fmt.Printf("成功写入文件！\n")
		c <- true
	}
	file.Close()
}

func main() {
	log := logger1{
		content:  "哇哈哈",
		fileName: "./log1.txt",
	}
	ch1 := make(chan bool)
	log.consoleWrite()
	go log.fileWrite(ch1)
	r := <-ch1
	if r {
		fmt.Printf("异步写文件完成！\n")
	}
}
