package main

import (
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

	// 写文件
	n, err := f.WriteString("Hello World\r\n")

	if err != nil {
		fmt.Println("WriteString error :",err)
		return
	}
	fmt.Println("n of writeString = ",n)


	// 修改读写位置

	offset,_ := f.Seek(-5,io.SeekEnd)
	fmt.Println("offset of Seek is ",offset) // n of Seek is  10

	//

	n,_ = f.WriteAt([]byte("11111"), offset)
	fmt.Println("n of writeAt is ",n)
}
