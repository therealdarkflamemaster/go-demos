package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// 打开读文件
	f_r, err := os.Open("E:/Programming/goLearning/testFile.txt")
	if err != nil {
		fmt.Println("Open err :", err)
		return
	}

	defer f_r.Close()

	// 创建写文件
	f_w,err := os.Create("E:/Programming/goLearning/testFileCopy.txt")
	if err != nil {
		fmt.Println("Create err :", err)
		return
	}
	defer f_w.Close()

	// 从读文件中获取数据，放入缓冲区

	buf := make([]byte, 4096)

	// 循环，从读文件中获取数据，“原封不动”地写入
	for{
		n,err := f_r.Read(buf)
		if err!=nil && err == io.EOF{
			fmt.Printf("Read finish n =%d \n", n)
			return
		}
		f_w.Write(buf[:n]) // 读多少，写多少
	}







}
