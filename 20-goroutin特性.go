package main

import (
	"fmt"
	"time"
)

func main() {

	go func() { // 创建一个 子 go 程
		for  i := 0;i<5;i++ {
			fmt.Println("... in the pesudo process ...")
			time.Sleep(time.Second)
		}

	}() // 匿名函数

	fmt.Println("... in the main process ...")

	//for i:=0;i<5;i++ {  // 主 go 程
	//	fmt.Println("... in the main process ...")
	//	time.Sleep(time.Second)
	//	if i == 2 {
	//		break
	//	}
	//}


}