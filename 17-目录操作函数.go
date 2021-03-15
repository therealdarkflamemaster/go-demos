package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)


// 拷贝mp3文件到指定目录

func copyMP3ToDir(src, dst string) {
	fmt.Println("src :", src)
	// 打开读文件
	f_r, err := os.Open(src)
	if err != nil {
		fmt.Println("Open err :", err)
		return
	}

	defer f_r.Close()

	// 创建写文件
	f_w,err := os.Create(dst)
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


func main() {

	// 获取用户输入的目录路径
	fmt.Println("请输入待查询的目录")
	var path string
	fmt.Scan(&path)

	// 打开目录
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("ReadDir err :",err)
		return
	}

	defer f.Close()

	// 读取目录项

	info, err := f.Readdir(-1) // 读取所有目录项

	// 一次调用 ，获取全部目录项，因为返回的是切片


	// 遍历返回的slice
	for _,fileInfo := range info {
		//if fileInfo.IsDir() { // 是目录
		//	fmt.Println(fileInfo.Name(), "是一个目录")
		//} else {
		//	fmt.Println(fileInfo.Name(), "是一个文件")
		//}
		//if !fileInfo.IsDir() { // 文件
		//	if strings.HasSuffix(fileInfo.Name(), ".jpg") {
		//		fmt.Println("jpg 文件 有: ", fileInfo.Name())
		//	}
		//}
		if !fileInfo.IsDir() { // 文件
			if strings.HasSuffix(fileInfo.Name(), ".mp3") {
				copyMP3ToDir(path + "/" + fileInfo.Name(), "./" + fileInfo.Name())
			}
		}

	}
}
