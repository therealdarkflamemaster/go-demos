package main

import (
	"fmt"
	"os"
)

//func main() {
//	f, err := os.Create("E:/Programming/goLearning/testFile.txt")
//	if err != nil {
//		fmt.Println("create err", err)
//		return
//	}
//
//	defer f.Close() // 关闭文件，避免内存泄漏
//	fmt.Println("successful")
//}

//func main() {
//	// 只读方式进行打开
//	f,err := os.Open("E:/Programming/goLearning/testFile.txt")
//	if err != nil {
//		fmt.Println("open err :", err)
//		return
//	}
//	defer f.Close()
//
//	_, err = f.WriteString("aaa")
//	if err != nil {
//		fmt.Println("writing err :", err)
//		return
//	}
//	// writing err : write E:/Programming/goLearning/testFile.txt: Access is denied.
//
//	fmt.Println("successful")
//}

func main() {
	// 只读方式进行打开
	f,err := os.OpenFile("E:/Programming/goLearning/testFile.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("open err :", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString("aaa")
	if err != nil {
		fmt.Println("writing err :", err)
		return
	}

	fmt.Println("successful")
}