package main

import (
	"fmt"
	"runtime"
)


func test22() {
	defer fmt.Println("cccccccccccc")


	runtime.Goexit() // 退出当前go程
	defer fmt.Println("ddddddddddd")
}

func main () {
	// 匿名函数
	go func() {
		defer fmt.Println("aaaaaaaaaa")
		// defer 延迟调用， 当函数结束运行时调用
		test22()
		fmt.Println("bbbbbbbbbb")
	}()

	for {
		// 保证主go程不结束
	}


}
