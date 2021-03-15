package main

import (
	"fmt"
	"time"
)


// 全局定义channel, 完成数据同步
var channel  = make(chan int) // 传递的数据和channel的数据没有关系


// 定义 打印机
func printer (s string) {
	for _,ch := range s {
		fmt.Printf("%c", ch) // 屏幕 ， stdout文件
		time.Sleep(300*time.Millisecond)
	}
}

// 定义两个人使用打印机
func person1 () { // person1 先执行

	printer("Hello")
	channel <- 8 // 随意写什么数字
	// 一个数字写入，如果没有人读，便会堵塞
}

func person2 () { // person2 后执行
	<- channel
	printer("world")
}


func main() {

	go person1()
	go person2()

	for {
		;
	}
}