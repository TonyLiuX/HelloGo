package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const fileName = "./test.txt"

// 检查文件是否存在
func checkFileExist(name string) bool {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// 打开文件，只读
func open(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		fmt.Printf("创建文件%s\n", name)
		file, err = os.Create(name)
	}
	return file
}

// 打开文件，读写
func open1(name string) *os.File {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return file
	}
	return file
}

// 读取文件，使用File.Read()，最多读取固定长度（字节数）的数据，当出现0或者io.EOF时结束读取
func readFile(f *os.File, o []byte) {
	n, err := f.Read(o)
	if err == io.EOF {
		fmt.Println("文件读取完毕")
		return
	}
	if err != nil {
		fmt.Println("read failed.")
		return
	}
	fmt.Printf("读取了%d字节数据\n", n)
}

// 循环读取文件，使用File.Read()
func readFile1(f *os.File) []byte {
	content := make([]byte, 0)
	temp := make([]byte, 128)
	for {
		n, err := f.Read(temp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return content
		}
		content = append(content, temp[:n]...)
	}
	return content
}

// bufio读取文件
func readFile2(f *os.File) string {
	reader := bufio.NewReader(f)
	var outStr string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				if len(line) != 0 {
					outStr += line
				}
				fmt.Println("文件读完了！")
				break
			} else {
				fmt.Println("read file failed, err:", err)
				return outStr
			}
		}
		if len(line) != 0 {
			outStr += line
		}
	}
	return outStr
}

// ioutil，一次性读取整个文件，这种方式读取最快？？？
func readFile3(f *os.File) string {
	content, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return ""
	}
	return string(content)
}

// 写入文件，file.Write()，覆盖写入内容！
func writeFile(f *os.File, str string) {
	//n, err := f.Write([]byte(str))
	n, err := f.WriteString(str)
	if err != nil {
		fmt.Println("写入文件出错!", n)
		return
	}
}

// bufio.NewWriter().WriteString()，覆盖写入内容！
func writeFile1(f *os.File, str string) {
	writer := bufio.NewWriter(f)
	for i := 0; i < 10; i++ {
		n, err := writer.WriteString(str) //将数据先写入缓存
		if err != nil {
			fmt.Println("写入出错！", n)
			break
		}
	}
	writer.Flush() // 将缓存写入文件
}

// ioutil.WriteFile()，覆盖写入内容！
func writeFile2(f *os.File, str string) {
	err := ioutil.WriteFile(f.Name(), []byte(str), 0o666)
	if err != nil {
		fmt.Println("写入出错！", err)
	}
}

func main() {
	// 打开文件
	var file *os.File
	//file = open(fileName)
	file = open1(fileName)

	// 读取文件，一次性读取xxx个字节
	//out := make([]byte, 0)
	//readFile(file, out)
	//fmt.Println("读取文件的数据：", out)

	// 循环读取文件
	//out1 := readFile1(file)
	//fmt.Println("读取文件的数据：", out1)

	// bufio按行读取文件
	//outStr := readFile2(file)
	//fmt.Println("读取文件的数据：", outStr)

	// ioutil读取整个文件
	//outStr := readFile3(file)
	//fmt.Println("读取文件的数据：", outStr)

	str1 := "hello world\n"
	// 写入文件
	//writeFile(file, str1)

	//writeFile1(file, str1)

	writeFile2(file, str1)

	// 关闭文件
	if file != nil {
		defer file.Close()
	}
}
