package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args //获取命令行参数

	if len(list) != 2 {
		fmt.Println("格式为 : go run xxx.go 文件名")
		return
	}

	// 提取文件名
	path := list[1]
	// 获取文件属性
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat err : ", err)
		return
	}

	fmt.Println("File Name : ", fileInfo.Name())
	fmt.Println("File Size : ", fileInfo.Size())

}
