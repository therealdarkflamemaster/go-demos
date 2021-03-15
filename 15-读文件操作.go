package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main () {
	f, err := os.OpenFile("E:/Programming/goLearning/testFile.txt", os.O_RDWR, 6)

	if err != nil {
		fmt.Println("OpenFile err :", err)
		return
	}

	defer f.Close()

	fmt.Println("successful")

	// 创建一个带有缓冲区的 (用户缓存) reader
	reader := bufio.NewReader(f) //参数是文件指针
	for {
		buf,err := reader.ReadBytes('\n') // 读一行数据
		if err != nil && err == io.EOF { // io.EOF 是结束符
			fmt.Println("ReadBytes finish")
			return
		} else if err != nil {
			fmt.Println("ReadBytes err :", err)
		}

		fmt.Println("buf = ", string(buf))
	}

}